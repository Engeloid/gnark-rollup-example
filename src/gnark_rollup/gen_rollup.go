package rollup

import (
	// "hash"

	"bytes"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"

	. "rollup/src/tools"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tidwall/gjson"
)

const (
	FileToRProof     = "./test/rollupg16.proof"
	FileToRPK        = "./test/rollupg16.pk"
	FileToRVK        = "./test/rollupg16.vk"
	FileToRCS        = "./test/rollupg16.cs"
	FileToZKVerifier = "./contracts/gnark_verifier.sol"
	ColorReset       = "\033[0m"
	ColorBlue        = "\033[33m"
	ColorGreen       = "\033[32m"
	ColorRed         = "\033[31m"
	Colorbold        = "\033[1m"
)

func (o *Operator) AllocateSlices() {
	o.Witnesses.allocateSlicesMerkleProofs()
}

func (o *Operator) CreateZKProof() (groth16.Proof, error) {
	var cs constraint.ConstraintSystem
	var pk groth16.ProvingKey
	var vk groth16.VerifyingKey
	var err error
	// If one file does not exist, create constraint system, new verifier contract etc.
	if FileExists(FileToRVK) || FileExists(FileToRPK) || FileExists(FileToRCS) || FileExists(FileToZKVerifier) {
		// Otherwise load them from local repo
		cs, pk, vk, err = o.SetupZKVerifierContract()
		if err != nil {
			log.Fatalf("Could not create zk verifier contract file: %v", err)
		}
	} else {
		// Load Constraint system
		cs, err = LoadCSFromFile(FileToRCS)
		if err != nil {
			log.Fatalf("Error loading CS: %v", err)
		}
		// Load Proofing key
		pk, err = LoadPKFromFile(FileToRPK)
		if err != nil {
			log.Fatalf("Error loading PK: %v", err)
		}
		// Load Verifying key
		vk, err = LoadVKFromFile(FileToRVK)
		if err != nil {
			log.Fatalf("Error loading VK: %v", err)
		}
		fmt.Println("Done Loading zk files")
	}

	// create witness and proof
	witness, err := frontend.NewWitness(&o.Witnesses, ecc.BN254.ScalarField())
	if err != nil {
		log.Fatalf("Could not create witnesses: %v", err)
	}
	publicWitness, err := witness.Public()
	if err != nil {
		log.Fatalf("Could not create public witness: %v", err)
	}
	proof, err := groth16.Prove(cs, pk, witness)
	if err != nil {
		log.Fatalf("Could not create proof using pk and witnesses: %v", err)
	}
	if !FileExists(FileToRProof) {
		err := SaveProofToFile(proof, FileToRProof)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// verify the proof
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		log.Fatalf("invalid proof: %v", err)
	}

	// Save proof to file for testing
	err = SaveProofToFile(proof, FileToRProof)
	if err != nil {
		log.Fatalln(err)
	}

	return proof, err
}

func (o *Operator) SetupZKVerifierContract() (constraint.ConstraintSystem, groth16.ProvingKey, groth16.VerifyingKey, error) {
	var circuit Circuit
	// allocating slice for the Merkle paths
	for i := 0; i < BatchSizeCircuit; i++ {
		circuit.MerkleProofReceiverBefore[i].Path = make([]frontend.Variable, 5)
		circuit.MerkleProofReceiverAfter[i].Path = make([]frontend.Variable, 5)
		circuit.MerkleProofSenderBefore[i].Path = make([]frontend.Variable, 5)
		circuit.MerkleProofSenderAfter[i].Path = make([]frontend.Variable, 5)
	}
	cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		log.Fatalln("Could not compile circuit:", err)
	}

	pk, vk, err := groth16.Setup(cs)
	if err != nil {
		log.Fatalln("Could not setup pk and vk:", err)
	}

	// Save files for zkProof if there are none for testing
	fmt.Println("Proof files will be created")
	err = SavePKToFile(pk, FileToRPK)
	if err != nil {
		log.Fatalln(err)
	}
	err = SaveVKToFile(vk, FileToRVK)
	if err != nil {
		log.Fatalln(err)
	}
	err = SaveCSToFile(cs, FileToRCS)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Colorbold + ColorBlue + "zkVerifier contract will be created" + ColorReset)
	// export solidity contract
	fSolidity, err := os.Create(filepath.Join("./contracts/", "gnark_verifier.sol"))
	if err != nil {
		log.Fatalln(err)
	}

	err = vk.ExportSolidity(fSolidity)
	if err != nil {
		log.Fatalln(err)
	}

	err = fSolidity.Close()
	if err != nil {
		log.Fatalln(err)
	}

	return cs, pk, vk, err
}

func (op *Operator) AddTransaction(tx Transfer) {
	op.q.listTransfers <- tx
}

func (op *Operator) StartProcessingTxs() {
	go func() {
		for tx := range op.q.listTransfers {
			// Process the transaction (e.g., add to the current batch)
			op.H.Reset()
			err := op.UpdateState(tx, op.batch)
			if err != nil {
				log.Fatalf("Error updating operator state: %v", err)
			}
			//fmt.Printf("TX index: %d amount: %d\n", op.batch, tx.amount.Uint64())

			op.batch++
			if op.batch >= BatchSizeCircuit {
				// Process the batch
				// For example, create a SNARK proof for these transactions

				// Reset the batch count
				op.batch = 0
			}
		}
	}()
}

// returns a single account on the rollup at the next free index in the
// merkle tree
func (o *Operator) addAccount(i uint64, address string) (Account, eddsa.PrivateKey) {

	var acc Account
	var rnd fr.Element
	var privkey eddsa.PrivateKey

	// create account
	acc.index = uint64(i)
	acc.Nonce = uint64(0)
	acc.balance.SetZero()
	rnd.SetUint64(uint64(i))
	src := rand.NewSource(int64(i))
	r := rand.New(src)

	pkey, err := eddsa.GenerateKey(r)
	if err != nil {
		panic(err)
	}
	privkey = *pkey

	acc.PubKey = privkey.PublicKey

	// fill the index map of the operator
	b := acc.PubKey.A.X.Bytes()
	o.AccountMap[string(b[:])] = acc.index
	// o.AccountMap[privkey.PublicKey.A.X.String()] = acc.index

	// Put pubkey.x into mapping to on-chain addresses
	o.EthToRollupMap[address] = string(b[:])

	return acc, privkey
}

// Applies Deposit event, i.e., checks whether account exists, if not ,
// creates one at the right index and increases balance according to the
// deposit amount.
// func (o *Operator) ApplyDepositEvent(address string, amount uint64) ([]byte, [][]byte, uint64) {
func (o *Operator) ApplyDepositEvent(address string, amount uint64, segmentSize int) ([]byte, [][]byte) {
	var indexInRollup uint64
	var depositAcc Account
	var err error
	if rollupAddress, exists := o.EthToRollupMap[address]; exists {
		// ethAddress is in the map, and rollupAddress contains the associated value
		// fmt.Println("Found in map, Rollup Address:", rollupAddress)
		indexInRollup = o.AccountMap[rollupAddress]
		depositAcc, err = o.readAccount(indexInRollup)
		if err != nil {
			log.Fatalf("failed to read deposit account: %v", err)
		}
	} else {
		// ethAddress is not in the map
		var key eddsa.PrivateKey
		fmt.Println("Address not found in map, account will be created")
		if len(o.AvailableIndices) > 0 {
			// Use an available index
			indexInRollup, o.AvailableIndices = o.AvailableIndices[0], o.AvailableIndices[1:]
		} else {
			// Use the next new index
			indexInRollup = o.NextIndex
			o.NextIndex++
		}
		//fmt.Println("NextIndex in creation is: ", indexInRollup)
		depositAcc, key = o.addAccount(indexInRollup, address)
		o.Keys = append(o.Keys, key)
	}

	var buf bytes.Buffer
	_, err = buf.Write(o.HashState)
	if err != nil {
		log.Fatalf("failed to read hash state of operator: %v", err)
	}
	// Then create merkle proof
	oldMerkleRoot, oldMerkleProofDeposit, _, err := merkletree.BuildReaderProof(&buf, o.H, segmentSize, indexInRollup)
	if err != nil {
		log.Fatalf("Failed to create merkle proof: %v", err)
	}
	merkletree.VerifyProof(o.H, oldMerkleRoot, oldMerkleProofDeposit, indexInRollup, 16)

	buf.Reset() // the buffer needs to be reset

	// increase balance by amount
	depositAcc.balance.SetUint64(depositAcc.balance.Uint64() + amount)

	// update the state of the operator
	copy(o.State[int(indexInRollup)*SizeAccount:], depositAcc.Serialize())
	o.H.Reset()
	_, _ = o.H.Write(depositAcc.Serialize())
	hashDepositAcc := o.H.Sum([]byte{})
	copy(o.HashState[int(indexInRollup)*segmentSize:(int(indexInRollup)+1)*segmentSize], hashDepositAcc)
	// create new roothash
	_, err = buf.Write(o.HashState)
	if err != nil {
		log.Fatalf("failed to read hash state of operator: %v", err)
	}
	o.H.Reset()
	newRootHash, _, _, err := merkletree.BuildReaderProof(&buf, o.H, segmentSize, indexInRollup)
	if err != nil {
		log.Fatalf("Failed to build new merkle tree: %v", err)
	}
	fmt.Printf("Balance of account %d is: %d\n", depositAcc.index, depositAcc.balance.Uint64())
	return newRootHash, oldMerkleProofDeposit //, indexInRollup
}

func (o *Operator) ApplyWithdrawalEvent(address string, amount uint64, segmentSize int) ([]byte, [][]byte) {
	// Assume similar setup for mapping and account retrieval as in ApplyDepositEvent
	var indexInRollup uint64
	var withdrawalAcc Account
	var err error
	if rollupAddress, exists := o.EthToRollupMap[address]; exists {
		// Account exists, retrieve it
		indexInRollup = o.AccountMap[rollupAddress]
		withdrawalAcc, err = o.readAccount(indexInRollup)
		if err != nil {
			log.Fatalf("failed to read withdrawal account: %v", err)
		}

		// Check if balance is sufficient for withdrawal
		if withdrawalAcc.balance.Uint64() < amount {
			log.Fatalf("Insufficient balance for withdrawal")
		}
	} else {
		log.Fatalf("Account does not exist for withdrawal")
	}

	// The rest of the function mirrors ApplyDepositEvent:
	// - Updating the operator's state
	// - Creating a new merkle proof for the updated account state
	// - Possibly, returning the new root hash and the old merkle proof

	// These steps ensure the system's state is updated correctly after the withdrawal and allows for verification of the state change.

	var buf bytes.Buffer
	_, err = buf.Write(o.HashState)
	if err != nil {
		log.Fatalf("failed to read hash state of operator: %v", err)
	}
	// Create merkle proof for the account before updating its balance
	oldMerkleRoot, oldMerkleProofWithdrawal, _, err := merkletree.BuildReaderProof(&buf, o.H, segmentSize, indexInRollup)
	if err != nil {
		log.Fatalf("Failed to create merkle proof: %v", err)
	}
	// Verify old state
	merkletree.VerifyProof(o.H, oldMerkleRoot, oldMerkleProofWithdrawal, indexInRollup, 16)

	// Decrease balance
	withdrawalAcc.balance.SetUint64(withdrawalAcc.balance.Uint64() - amount)

	// Update operator's state for the account
	buf.Reset()
	copy(o.State[int(indexInRollup)*SizeAccount:], withdrawalAcc.Serialize())
	o.H.Reset()
	_, _ = o.H.Write(withdrawalAcc.Serialize())
	hashWithdrawalAcc := o.H.Sum([]byte{})
	copy(o.HashState[int(indexInRollup)*segmentSize:(int(indexInRollup)+1)*segmentSize], hashWithdrawalAcc)

	// Create new root hash after the withdrawal
	_, err = buf.Write(o.HashState)
	if err != nil {
		log.Fatalf("failed to read hash state of operator: %v", err)
	}
	o.H.Reset()
	newRootHash, _, _, err := merkletree.BuildReaderProof(&buf, o.H, segmentSize, indexInRollup)
	if err != nil {
		log.Fatalf("Failed to build new merkle tree: %v", err)
	}
	fmt.Printf("Balance of account %d after withdrawal: %d\n", withdrawalAcc.index, withdrawalAcc.balance.Uint64())

	return newRootHash, oldMerkleProofWithdrawal
}

func (o *Operator) ListenToDepositsWithdrawals(rollup *Rollup, authOwner *bind.TransactOpts) {
	o.H.Reset()
	segmentSize := o.H.Size()
	depositSink := make(chan *RollupDeposit)
	withdrawalSink := make(chan *RollupWithdraw)

	depositSub, err := rollup.WatchDeposit(nil, depositSink, nil) // Listen to all deposit events
	if err != nil {
		log.Fatalf("Failed to watch deposit events: %v", err)
	}
	defer depositSub.Unsubscribe()

	withdrawalSub, err := rollup.WatchWithdraw(nil, withdrawalSink, nil) // Listen to all withdrawal events
	if err != nil {
		log.Fatalf("Failed to watch withdrawal events: %v", err)
	}
	defer withdrawalSub.Unsubscribe()

	for {
		select {
		case err := <-depositSub.Err():
			if err != nil { // Only log if there's an actual error
				log.Fatalf("Error while watching for deposit events: %v", err)
			}
		case depositEvent := <-depositSink:
			// Handle the deposit event
			log.Printf("Deposit received: %s deposited %s wei", depositEvent.User.Hex(), depositEvent.Amount.String())
			merkleRoot, oldMerkleProofDeposit := o.ApplyDepositEvent(depositEvent.User.Hex(), depositEvent.Amount.Uint64(), segmentSize)
			ethAccount, err := o.readAccount(o.AccountMap[o.EthToRollupMap[depositEvent.User.Hex()]])
			depositAcc, err := RollupAccToEth(ethAccount)
			if err != nil {
				log.Fatalf("Failed to fetch eth account during deposit: %v", err)
			}
			merkleProof, err := MerkleProofToEth(oldMerkleProofDeposit)
			if err != nil {
				log.Fatalf("Failed to convert Merkle proof: %v", err)
			}
			_, err = rollup.ProcessDeposit(authOwner, merkleProof.Proof, depositEvent.Amount, depositEvent.User, depositAcc) // big.NewInt(0).SetBytes(merkleRoot)
			if err != nil {
				log.Fatalf("Failed to process deposits: %v", err)
			}
			latestRoot, _ := rollup.GetMerkleRoot(nil)
			merkleRootBigInt := new(big.Int).SetBytes(merkleRoot)
			if merkleRootBigInt.Cmp(latestRoot) == 0 {
				fmt.Printf(Colorbold+ColorGreen+"Successfully updated accounts after deposit: %d\n"+ColorReset, merkleRoot)
			} else {
				fmt.Printf("roots not identical!!! \nlocal root: %d, \ncontract root: %d\n", merkleRootBigInt, latestRoot)
			}

		case err := <-withdrawalSub.Err():
			if err != nil { // Only log if there's an actual error
				log.Fatalf("Error while watching for withdrawal events: %v", err)
			}
		case withdrawalEvent := <-withdrawalSink:
			// Handle the withdrawal event
			log.Printf("Withdrawal made: %s withdrew %s wei", withdrawalEvent.User.Hex(), withdrawalEvent.Amount.String())
			// Process withdrawal event...
			merkleRoot, _ := o.ApplyWithdrawalEvent(withdrawalEvent.User.Hex(), withdrawalEvent.Amount.Uint64(), segmentSize)
			latestRoot, _ := rollup.GetMerkleRoot(nil)
			merkleRootBigInt := new(big.Int).SetBytes(merkleRoot)
			if merkleRootBigInt.Cmp(latestRoot) == 0 {
				fmt.Printf(Colorbold+ColorGreen+"Successfully updated accounts after withdrawal: %d\n"+ColorReset, merkleRoot)
			} else {
				fmt.Printf("Roots not identical! \nlocal root: %d, \ncontract root: %d\n", merkleRootBigInt, latestRoot)
			}

			// Optionally, handle other cases such as system signals for graceful shutdown.
		}
	}
}

// readAccount reads the account located at index i
func (o *Operator) ReadAccount(i uint64) (Account, error) {

	var res Account
	err := Deserialize(&res, o.State[int(i)*SizeAccount:int(i)*SizeAccount+SizeAccount])
	if err != nil {
		return res, err
	}
	return res, nil
}

func RollupAccToEth(_acc Account) (RollupAccount, error) {
	var acc RollupAccount
	publicKey := RollupPublicKey{
		X: _acc.PubKey.A.X.BigInt(new(big.Int)),
		Y: _acc.PubKey.A.Y.BigInt(new(big.Int)),
	}

	acc.Index = uint64(_acc.index)
	acc.Nonce = new(big.Int).SetUint64(_acc.Nonce)
	acc.Balance = new(big.Int).SetUint64(_acc.balance.Uint64())
	acc.PubKey = publicKey

	return acc, nil
}

// ################# old ##########

// ---------- RollupGen ---------
// RollupGen represents a RollupGen, NOTE: not used yet
type RollupGen struct {
	// contract stuff
	op       Operator
	userKeys []eddsa.PrivateKey // list of private keys
	// Mapping of public keys to RollupGen public keys
	contractBalance int64
	rootHash        []byte

	// Network stuff
	networkURL      string         // network URL, e.g.: "HTTP://127.0.0.1:7545"
	contractAddress common.Address // Ethereum Address of the RollupGen contract
}

func NewRollupGen(nbAccounts int) RollupGen {

	var _RollupGen RollupGen
	_op, _keys := GenOperator(nbAccounts)
	_RollupGen.op = _op
	_RollupGen.userKeys = _keys

	_RollupGen.networkURL = "HTTP://127.0.0.1:7545"

	_RollupGen.contractBalance = 0
	_rootHash := "0xab178956422ef"
	_RollupGen.rootHash = []byte(_rootHash)

	// Read contract address from migration artefact json
	RollupGen_json, err := os.ReadFile("build/contracts/RollupGen.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v\n", err)
	}
	json_contract_address := gjson.Get(string(RollupGen_json), "networks.5777.address")
	contractAddress := common.HexToAddress(json_contract_address.String())
	_RollupGen.contractAddress = contractAddress

	return _RollupGen
}

func (o *Operator) MakeTransfer(amount uint64, fromPriv eddsa.PrivateKey, from, to Account) error {

	// create the transfer and sign it
	transfer := NewTransfer(amount, from.PubKey, to.PubKey, from.Nonce)
	transfer.Sign(fromPriv, o.H)

	err := o.UpdateState(transfer, 0)
	if err != nil {
		return err
	}

	return nil
}

// returns a single account on the rollup
func genAccount(i int) (Account, eddsa.PrivateKey) {

	var acc Account
	var rnd fr.Element
	var privkey eddsa.PrivateKey

	// create account, the i-th account has a balance of 20+i
	acc.index = uint64(i)
	acc.Nonce = uint64(i)
	acc.balance.SetUint64(uint64(i) + 20)
	rnd.SetUint64(uint64(i))
	src := rand.NewSource(int64(i))
	r := rand.New(src)

	pkey, err := eddsa.GenerateKey(r)
	if err != nil {
		panic(err)
	}
	privkey = *pkey

	acc.PubKey = privkey.PublicKey

	return acc, privkey
}

// Returns a newly created operator and the private keys of the associated accounts
func GenOperator(nbAccounts int) (Operator, []eddsa.PrivateKey) {

	operator := NewOperator(nbAccounts)

	operator.Witnesses.allocateSlicesMerkleProofs()

	userAccounts := make([]eddsa.PrivateKey, nbAccounts)

	// randomly fill the accounts
	for i := 0; i < nbAccounts; i++ {

		acc, privkey := genAccount(i)

		// fill the index map of the operator
		b := acc.PubKey.A.X.Bytes()
		operator.AccountMap[string(b[:])] = acc.index

		// fill user accounts list
		userAccounts[i] = privkey
		baccount := acc.Serialize()

		copy(operator.State[SizeAccount*i:], baccount)

		// create the list of hashes of account
		operator.H.Reset()
		operator.H.Write(acc.Serialize())
		buf := operator.H.Sum([]byte{})
		copy(operator.HashState[operator.H.Size()*i:], buf)
	}

	return operator, userAccounts

}

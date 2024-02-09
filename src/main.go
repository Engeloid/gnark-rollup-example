package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	. "rollup/src/gnark_rollup"
	. "rollup/src/tools"
	"time"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc"

	//"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tidwall/gjson"
)

const (
	// fileKeys      = "./test/keys.json"
	fileKeys      = "./test/keys2.json"
	fileToSolFile = "./contracts/gnark_verifier.sol"
	NumOfAcc      = 16
)

type GanachAccount struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

func main() {
	var proof groth16.Proof
	var err error

	// Connect to your Ethereum node (Ganache instance or other network)
	client, err := ethclient.Dial("ws://127.0.0.1:7545") // Update with your Ganache URL
	if err != nil {
		log.Fatalf("Failed to connect to ethclient/ganache: %v", err)
	}
	defer client.Close()
	// Read private/pub keys of ganache accounts
	file, err := os.ReadFile(fileKeys)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v\n", err)
	}
	var accounts []GanachAccount
	err = json.Unmarshal([]byte(file), &accounts)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v\n", err)
	}

	// Create contract instance
	rollup_contract_json, err := os.ReadFile("build/contracts/Rollup.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v\n", err)
	}
	json_rollup_contract_address := gjson.Get(string(rollup_contract_json), "networks.5777.address")
	rollupContractAddress := common.HexToAddress(json_rollup_contract_address.String())
	rollupContract, err := NewRollup(rollupContractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate Rollup contract: %v", err)
	}

	// Create operator with a State of 16 accounts
	op := NewOperator(NumOfAcc)
	// If one file does not exist, create constraint system, new verifier contract etc.
	if FileExists(FileToRVK) || FileExists(FileToRPK) || FileExists(FileToRCS) || FileExists(FileToZKVerifier) {
		// Otherwise load them from local repo
		_, _, _, err = op.SetupZKVerifierContract()
		if err != nil {
			log.Fatalf("Could not create zk verifier contract file: %v", err)
		}
	}

	//Create initial merkle tree
	var buf bytes.Buffer
	_, err = buf.Write(op.HashState)
	if err != nil {
		log.Fatalf("failed to read hash state of operator: %v", err)
	}
	merkleRootInit, _, _, err := merkletree.BuildReaderProof(&buf, op.H, op.H.Size(), 0)
	if err != nil {
		log.Fatalf("Failed to init merkle tree: %v", err)
	}
	fmt.Printf("Init merkle root: %d\n", big.NewInt(0).SetBytes(merkleRootInit))

	// -------- on-chain stuff -----------
	owner, err := crypto.HexToECDSA(accounts[0].PrivateKey[2:])
	if err != nil {
		log.Fatalf("Failed to transform owner key to ECDSA: %v", err)
	}
	authOwner, err := bind.NewKeyedTransactorWithChainID(owner, big.NewInt(1337)) // Use the correct chain ID for Ganache
	if err != nil {
		log.Fatalf("Failed to create owner transactor: %v", err)
	}

	// Start listening to deposit and withdrawal events in a new goroutine
	go op.ListenToDepositsWithdrawals(rollupContract, authOwner)

	// ------------------ Deposit and withdraw tests -------------
	senderPriv, err := crypto.HexToECDSA(accounts[1].PrivateKey[2:])
	if err != nil {
		log.Fatalf("Failed to transform to ECDSA: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(senderPriv, big.NewInt(1337)) // Use the correct chain ID for Ganache
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	// Specify the amount to deposit
	depositAmount := big.NewInt(1000000000000000) // 0.001 ETH in Wei big.NewInt(1000000000000000000) // 1 ETH in Wei
	// Set up the transact opts
	auth.Value = depositAmount      // Value in wei to send
	auth.GasLimit = uint64(6721975) // Set the gas limit
	_, err = rollupContract.Deposit(auth)
	if err != nil {
		log.Fatalf("Failed to deposit to rollup contract: %v", err)
	}
	// Second deposit to first acount
	time.Sleep(5 * time.Second)
	_, err = rollupContract.Deposit(auth)
	if err != nil {
		log.Fatalf("Failed to deposit to rollup contract: %v", err)
	}

	// Second accout on rollup
	time.Sleep(5 * time.Second)
	senderPriv2, err := crypto.HexToECDSA(accounts[3].PrivateKey[2:])
	if err != nil {
		log.Fatalf("Failed to transform to ECDSA: %v", err)
	}
	auth2, err := bind.NewKeyedTransactorWithChainID(senderPriv2, big.NewInt(1337)) // Use the correct chain ID for Ganache
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	// Set up the transact opts
	auth2.Value = depositAmount      // Value in wei to send
	auth2.GasLimit = uint64(6721975) // Set the gas limit
	_, err = rollupContract.Deposit(auth2)
	if err != nil {
		log.Fatalf("Failed to deposit to rollup contract: %v", err)
	}

	time.Sleep(5 * time.Second)
	// Fetch and print contract balance
	contractBalance, err := client.BalanceAt(context.Background(), rollupContractAddress, nil)
	if err != nil {
		log.Fatalf("Failed to fetch contract balance: %v", err)
	}
	fmt.Println("Contract balance:", contractBalance)

	var shift uint64 = 256
	var transferAmount uint64 = 500000000

	// get info on both parties of the transaction
	sender, err := op.ReadAccount(0)
	if err != nil {
		log.Fatalf("Error reading sender account: %v", err)
	}

	receiver, err := op.ReadAccount(1)
	if err != nil {
		log.Fatalf("Error reading receiver account: %v", err)
	}

	//----------- Transactions -------

	// Start listening/processing transactions
	op.StartProcessingTxs()
	var hashF = mimc.NewMiMC()

	// Simulate adding transactions
	for i := 0; i < BatchSizeCircuit; i++ {
		tx := NewTransfer(transferAmount, sender.PubKey, receiver.PubKey, sender.Nonce)
		tx.Sign(op.Keys[0], op.H)
		hashF.Reset()
		op.AddTransaction(tx)

		time.Sleep(1 * time.Second)
		// Read account from operator state to get most recent nonce
		sender, err = op.ReadAccount(0)
		if err != nil {
			log.Fatalf("Error reading sender account: %v", err)
		}
	}

	fmt.Println("----------- After tx loop ----------")
	fmt.Printf("Balance of 1st account: %d\n", big.NewInt(0).SetBytes(op.State[64:96]))
	fmt.Printf("Balance of 2nd account: %d\n", big.NewInt(0).SetBytes(op.State[shift-32:shift]))

	// Submit zkProof and root to contract
	proof, err = op.CreateZKProof()
	if err != nil {
		log.Fatalf("Failed to generate zkProof: %v", err)
	}
	eth_proof, _ := ProofToEthereumProof(proof)
	// Transform data to fit contract call
	// TODO, get root from state
	input := new(big.Int).SetBytes(op.RootHashA)

	_, err = rollupContract.SubmitVerifyBatch(authOwner, eth_proof.Proof, input)
	if err != nil {
		log.Fatalf("Failed to submit batch: %v", err)
	}
	fmt.Printf(Colorbold+ColorGreen+"zkProof sucessful on-chain: %d\n"+ColorReset, op.RootHashA)

	// Withdrawal from contract
	time.Sleep(5 * time.Second)
	withdrawAccount, err := op.ReadAccount(1)
	if err != nil {
		log.Fatalf("Failed to fetch withdraw account: %v", err)
	}
	withdrawAccountEth, err := RollupAccToEth(withdrawAccount)
	b := withdrawAccount.PubKey.A.X.Bytes()
	withdrawAccountIndex := op.AccountMap[string(b[:])]

	buf.Reset() // the buffer needs to be reset
	_, err = buf.Write(op.HashState)
	if err != nil {
		log.Fatalf("failed to read hash state of operator: %v", err)
	}
	_, withdrawProof, _, err := merkletree.BuildReaderProof(&buf, op.H, op.H.Size(), withdrawAccountIndex)
	if err != nil {
		log.Fatalf("Failed to create Merkle proof during withdrawal: %v", err)
	}

	withdrawEthProof, _ := MerkleProofToEth(withdrawProof) // Merkle proof for contract call
	auth.Value = big.NewInt(0)                             // Reset to 0 as we're not sending ETH for a withdrawal
	auth2.Value = big.NewInt(0)                            // Reset to 0 as we're not sending ETH for a withdrawal                        // Adjust the Gas Limit as per the withdrawal transaction requirement
	withdrawAmount := big.NewInt(1000000002000000)         // 0.001 ETH in Wei, adjust as needed
	_, err = rollupContract.Withdraw(auth2, withdrawAmount, withdrawEthProof.Proof, withdrawAccountEth)
	if err != nil {
		log.Fatalf("Failed to withdraw: %v", err)
	}

	time.Sleep(2 * time.Second)
	// Fetch and print contract balance
	contractBalance, err = client.BalanceAt(context.Background(), rollupContractAddress, nil)
	if err != nil {
		log.Fatalf("Failed to fetch contract balance: %v", err)
	}

	fmt.Println("Contract balance after withdrawal:", contractBalance)

	// Compare latest merkle root
	time.Sleep(2 * time.Second)
	buf.Reset()
	_, err = buf.Write(op.HashState)
	if err != nil {
		log.Fatalf("failed to read hash state of operator: %v", err)
	}
	latestLocalRoot, _, _, err := merkletree.BuildReaderProof(&buf, op.H, op.H.Size(), 0)
	if err != nil {
		log.Fatalf("Failed to create merkle proof: %v", err)
	}
	latestLocalBigRoot := new(big.Int).SetBytes(latestLocalRoot)
	latestRoot, _ := rollupContract.GetMerkleRoot(nil)
	if latestLocalBigRoot.Cmp(latestRoot) == 0 {
		fmt.Printf(Colorbold+ColorGreen+"Successfully updated accounts after withdrawal: %d\n"+ColorReset, latestLocalBigRoot)
	} else {
		fmt.Printf("Roots not identical! \nlocal root: %d, \ncontract root: %d\n", latestLocalBigRoot, latestRoot)
	}

	os.Exit(3)

	//######################## Old Tests #########################

	// ------------------ Circuit stuff ----------------------

	var circuit Circuit
	for i := 0; i < BatchSizeCircuit; i++ {
		// allocating slice for the Merkle paths
		circuit.MerkleProofReceiverBefore[i].Path = make([]frontend.Variable, 5)
		circuit.MerkleProofReceiverAfter[i].Path = make([]frontend.Variable, 5)
		circuit.MerkleProofSenderBefore[i].Path = make([]frontend.Variable, 5)
		circuit.MerkleProofSenderAfter[i].Path = make([]frontend.Variable, 5)
	}
	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		log.Fatalln("Could not compile circuit:", err)
	}

	// groth16 zkSNARK: Setup
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		log.Fatalln("Could not setup pk and vk:", err)
	}

	// create witness and proof
	witness, err := frontend.NewWitness(&op.Witnesses, ecc.BN254.ScalarField())
	// witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		log.Fatalf("Could not create witnesses: %v", err)
	}
	// fmt.Println("witness created: ", witness)
	publicWitness, err := witness.Public()
	if err != nil {
		log.Fatalf("Could not create public witness: %v", err)
	}

	// generate the proof
	proof, err = groth16.Prove(r1cs, pk, witness)
	if err != nil {
		log.Fatalf("Could not create proof using pk and witnesses: %v", err)
	}

	// verify the proof
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		log.Fatalf("invalid proof: %v", err)
	}

	var useGenProof bool = false
	if !FileExists(FileToRProof) || !FileExists(FileToRPK) || !FileExists(FileToRCS) {
		fmt.Println("Proof files will be created")
		// Setup and generate files (solidity contract, pk, etc.)
		err := SaveProofToFile(proof, FileToRProof)
		if err != nil {
			log.Fatalln(err)
		}
		err = SavePKToFile(pk, FileToRPK)
		if err != nil {
			log.Fatalln(err)
		}
		err = SaveCSToFile(r1cs, FileToRCS)
		if err != nil {
			log.Fatalln(err)
		}

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
		if err != nil {
			log.Fatalf("Error setting up proof and contract files: %v", err)
		}
	} else if useGenProof {
		// using generated proof
		fmt.Println("Using Proof just generated")
	} else {
		// Load proof
		proof, err = LoadProofFromFile(FileToRProof)
		if err != nil {
			log.Fatalf("Error loading proof: %v", err)
		} else {
			fmt.Println("Proof loaded from", FileToRProof)
			// Now you can use the 'proof' variable in the code.
		}
	}

	eth_proofs, _ := ProofToEthereumProof(proof)

	// Transform data to fit contract call
	input_value := [2]*big.Int{}
	input_value[0] = new(big.Int).SetBytes(op.RootHashB)
	input_value[1] = new(big.Int).SetBytes(op.RootHashA)

	// sender.PubKey.A.X.BigInt(new(big.Int))
	// input_value[3] = new(big.Int).SetUint64(35)
	fmt.Printf("Input Root hash before = %x\n", input_value[0])
	fmt.Printf("Input Root hash before as digits = %d\n", input_value[0])
	fmt.Printf("Input Root hash after = %x\n", input_value[1])

	// ---------------- Truffle/Ganache contract call --------

	// ---------- Gnark verification ---------
	// Read contract address from migration artefact json
	/* verifier_json, err := os.ReadFile("build/contracts/Verifier.json")
	if err != nil {
		log.Fatalf("Error reading JSON file: %v\n", err)
	}
	json_contract_address := gjson.Get(string(verifier_json), "networks.5777.address")
	contractAddress := common.HexToAddress(json_contract_address.String())

	// Create a contract instance
	verifier, err := NewVerifier(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate Verifier contract: %v", err)
	}

	//var result []interface{}
	// err = verifier.VerifyProof(nil, eth_proof.A, eth_proof.B, eth_proof.C, input_value)
	err = verifier.VerifyProof(nil, eth_proofs.Proof, input_value)
	if err != nil {
		log.Fatalf("Failed to validate proof: %v", err)
	}

	fmt.Println(bold + color + "Proof verified on blockchain!!!" + colorReset) */

	// ------------- Rollup contract tests ------------------

	/* 	fmt.Printf("Balance of 1st account: %d\n", operatorA.State[64:96])
	   	fmt.Printf("Balance of 2nd account: %d\n", operatorA.State[shift-32:shift])
	   	fmt.Printf("Nonce of 1st account after update: %d\n", sender.Nonce)
	   	fmt.Printf("Nonce of 1st account after after update of state: %d\n", operatorA.State[32:64]) */

	// r1cs, pk, vk, _ := operatorA.SetupRollupProofAndContract()
	//operatorA.SetupRollupProofAndContract()

	// ---- zk-proof verification on chain ------------
	// ownerPub := common.HexToAddress(accounts[5].PublicKey)

	// get merkleroot after inital set
	getRoot, err := rollupContract.GetMerkleRoot(nil)
	fmt.Printf("Initial root hash in contract = %x\n", getRoot)

	// verify new state
	verifytx, err := rollupContract.SubmitVerifyBatch(authOwner, eth_proofs.Proof, new(big.Int).SetBytes(op.RootHashA))
	if err != nil {
		log.Fatalf("Failed to verify: %v", err)
	}
	fmt.Printf(Colorbold+ColorGreen+"ZK-Proof valid: %s\n"+ColorReset, verifytx.Hash().Hex())

	// get merkle root after state update
	getRootAfter, err := rollupContract.GetMerkleRoot(nil)
	fmt.Printf("Root hash in contract after update = %x\n", getRootAfter)
}

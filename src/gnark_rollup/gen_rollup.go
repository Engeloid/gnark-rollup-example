package rollup

import (
	// "hash"

	"bytes"
	"fmt"
	"log"
	"math/big"
	"math/rand"

	// . "rollup/src/tools"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/ethereum/go-ethereum/common"
)

const (
	fileToProof = "./test/rollupg16.proof"
	fileToPK    = "./test/rollupg16.pk"
	fileToCS    = "./test/rollupg16.cs"
)

// ---------- Rollup ---------
// Rollup represents a rollup, NOTE: not used yet
type Rollup struct {
	// contract stuff
	op       Operator
	userKeys []eddsa.PrivateKey // list of private keys
	// Mapping of public keys to rollup public keys
	contractBalance int64

	// Network stuff
	networkURL      string         // network URL, e.g.: "HTTP://127.0.0.1:7545"
	contractAddress common.Address // Ethereum Address of the rollup contract
}

// func SetupRollupProofAndContract(fileNameProof string, fileNamePK string, fileNameCS string) (groth16.Proof, error) {
func (o *Operator) SetupRollupProofAndContract() (constraint.ConstraintSystem, groth16.ProvingKey, groth16.VerifyingKey, error) {
	// compiles our circuit into a R1CS
	var circuit Circuit

	_ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit, frontend.IgnoreUnconstrainedInputs())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Circuit sender key: %x\n", circuit.PublicKeysSender[0].A.X)
	//fmt.Printf("constraint system: %v\n", ccs)
	// err := SaveCSToFile(ccs, fileToCS)

	// groth16 zkSNARK: Setup
	_pk, _vk, err := groth16.Setup(_ccs)
	if err != nil {
		log.Fatalln(err)
	}
	// err = SavePKToFile(pk, fileToPK)

	// -------- witness definition ---------

	return _ccs, _pk, _vk, err
}

func CreateOP(nbAccounts int) (Operator, []eddsa.PrivateKey) {
	// TODO: create Rollup object and create respective smart contract
	op, userKeys := genOperator(nbAccounts)

	return op, userKeys
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

// ReadAcc reads the account located at index i
func (o *Operator) ReadAcc(i uint64) (Account, error) {
	var res Account
	err := Deserialize(&res, o.State[int(i)*SizeAccount:int(i)*SizeAccount+SizeAccount])
	if err != nil {
		return res, err
	}
	return res, nil
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
func genOperator(nbAccounts int) (Operator, []eddsa.PrivateKey) {

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

func (o *Operator) ReadAccounts() ([]Account, error) {
	var i uint64
	res := make([]Account, o.nbAccounts)
	for i = 0; i < uint64(o.nbAccounts); i++ {
		res[i], _ = o.readAccount(i)
	}

	return res, nil
}

func (o *Operator) GetNbAccounts() uint64 {

	return uint64(o.nbAccounts)
}

// updateAccount updates the state according to transfer
// numTransfer is the number of the transfer currently handled (between 0 and BatchSizeCircuit)
func (o *Operator) RollupUpdateState(t Transfer, numTransfer int) error {

	var posSender, posReceiver uint64
	var ok bool

	// ext := strconv.Itoa(numTransfer)
	segmentSize := o.H.Size()

	// read sender's account
	b := t.senderPubKey.A.X.Bytes()
	if posSender, ok = o.AccountMap[string(b[:])]; !ok {
		return ErrNonExistingAccount
	}
	senderAccount, err := o.readAccount(posSender)
	if err != nil {
		return err
	}
	if senderAccount.index != posSender {
		return ErrIndexConsistency
	}

	// read receiver's account
	b = t.receiverPubKey.A.X.Bytes()
	if posReceiver, ok = o.AccountMap[string(b[:])]; !ok {
		return ErrNonExistingAccount
	}
	receiverAccount, err := o.readAccount(posReceiver)
	if err != nil {
		return err
	}
	if receiverAccount.index != posReceiver {
		return ErrIndexConsistency
	}

	// set witnesses for the leaves
	o.Witnesses.LeafReceiver[numTransfer] = posReceiver
	o.Witnesses.LeafSender[numTransfer] = posSender

	// set witnesses for the public keys
	o.Witnesses.PublicKeysSender[numTransfer].A.X = senderAccount.PubKey.A.X
	o.Witnesses.PublicKeysSender[numTransfer].A.Y = senderAccount.PubKey.A.Y
	o.Witnesses.PublicKeysReceiver[numTransfer].A.X = receiverAccount.PubKey.A.X
	o.Witnesses.PublicKeysReceiver[numTransfer].A.Y = receiverAccount.PubKey.A.Y

	// set witnesses for the accounts before update
	o.Witnesses.SenderAccountsBefore[numTransfer].Index = senderAccount.index
	o.Witnesses.SenderAccountsBefore[numTransfer].Nonce = senderAccount.Nonce
	o.Witnesses.SenderAccountsBefore[numTransfer].Balance = senderAccount.balance

	o.Witnesses.ReceiverAccountsBefore[numTransfer].Index = receiverAccount.index
	o.Witnesses.ReceiverAccountsBefore[numTransfer].Nonce = receiverAccount.Nonce
	o.Witnesses.ReceiverAccountsBefore[numTransfer].Balance = receiverAccount.balance

	//  Set witnesses for the proof of inclusion of sender and receivers account before update
	var buf bytes.Buffer
	_, err = buf.Write(o.HashState)
	if err != nil {
		return err
	}
	merkleRootBefore, proofInclusionSenderBefore, numLeaves, err := merkletree.BuildReaderProof(&buf, o.H, segmentSize, posSender)
	if err != nil {
		return err
	}

	// verify the proof in plain go...
	merkletree.VerifyProof(o.H, merkleRootBefore, proofInclusionSenderBefore, posSender, numLeaves)

	buf.Reset() // the buffer needs to be reset
	_, err = buf.Write(o.HashState)
	if err != nil {
		return err
	}
	_, proofInclusionReceiverBefore, _, err := merkletree.BuildReaderProof(&buf, o.H, segmentSize, posReceiver)
	if err != nil {
		return err
	}
	o.Witnesses.RootHashesBefore[numTransfer] = merkleRootBefore
	o.Witnesses.MerkleProofReceiverBefore[numTransfer].RootHash = merkleRootBefore
	o.Witnesses.MerkleProofSenderBefore[numTransfer].RootHash = merkleRootBefore

	for i := 0; i < len(proofInclusionSenderBefore); i++ {
		o.Witnesses.MerkleProofReceiverBefore[numTransfer].Path[i] = proofInclusionReceiverBefore[i]
		o.Witnesses.MerkleProofSenderBefore[numTransfer].Path[i] = proofInclusionSenderBefore[i]
	}

	// set witnesses for the transfer
	o.Witnesses.Transfers[numTransfer].Amount = t.amount
	o.Witnesses.Transfers[numTransfer].Signature.R.X = t.signature.R.X
	o.Witnesses.Transfers[numTransfer].Signature.R.Y = t.signature.R.Y
	o.Witnesses.Transfers[numTransfer].Signature.S = t.signature.S[:]

	// verifying the signature. The msg is the hash (o.h) of the transfer
	// nonce ∥ amount ∥ senderpubKey(x&y) ∥ receiverPubkey(x&y)
	resSig, err := t.Verify(o.H)
	if err != nil {
		return err
	}
	if !resSig {
		return ErrWrongSignature
	}

	// checks if the amount is correct
	var bAmount, bBalance big.Int
	receiverAccount.balance.BigInt(&bBalance)
	t.amount.BigInt(&bAmount)
	if bAmount.Cmp(&bBalance) == 1 {
		return ErrAmountTooHigh
	}

	// check if the nonce is correct
	if t.nonce != senderAccount.Nonce {
		return ErrNonce
	}

	// update balances
	senderAccount.balance.Sub(&senderAccount.balance, &t.amount)
	receiverAccount.balance.Add(&receiverAccount.balance, &t.amount)

	// update the nonce of the sender
	senderAccount.Nonce++

	// set the witnesses for the account after update
	o.Witnesses.ReceiverAccountsAfter[numTransfer].Index = receiverAccount.index
	o.Witnesses.ReceiverAccountsAfter[numTransfer].Nonce = receiverAccount.Nonce
	o.Witnesses.ReceiverAccountsAfter[numTransfer].Balance = receiverAccount.balance

	o.Witnesses.SenderAccountsAfter[numTransfer].Index = senderAccount.index
	o.Witnesses.SenderAccountsAfter[numTransfer].Nonce = senderAccount.Nonce
	o.Witnesses.SenderAccountsAfter[numTransfer].Balance = senderAccount.balance

	// update the state of the operator
	copy(o.State[int(posSender)*SizeAccount:], senderAccount.Serialize())
	o.H.Reset()
	_, _ = o.H.Write(senderAccount.Serialize())
	bufSender := o.H.Sum([]byte{})
	copy(o.HashState[int(posSender)*o.H.Size():(int(posSender)+1)*o.H.Size()], bufSender)

	copy(o.State[int(posReceiver)*SizeAccount:], receiverAccount.Serialize())
	o.H.Reset()
	_, _ = o.H.Write(receiverAccount.Serialize())
	bufReceiver := o.H.Sum([]byte{})
	copy(o.HashState[int(posReceiver)*o.H.Size():(int(posReceiver)+1)*o.H.Size()], bufReceiver)

	//  Set witnesses for the proof of inclusion of sender and receivers account after update
	buf.Reset()
	_, err = buf.Write(o.HashState)
	if err != nil {
		return err
	}
	merkleRootAfer, proofInclusionSenderAfter, _, err := merkletree.BuildReaderProof(&buf, o.H, segmentSize, posSender)
	if err != nil {
		return err
	}
	// merkleProofHelperSenderAfter := merkle.GenerateProofHelper(proofInclusionSenderAfter, posSender, numLeaves)

	buf.Reset() // the buffer needs to be reset
	_, err = buf.Write(o.HashState)
	if err != nil {
		return err
	}
	_, proofInclusionReceiverAfter, _, err := merkletree.BuildReaderProof(&buf, o.H, segmentSize, posReceiver)
	if err != nil {
		return err
	}
	// merkleProofHelperReceiverAfter := merkle.GenerateProofHelper(proofInclusionReceiverAfter, posReceiver, numLeaves)

	o.Witnesses.RootHashesAfter[numTransfer] = merkleRootAfer
	o.Witnesses.MerkleProofReceiverAfter[numTransfer].RootHash = merkleRootAfer
	o.Witnesses.MerkleProofSenderAfter[numTransfer].RootHash = merkleRootAfer

	for i := 0; i < len(proofInclusionSenderAfter); i++ {
		o.Witnesses.MerkleProofReceiverAfter[numTransfer].Path[i] = proofInclusionReceiverAfter[i]
		o.Witnesses.MerkleProofSenderAfter[numTransfer].Path[i] = proofInclusionSenderAfter[i]
	}

	return nil
}

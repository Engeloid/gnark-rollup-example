/*
Copyright © 2020 ConsenSys

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rollup

import (
	"bytes"
	"hash"
	"math/big"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
)

var hFunc = mimc.NewMiMC()

// BatchSize size of a batch of transactions to put in a snark
var BatchSize = BatchSizeCircuit

// Queue queue for storing the transfers (fixed size queue)
type Queue struct {
	listTransfers chan Transfer
}

// NewQueue creates a new queue, BatchSizeCircuit is the capaciy
func NewQueue(BatchSizeCircuit int) Queue {
	resChan := make(chan Transfer, BatchSizeCircuit)
	var res Queue
	res.listTransfers = resChan
	return res
}

// Operator represents a rollup operator
type Operator struct {
	State            []byte            // list of accounts: index ∥ nonce ∥ balance ∥ pubkeyX ∥ pubkeyY, each chunk is 256 bits
	HashState        []byte            // Hashed version of the state, each chunk is 256bits: ... ∥ H(index ∥ nonce ∥ balance ∥ pubkeyX ∥ pubkeyY)) ∥ ...
	AccountMap       map[string]uint64 // hashmap of all available accounts (the key is the account.pubkey.X), the value is the index of the account in the state
	nbAccounts       int               // number of accounts managed by this operator
	H                hash.Hash         // hash function used to build the Merkle Tree
	q                Queue             // queue of transfers
	batch            int               // current number of transactions in a batch
	Witnesses        Circuit           // witnesses for the snark cicruit
	EthToRollupMap   map[string]string // EthToRollupMap maps Ethereum addresses (as hex strings) to rollup addresses (pubkey.x as strings).
	AvailableIndices []uint64          // List of available indices
	NextIndex        uint64            // Next new index to be used if no available indices
	BatchID          uint64            // Id of batch currently working with
	// Only for testing/simulation:
	Keys      []eddsa.PrivateKey
	RootHashB []byte
	RootHashA []byte
}

// NewOperator creates an operator with all being set to zero.
// nbAccounts is the number of accounts managed by this operator, h is the hash function for the merkle proofs
// TODO: adjust function so that compressed transactions can be taken in to reconstruct state for new operators
func NewOperator(nbAccounts int) Operator {

	res := Operator{}

	// create a list of empty accounts
	res.State = make([]byte, SizeAccount*nbAccounts)

	// initialize hash of the state
	res.HashState = make([]byte, hFunc.Size()*nbAccounts)
	for i := 0; i < nbAccounts; i++ {
		hFunc.Reset()
		_, _ = hFunc.Write(res.State[i*SizeAccount : i*SizeAccount+SizeAccount])
		s := hFunc.Sum([]byte{})
		// fmt.Printf("ZERO_VALUE = %d\n", big.NewInt(0).SetBytes(s))
		copy(res.HashState[i*hFunc.Size():(i+1)*hFunc.Size()], s)
	}

	res.AccountMap = make(map[string]uint64)
	res.nbAccounts = nbAccounts
	res.H = hFunc
	res.q = NewQueue(BatchSize)
	res.batch = 0
	res.EthToRollupMap = make(map[string]string)
	res.AvailableIndices = []uint64{}
	res.NextIndex = 0
	res.BatchID = 0

	res.Keys = []eddsa.PrivateKey{}
	res.Witnesses.allocateSlicesMerkleProofs()
	return res
}

// readAccount reads the account located at index i
func (o *Operator) readAccount(i uint64) (Account, error) {

	var res Account
	err := Deserialize(&res, o.State[int(i)*SizeAccount:int(i)*SizeAccount+SizeAccount])
	if err != nil {
		return res, err
	}
	return res, nil
}

// updateAccount updates the state according to transfer
// numTransfer is the number of the transfer currently handled (between 0 and BatchSizeCircuit)
func (o *Operator) UpdateState(t Transfer, numTransfer int) error {

	var posSender, posReceiver uint64
	var ok bool
	o.H.Reset()

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
	o.RootHashB = merkleRootBefore
	o.Witnesses.RootHashesBefore[numTransfer] = merkleRootBefore
	o.Witnesses.MerkleProofReceiverBefore[numTransfer].RootHash = merkleRootBefore
	o.Witnesses.MerkleProofSenderBefore[numTransfer].RootHash = merkleRootBefore

	for i := 0; i < len(proofInclusionSenderBefore); i++ {
		o.Witnesses.MerkleProofReceiverBefore[numTransfer].Path[i] = proofInclusionReceiverBefore[i]
		o.Witnesses.MerkleProofSenderBefore[numTransfer].Path[i] = proofInclusionSenderBefore[i]
	}

	// set witnesses for the transfer
	o.Witnesses.Transfers[numTransfer].Amount = t.amount //.BigInt(&big.Int{})
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

	//TODO remove next line
	o.RootHashA = merkleRootAfer

	o.Witnesses.RootHashesAfter[numTransfer] = merkleRootAfer
	o.Witnesses.MerkleProofReceiverAfter[numTransfer].RootHash = merkleRootAfer
	o.Witnesses.MerkleProofSenderAfter[numTransfer].RootHash = merkleRootAfer

	for i := 0; i < len(proofInclusionSenderAfter); i++ {
		o.Witnesses.MerkleProofReceiverAfter[numTransfer].Path[i] = proofInclusionReceiverAfter[i]
		o.Witnesses.MerkleProofSenderAfter[numTransfer].Path[i] = proofInclusionSenderAfter[i]
	}

	return nil
}

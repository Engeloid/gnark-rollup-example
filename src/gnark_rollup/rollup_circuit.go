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
	tedwards "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	"github.com/consensys/gnark/std/algebra/twistededwards"
	"github.com/consensys/gnark/std/hash/mimc"
	"github.com/consensys/gnark/std/signature/eddsa"
)

// Circuit "toy" rollup circuit where an operator can generate a proof that he processed
// some transactions
type RollupCircuit struct {
	// ---------------------------------------------------------------------------------------------
	// SECRET INPUTS

	// list of accounts involved before update and their public keys
	SenderAccountsBefore   RollupAccountConstraints
	ReceiverAccountsBefore RollupAccountConstraints

	// list of accounts involved after update and their public keys
	SenderAccountsAfter   RollupAccountConstraints
	ReceiverAccountsAfter RollupAccountConstraints

	// list of transactions
	Amount         frontend.Variable
	Nonce          frontend.Variable `gnark:"-"`
	SenderPubKey   eddsa.PublicKey
	ReceiverPubKey eddsa.PublicKey
	Signature      eddsa.Signature

	// list of proofs corresponding to sender and receiver accounts
	MerkleProofReceiverBefore merkle.MerkleProof
	MerkleProofReceiverAfter  merkle.MerkleProof
	MerkleProofSenderBefore   merkle.MerkleProof
	MerkleProofSenderAfter    merkle.MerkleProof
	LeafReceiver              frontend.Variable
	LeafSender                frontend.Variable

	// ---------------------------------------------------------------------------------------------
	// PUBLIC INPUTS

	// list of root hashes
	RootHashesBefore frontend.Variable `gnark:",public"`
	RootHashesAfter  frontend.Variable `gnark:",public"`
}

// AccountConstraints accounts encoded as constraints
type RollupAccountConstraints struct {
	Index   frontend.Variable // index in the tree
	Nonce   frontend.Variable // nb transactions done so far from this account
	Balance frontend.Variable
	PubKey  eddsa.PublicKey `gnark:"-"`
}

func (circuit *RollupCircuit) postInit2(api frontend.API) error {

	for i := 0; i < BatchSizeCircuit; i++ {

		// setting the sender accounts before update
		circuit.SenderAccountsBefore.PubKey = circuit.SenderPubKey

		// setting the sender accounts after update
		circuit.SenderAccountsAfter.PubKey = circuit.SenderPubKey

		// setting the receiver accounts before update
		circuit.ReceiverAccountsBefore.PubKey = circuit.ReceiverPubKey

		// setting the receiver accounts after update
		circuit.ReceiverAccountsAfter.PubKey = circuit.ReceiverPubKey

		// setting the transfers
		circuit.Nonce = circuit.SenderAccountsBefore.Nonce

		// allocate the slices for the Merkle proofs
		circuit.AllocateSlicesMerkleProofss()

	}
	return nil
}

func (circuit *RollupCircuit) AllocateSlicesMerkleProofss() {

	for i := 0; i < BatchSizeCircuit; i++ {
		// allocating slice for the Merkle paths
		circuit.MerkleProofReceiverBefore.Path = make([]frontend.Variable, depth)
		circuit.MerkleProofReceiverAfter.Path = make([]frontend.Variable, depth)
		circuit.MerkleProofSenderBefore.Path = make([]frontend.Variable, depth)
		circuit.MerkleProofSenderAfter.Path = make([]frontend.Variable, depth)
	}

}

// Define declares the circuit's constraints
func (circuit *RollupCircuit) Define(api frontend.API) error {

	if err := circuit.postInit2(api); err != nil {
		return err
	}
	// hash function for the merkle proof and the eddsa signature
	hFunc, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}

	// verifications of:
	// - Merkle proofs of the accounts
	// - the signatures
	// - accounts' balance consistency
	for i := 0; i < BatchSizeCircuit; i++ {

		// the root hashes of the Merkle path must match the public ones given in the circuit
		api.AssertIsEqual(circuit.RootHashesBefore, circuit.MerkleProofReceiverBefore.RootHash)
		api.AssertIsEqual(circuit.RootHashesBefore, circuit.MerkleProofSenderBefore.RootHash)
		api.AssertIsEqual(circuit.RootHashesAfter, circuit.MerkleProofReceiverAfter.RootHash)
		api.AssertIsEqual(circuit.RootHashesAfter, circuit.MerkleProofSenderAfter.RootHash)

		// the leafs of the Merkle proofs must match the index of the accounts
		api.AssertIsEqual(circuit.ReceiverAccountsBefore.Index, circuit.LeafReceiver)
		api.AssertIsEqual(circuit.ReceiverAccountsAfter.Index, circuit.LeafReceiver)
		api.AssertIsEqual(circuit.SenderAccountsBefore.Index, circuit.LeafSender)
		api.AssertIsEqual(circuit.SenderAccountsAfter.Index, circuit.LeafSender)

		// Reset the hash state!
		hFunc.Reset()

		// the signature is on h(nonce ∥ amount ∥ senderpubKey (x&y) ∥ receiverPubkey(x&y))
		hFunc.Write(circuit.Nonce, circuit.Amount, circuit.SenderPubKey.A.X, circuit.SenderPubKey.A.Y, circuit.ReceiverPubKey.A.X, circuit.ReceiverPubKey.A.Y)
		htransfer := hFunc.Sum()

		curve, err := twistededwards.NewEdCurve(api, tedwards.BN254)
		if err != nil {
			return err
		}

		hFunc.Reset()
		err = eddsa.Verify(curve, circuit.Signature, htransfer, circuit.SenderPubKey, &hFunc)
		if err != nil {
			return err
		}

		// ensure that nonce is correctly updated
		nonceUpdated := api.Add(circuit.SenderAccountsBefore.Nonce, 1)
		api.AssertIsEqual(nonceUpdated, circuit.SenderAccountsAfter.Nonce)
		api.AssertIsEqual(circuit.ReceiverAccountsBefore.Nonce, circuit.ReceiverAccountsAfter.Nonce)

		// ensures that the amount is less than the balance
		api.AssertIsLessOrEqual(circuit.Amount, circuit.SenderAccountsBefore.Balance)

		// ensure that balance is correctly updated
		fromBalanceUpdated := api.Sub(circuit.SenderAccountsBefore.Balance, circuit.Amount)
		api.AssertIsEqual(fromBalanceUpdated, circuit.SenderAccountsAfter.Balance)

		toBalanceUpdated := api.Add(circuit.ReceiverAccountsBefore.Balance, circuit.Amount)
		api.AssertIsEqual(toBalanceUpdated, circuit.ReceiverAccountsAfter.Balance)
	}

	return nil
}

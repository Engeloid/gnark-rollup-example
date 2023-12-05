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
	SenderAccountsBefore   [BatchSizeCircuit]RollupAccountConstraints
	ReceiverAccountsBefore [BatchSizeCircuit]RollupAccountConstraints
	PublicKeysSender       [BatchSizeCircuit]eddsa.PublicKey

	// list of accounts involved after update and their public keys
	SenderAccountsAfter   [BatchSizeCircuit]RollupAccountConstraints
	ReceiverAccountsAfter [BatchSizeCircuit]RollupAccountConstraints
	PublicKeysReceiver    [BatchSizeCircuit]eddsa.PublicKey

	// list of transactions
	Transfers [BatchSizeCircuit]RollupTransferConstraints

	// list of proofs corresponding to sender and receiver accounts

	LeafReceiver [BatchSizeCircuit]frontend.Variable
	LeafSender   [BatchSizeCircuit]frontend.Variable

	// ---------------------------------------------------------------------------------------------
	// PUBLIC INPUTS

	// list of root hashes
	RootHashesBefore [BatchSizeCircuit]frontend.Variable `gnark:",public"`
	RootHashesAfter  [BatchSizeCircuit]frontend.Variable `gnark:",public"`
}

// AccountConstraints accounts encoded as constraints
type RollupAccountConstraints struct {
	Index   frontend.Variable // index in the tree
	Nonce   frontend.Variable // nb transactions done so far from this account
	Balance frontend.Variable
	PubKey  eddsa.PublicKey `gnark:"-"`
}

// TransferConstraints transfer encoded as constraints
type RollupTransferConstraints struct {
	Amount         frontend.Variable
	Nonce          frontend.Variable `gnark:"-"`
	SenderPubKey   eddsa.PublicKey   `gnark:"-"`
	ReceiverPubKey eddsa.PublicKey   `gnark:"-"`
	Signature      eddsa.Signature
}

func (circuit *RollupCircuit) postInit2(api frontend.API) error {

	for i := 0; i < BatchSizeCircuit; i++ {

		// setting the sender accounts before update
		circuit.SenderAccountsBefore[i].PubKey = circuit.PublicKeysSender[i]

		// setting the sender accounts after update
		circuit.SenderAccountsAfter[i].PubKey = circuit.PublicKeysSender[i]

		// setting the receiver accounts before update
		circuit.ReceiverAccountsBefore[i].PubKey = circuit.PublicKeysReceiver[i]

		// setting the receiver accounts after update
		circuit.ReceiverAccountsAfter[i].PubKey = circuit.PublicKeysReceiver[i]

		// setting the transfers
		circuit.Transfers[i].Nonce = circuit.SenderAccountsBefore[i].Nonce
		circuit.Transfers[i].SenderPubKey = circuit.PublicKeysSender[i]
		circuit.Transfers[i].ReceiverPubKey = circuit.PublicKeysReceiver[i]

		// allocate the slices for the Merkle proofs
		// circuit.allocateSlicesMerkleProofs()

	}
	return nil
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

		// the leafs of the Merkle proofs must match the index of the accounts
		api.AssertIsEqual(circuit.ReceiverAccountsBefore[i].Index, circuit.LeafReceiver[i])
		api.AssertIsEqual(circuit.ReceiverAccountsAfter[i].Index, circuit.LeafReceiver[i])
		api.AssertIsEqual(circuit.SenderAccountsBefore[i].Index, circuit.LeafSender[i])
		api.AssertIsEqual(circuit.SenderAccountsAfter[i].Index, circuit.LeafSender[i])

		/*
			// verify the transaction transfer
			err := verifyTransferSignature(api, circuit.Transfers[i], hFunc)
			if err != nil {
				return err
			}

			// update the accounts
			verifyAccountUpdated(api, circuit.SenderAccountsBefore[i], circuit.ReceiverAccountsBefore[i], circuit.SenderAccountsAfter[i], circuit.ReceiverAccountsAfter[i], circuit.Transfers[i].Amount)
		*/

		// Reset the hash state!
		hFunc.Reset()

		// the signature is on h(nonce ∥ amount ∥ senderpubKey (x&y) ∥ receiverPubkey(x&y))
		hFunc.Write(circuit.Transfers[i].Nonce, circuit.Transfers[i].Amount, circuit.Transfers[i].SenderPubKey.A.X, circuit.Transfers[i].SenderPubKey.A.Y, circuit.Transfers[i].ReceiverPubKey.A.X, circuit.Transfers[i].ReceiverPubKey.A.Y)
		htransfer := hFunc.Sum()

		curve, err := twistededwards.NewEdCurve(api, tedwards.BN254)
		if err != nil {
			return err
		}

		hFunc.Reset()
		err = eddsa.Verify(curve, circuit.Transfers[i].Signature, htransfer, circuit.Transfers[i].SenderPubKey, &hFunc)
		if err != nil {
			return err
		}
	}

	return nil
}

package main

import (
	"fmt"
	"log"
	. "rollup/src/gnark_rollup"

	// . "rollup/src/tools"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func main() {
	var amount uint64 = 10
	var operatorA, privKeys = CreateOP(10)

	// get info on both parties of the transaction
	sender, err := operatorA.ReadAcc(0)
	if err != nil {
		log.Fatalf("Error reading sender account: %v", err)
	}

	receiver, err := operatorA.ReadAcc(1)
	if err != nil {
		log.Fatalf("Error reading receiver account: %v", err)
	}

	// Transactions

	firstTx := NewTransfer(amount, sender.PubKey, receiver.PubKey, sender.Nonce)
	firstTx.Sign(privKeys[0], operatorA.H)

	// "0" here, as it probably needs to be within a loop iterating over every tx (between 0 and BatchSizeCircuit)
	err = operatorA.UpdateState(firstTx, 0)
	if err != nil {
		log.Fatalf("Error updating operator state: %v", err)
	}

	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &operatorA.Witnesses)
	if err != nil {
		log.Fatalln("Could not compile circuit:", err)
	}
	fmt.Println("Circuit compiled")

	// groth16 zkSNARK: Setup
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		log.Fatalln("Could not setup pk and vk:", err)
	}
	fmt.Println("Setup groth16 rank 1 constraint system")

	// create witness and proof
	witness, err := frontend.NewWitness(&operatorA.Witnesses, ecc.BN254.ScalarField())
	if err != nil {
		log.Fatalf("Could not create witnesses: %v", err)
	}
	fmt.Println("witness created: ", witness)
	publicWitness, err := witness.Public()
	if err != nil {
		log.Fatalf("Could not create public witness: %v", err)
	}
	// generate the proof
	proof, err := groth16.Prove(r1cs, pk, witness)
	if err != nil {
		log.Fatalf("Could not create proof using pk and witnesses: %v", err)
	}
	fmt.Println("Proof created")

	// verify the proof
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		log.Fatalf("invalid proof: %v", err)
	}

	// r1cs, pk, vk, _ := operatorA.SetupRollupProofAndContract()
	//operatorA.SetupRollupProofAndContract()

	// declare and assign Circuit witnesses
	//var assignment Circuit

	/*
		// Transactions
		assignment.Transfers[0].Amount = operatorA.Witnesses.Transfers[0].Amount
		assignment.Transfers[0].Signature = operatorA.Witnesses.Transfers[0].Signature

		// rest of account constraints
		assignment.SenderAccountsBefore[0].Index = operatorA.Witnesses.SenderAccountsBefore[0].Index
		assignment.SenderAccountsBefore[0].Nonce = operatorA.Witnesses.SenderAccountsBefore[0].Nonce
		assignment.SenderAccountsBefore[0].Balance = operatorA.Witnesses.SenderAccountsBefore[0].Balance
		assignment.ReceiverAccountsBefore[0].Index = operatorA.Witnesses.ReceiverAccountsBefore[0].Index
		assignment.ReceiverAccountsBefore[0].Nonce = operatorA.Witnesses.ReceiverAccountsBefore[0].Nonce
		assignment.ReceiverAccountsBefore[0].Balance = operatorA.Witnesses.ReceiverAccountsBefore[0].Balance
		assignment.SenderAccountsAfter[0].Index = operatorA.Witnesses.SenderAccountsAfter[0].Index
		assignment.SenderAccountsAfter[0].Nonce = operatorA.Witnesses.SenderAccountsAfter[0].Nonce
		assignment.SenderAccountsAfter[0].Balance = operatorA.Witnesses.SenderAccountsAfter[0].Balance
		assignment.ReceiverAccountsAfter[0].Index = operatorA.Witnesses.ReceiverAccountsAfter[0].Index
		assignment.ReceiverAccountsAfter[0].Nonce = operatorA.Witnesses.ReceiverAccountsAfter[0].Nonce
		assignment.ReceiverAccountsAfter[0].Balance = operatorA.Witnesses.ReceiverAccountsAfter[0].Balance

		// assign public key values of senders and receivers respectively
		assignment.PublicKeysSender = operatorA.Witnesses.PublicKeysSender
		assignment.PublicKeysReceiver = operatorA.Witnesses.PublicKeysReceiver

		// Merkle proofs, roots and leafs
		assignment.MerkleProofReceiverBefore = operatorA.Witnesses.MerkleProofReceiverBefore
		assignment.MerkleProofReceiverAfter = operatorA.Witnesses.MerkleProofReceiverAfter
		assignment.MerkleProofSenderBefore = operatorA.Witnesses.MerkleProofSenderBefore
		assignment.MerkleProofSenderAfter = operatorA.Witnesses.MerkleProofSenderAfter
		assignment.LeafReceiver = operatorA.Witnesses.LeafReceiver
		assignment.LeafSender = operatorA.Witnesses.LeafSender

		// PUBLIC INPUTS
		assignment.RootHashesBefore = operatorA.Witnesses.RootHashesBefore
		assignment.RootHashesAfter = operatorA.Witnesses.RootHashesAfter

		// create witness and proof
		witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
		publicWitness, err := witness.Public()
		// generate the proof
		proof, err = groth16.Prove(r1cs, pk, witness)

		// verify the proof
		err = groth16.Verify(proof, vk, publicWitness)
		if err != nil {
			log.Fatalf("invalid proof: %v", err)
		}
	*/

	/* // declare and assign the account witnesses
	var senderAccBefore [BatchSizeCircuit]AccountConstraints = operatorA.Witnesses.SenderAccountsBefore
	var receiverAccBefore [BatchSizeCircuit]AccountConstraints = operatorA.Witnesses.ReceiverAccountsBefore
	var senderAccAfter [BatchSizeCircuit]AccountConstraints = operatorA.Witnesses.SenderAccountsAfter
	var receiverAccAfter [BatchSizeCircuit]AccountConstraints = operatorA.Witnesses.ReceiverAccountsAfter

	assignment.SenderAccountsBefore = senderAccBefore
	assignment.ReceiverAccountsBefore = receiverAccBefore
	assignment.SenderAccountsAfter = senderAccAfter
	assignment.ReceiverAccountsAfter = receiverAccAfter

	// assign transaction values
	var transfers [BatchSizeCircuit]TransferConstraints = operatorA.Witnesses.Transfers

	assignment.Transfers[0] = transfers[0]

	// PublicKeys
	var publicKeySender, publicKeyReceiver eddsa.PublicKey

	publicKeySender = operatorA.Witnesses.PublicKeysSender[0]
	publicKeyReceiver = operatorA.Witnesses.PublicKeysReceiver[0]

	assignment.PublicKeysSender[0] = publicKeySender
	assignment.PublicKeysReceiver[0] = publicKeyReceiver

	// Merkle proofs, roots and leafs

	fmt.Println("Merkle path = ", operatorA.Witnesses.MerkleProofReceiverBefore)
	assignment.MerkleProofReceiverBefore = operatorA.Witnesses.MerkleProofReceiverBefore
	assignment.MerkleProofReceiverAfter = operatorA.Witnesses.MerkleProofReceiverAfter
	assignment.MerkleProofSenderBefore = operatorA.Witnesses.MerkleProofSenderBefore
	assignment.MerkleProofSenderAfter = operatorA.Witnesses.MerkleProofSenderAfter
	assignment.LeafReceiver = operatorA.Witnesses.LeafReceiver
	assignment.LeafSender = operatorA.Witnesses.LeafSender

	// PUBLIC INPUTS
	assignment.RootHashesBefore = operatorA.Witnesses.RootHashesBefore
	assignment.RootHashesAfter = operatorA.Witnesses.RootHashesAfter */

}

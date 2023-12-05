package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	. "rollup/src/gnark_rollup"
	. "rollup/src/tools"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tidwall/gjson"
)

const (
	fileToProof = "./test/cubicg16.proof"
	fileToPK    = "./test/cubicg16.pk"
	fileToCS    = "./test/cubicg16.cs"
)

func main() {
	// Initial setup
	var proof groth16.Proof
	var err error
	if FileExists(fileToProof) || FileExists(fileToPK) || FileExists(fileToCS) {
		fmt.Println("Proof files will be created")
		// Setup and generate files (solidity contract, pk, etc.)
		proof, err = SetupCubicProofAndContract(fileToProof, fileToPK, fileToCS) // Cubic example
		if err != nil {
			log.Fatalf("Error setting up proof and contract files: %v", err)
		}
	} else {
		// Load proof
		proof, err = LoadProofFromFile(fileToProof)
		if err != nil {
			log.Fatalf("Error loading proof: %v", err)
		} else {
			fmt.Println("Proof loaded from", fileToProof)
			// Now you can use the 'proof' variable in the code.
		}
	}

	// ----------------------- Main part -------------------
	// var shift uint64 = 256
	var amount uint64 = 10
	var operatorA, privKeys = CreateOP(10)

	/* fmt.Printf("Index of 1st account: %d\n", operatorA.State[0:32])
	fmt.Printf("Nonce of 1st account: %d\n", operatorA.State[32:64])
	fmt.Printf("Balance of 1st account: %d\n", operatorA.State[64:96])
	fmt.Printf("Balance of 2nd account: %d\n", operatorA.State[shift-32:shift])
	fmt.Printf("Public keys of accounts: %x\n", operatorA.AccountMap)
	fmt.Printf("Private key of 1st account: %x\n\n", privKeys[0]) */

	// get info on both parties of the transaction
	sender, err := operatorA.ReadAcc(0)
	if err != nil {
		log.Fatalf("Error reading sender account: %v", err)
	}

	receiver, err := operatorA.ReadAcc(1)
	if err != nil {
		log.Fatalf("Error reading receiver account: %v", err)
	}

	/* fmt.Printf("sender account: %x\n", sender.PubKey)
	fmt.Printf("sender account A.X: %x\n", sender.PubKey.A.X)
	fmt.Printf("sender account A.Y: %x\n", sender.PubKey.A.Y)
	fmt.Printf("sender account A.X.Bytes(): %x\n", sender.PubKey.A.X.Bytes()) */

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
	proof, err = groth16.Prove(r1cs, pk, witness)
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

	// ------------------ Proof (Cubic) --------------------
	eth_proof, _ := ProofToEthereumProof(proof)

	// ---------------- Truffle/Ganache contract call --------
	// Connect to your Ethereum node (Ganache instance or other network)
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545") // Update with your Ganache URL
	if err != nil {
		log.Fatalf("Failed to connect to ethclient/ganache: %v", err)
	}
	defer client.Close()

	// Read contract address from migration artefact json
	verifier_json, err := os.ReadFile("build/contracts/Verifier.json")
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

	// Call the contract function
	input_value := [1]*big.Int{}
	input_value[0] = new(big.Int).SetUint64(35)
	fmt.Println("Input y=", input_value)

	//var result []interface{}
	result, err := verifier.VerifyProof(nil, eth_proof.A, eth_proof.B, eth_proof.C, input_value)
	if err != nil {
		log.Fatalf("Failed to retrieve value: %v", err)
	}

	fmt.Println("Function Result:", result)

	/*
		// generate assets
		// gnark-solidity-checker generate --dir tmpdir --solidity contract_g16.sol
		cmd := exec.Command("gnark-solidity-checker", "generate", "--dir", tmpDir, "--solidity", "gnark_verifier.sol")
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("out: %x\n", out)
	*/
}

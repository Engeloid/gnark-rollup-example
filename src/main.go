package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	. "rollup/src/gnark_rollup"
	. "rollup/src/tools"
	"time"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"

	//"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

	// Create an operator for two different rollups with 16 accounts
	op := NewOperator(NumOfAcc)
	op2 := NewOperator(NumOfAcc)
	// If one file does not exist, create constraint system, new verifier contract etc.
	if FileExists(FileToRVK) || FileExists(FileToRPK) || FileExists(FileToRCS) || FileExists(FileToZKVerifier) {
		// Otherwise load them from local repo
		_, _, _, err = op.SetupZKVerifierContract(FileToRVK, FileToRPK, FileToRCS, FileToZKVerifier)
		if err != nil {
			log.Fatalf("Could not create zk verifier contract file: %v", err)
		}
	}
	if FileExists(FileToR2VK) || FileExists(FileToR2PK) || FileExists(FileToR2CS) || FileExists(FileToZKVerifier2) {
		// Otherwise load them from local repo
		_, _, _, err = op2.SetupZKVerifierContract(FileToR2VK, FileToR2PK, FileToR2CS, FileToZKVerifier2)
		if err != nil {
			log.Fatalf("Could not create zk verifier contract file: %v", err)
		}
	}

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

	// Create contract instance of first rollup
	r1Contract, r1ContractAddress := InstantiateRollup("build/contracts/Rollup-1.json", client, true)
	// Create contract instance of second rollup
	r2Contract, r2ContractAddress := InstantiateRollup("build/contracts/Rollup.json", client, false)

	latestRootafter, _ := r2Contract.GetMerkleRoot(nil)
	fmt.Printf("Init root on second contract: %x\n", latestRootafter.Bytes())

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
	fmt.Printf("Init merkle root: %x\n", big.NewInt(0).SetBytes(merkleRootInit))

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
	go op.ListenToDepositsWithdrawals(r1Contract, authOwner)

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
	_, err = r1Contract.Deposit(auth)
	if err != nil {
		log.Fatalf("Failed to deposit to rollup contract: %v", err)
	}
	// Second deposit to first acount
	time.Sleep(4 * time.Second)
	_, err = r1Contract.Deposit(auth)
	if err != nil {
		log.Fatalf("Failed to deposit to rollup contract: %v", err)
	}

	// Second accout on rollup
	time.Sleep(4 * time.Second)
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
	_, err = r1Contract.Deposit(auth2)
	if err != nil {
		log.Fatalf("Failed to deposit to rollup contract: %v", err)
	}

	time.Sleep(2 * time.Second)
	// Fetch and print contract balance
	contractBalance, err := client.BalanceAt(context.Background(), r1ContractAddress, nil)
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
	proof, err = op.CreateZKProof(FileToRVK, FileToRPK, FileToRCS, FileToZKVerifier)
	if err != nil {
		log.Fatalf("Failed to generate zkProof: %v", err)
	}

	// Transform data to fit contract call
	eth_proof, _ := ProofToEthereumProof(proof)
	input := [PublicInputSize]*big.Int{}
	// Note that first all of the first public inputs must be provided then the second and so on
	for i := 0; i < BatchSizeCircuit; i++ {
		input[i] = new(big.Int).SetBytes(op.RootHashesBef[i])
		input[i+BatchSizeCircuit] = new(big.Int).SetBytes(op.RootHashesAf[i])
	}

	// Submit batch
	_, err = r1Contract.SubmitVerifyBatch(authOwner, eth_proof.Proof, input)
	if err != nil {
		log.Fatalf("Failed to submit batch: %v", err)
	}
	fmt.Printf(Colorbold+ColorGreen+"zkProof sucessful on-chain: %x\n"+ColorReset, op.RootHashesAf[len(op.RootHashesAf)-1])
	/* latestRootafter, _ := rollupContract.GetMerkleRoot(nil)
	fmt.Printf("Latest root after batch: %x\n", latestRootafter.Bytes()) */

	// Withdrawal from contract
	time.Sleep(4 * time.Second)
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
	_, err = r1Contract.Withdraw(auth2, withdrawAmount, withdrawEthProof.Proof, withdrawAccountEth)
	if err != nil {
		log.Fatalf("Failed to withdraw: %v", err)
	}

	time.Sleep(2 * time.Second)
	// Fetch and print contract balance
	contractBalance, err = client.BalanceAt(context.Background(), r1ContractAddress, nil)
	if err != nil {
		log.Fatalf("Failed to fetch contract balance: %v", err)
	}
	fmt.Println("Contract balance after withdrawal:", contractBalance)
	contract2Balance, err := client.BalanceAt(context.Background(), r2ContractAddress, nil)
	if err != nil {
		log.Fatalf("Failed to fetch contract balance: %v", err)
	}
	fmt.Println("Contract balance of second rollup:", contract2Balance)

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
	latestRoot, _ := r1Contract.GetMerkleRoot(nil)
	if latestLocalBigRoot.Cmp(latestRoot) == 0 {
		fmt.Printf(Colorbold+ColorGreen+"Successfully updated accounts after withdrawal: %x\n"+ColorReset, latestLocalBigRoot)
	} else {
		fmt.Printf("Roots not identical! \nlocal root: %x, \ncontract root: %x\n", latestLocalBigRoot, latestRoot)
	}

	os.Exit(3)

	// ---- zk-proof verification on chain ------------
	// ownerPub := common.HexToAddress(accounts[5].PublicKey)

}

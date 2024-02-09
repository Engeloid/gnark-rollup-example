package tools

import (
	"bytes"
	"fmt"
	"math/big"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
)

type EthMerkleProof struct {
	Proof []*big.Int
}

type EthereumProof struct {
	Proof [8]*big.Int
}

func MerkleProofToEth(p [][]byte) (*EthMerkleProof, error) {
	var proof EthMerkleProof

	// Iterate over each element in the byte slice
	for _, proofBytes := range p {
		// Convert each byte slice into a *big.Int and append it to the Proof slice
		bigIntProofElement := new(big.Int).SetBytes(proofBytes)
		proof.Proof = append(proof.Proof, bigIntProofElement)
	}

	return &proof, nil
}

func ProofToEthereumProof(p groth16.Proof) (*EthereumProof, error) {
	var proof EthereumProof

	var buf bytes.Buffer
	_, err := p.WriteRawTo(&buf)
	if err != nil {
		return nil, fmt.Errorf("write raw proof to: %w", err)
	}
	proofBytes := buf.Bytes()

	proof.Proof[0] = new(big.Int).SetBytes(proofBytes[fp.Bytes*0 : fp.Bytes*1])
	proof.Proof[1] = new(big.Int).SetBytes(proofBytes[fp.Bytes*1 : fp.Bytes*2])
	proof.Proof[2] = new(big.Int).SetBytes(proofBytes[fp.Bytes*2 : fp.Bytes*3])
	proof.Proof[3] = new(big.Int).SetBytes(proofBytes[fp.Bytes*3 : fp.Bytes*4])
	proof.Proof[4] = new(big.Int).SetBytes(proofBytes[fp.Bytes*4 : fp.Bytes*5])
	proof.Proof[5] = new(big.Int).SetBytes(proofBytes[fp.Bytes*5 : fp.Bytes*6])
	proof.Proof[6] = new(big.Int).SetBytes(proofBytes[fp.Bytes*6 : fp.Bytes*7])
	proof.Proof[7] = new(big.Int).SetBytes(proofBytes[fp.Bytes*7 : fp.Bytes*8])

	return &proof, nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err) // returns false if file exists
}

func SavePKToFile(_pk groth16.ProvingKey, fileName string) error {
	// Create or open the file for writing.
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffer to store byte data
	//var buf bytes.Buffer
	_, err = _pk.WriteRawTo(file)
	if err != nil {
		return err
	}

	return nil

	/* // gen verifying key raw file
	f, err := os.Create("cubicg16.vk")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = vk.WriteRawTo(f)
	if err != nil {
		log.Fatalln(err)
	} */
}

func LoadPKFromFile(fileName string) (groth16.ProvingKey, error) {
	pk := groth16.NewProvingKey(ecc.BN254)
	// Open the file for reading.
	file, err := os.Open(fileName)
	if err != nil {
		return pk, err
	}
	defer file.Close()

	// Read and deserialize the proving key from the file.
	_, err = pk.ReadFrom(file)
	if err != nil {
		println("Error when reading raw file")
		return pk, err
	}

	return pk, nil
}

func SaveVKToFile(vk groth16.VerifyingKey, fileName string) error {
	// Create or open the file for writing.
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the verifying key to the file directly.
	_, err = vk.WriteRawTo(file)
	if err != nil {
		return err
	}

	return nil
}

func LoadVKFromFile(fileName string) (groth16.VerifyingKey, error) {
	vk := groth16.NewVerifyingKey(ecc.BN254)
	// Open the file for reading.
	file, err := os.Open(fileName)
	if err != nil {
		return vk, err
	}
	defer file.Close()

	// Read and deserialize the verifying key from the file.
	_, err = vk.ReadFrom(file)
	if err != nil {
		println("Error when reading raw file")
		return vk, err
	}

	return vk, nil
}

func SaveCSToFile(_cs constraint.ConstraintSystem, fileName string) error {
	// Create or open the file for writing.
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Store byte data for the contraint system in file
	_, err = _cs.WriteTo(file)
	if err != nil {
		return err
	}

	return nil
}

func LoadCSFromFile(fileName string) (constraint.ConstraintSystem, error) {
	cs := groth16.NewCS(ecc.BN254)
	// Open the file for reading.
	file, err := os.Open(fileName)
	if err != nil {
		return cs, err
	}
	defer file.Close()

	// Read and deserialize the constraint system from the file.
	_, err = cs.ReadFrom(file)
	if err != nil {
		println("Error when reading raw file")
		return cs, err
	}

	return cs, nil
}

func SaveProofToFile(_proof groth16.Proof, fileName string) error {
	// Create or open the file for writing.
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffer to store byte data
	//var buf bytes.Buffer
	_, err = _proof.WriteRawTo(file)
	if err != nil {
		return err
	}

	return nil
}

// Assuming 'fileName' is the name of the file where the proof is saved.
func LoadProofFromFile(fileName string) (groth16.Proof, error) {
	proof := groth16.NewProof(ecc.BN254)
	// Open the file for reading.
	file, err := os.Open(fileName)
	if err != nil {
		return proof, err
	}
	defer file.Close()

	// Create a decoder to deserialize the proof from the file.

	_, err = proof.ReadFrom(file)
	if err != nil {
		println("Error when reading raw file")
		return proof, err
	}

	return proof, nil
}

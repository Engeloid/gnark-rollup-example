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

type EthereumProof struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

func ProofToEthereumProof(p groth16.Proof) (*EthereumProof, error) {

	var proof EthereumProof

	var buf bytes.Buffer
	_, err := p.WriteRawTo(&buf)
	if err != nil {
		return nil, fmt.Errorf("write raw proof to: %w", err)
	}
	proofBytes := buf.Bytes()

	proof.A[0] = new(big.Int).SetBytes(proofBytes[fp.Bytes*0 : fp.Bytes*1])
	proof.A[1] = new(big.Int).SetBytes(proofBytes[fp.Bytes*1 : fp.Bytes*2])
	proof.B[0][0] = new(big.Int).SetBytes(proofBytes[fp.Bytes*2 : fp.Bytes*3])
	proof.B[0][1] = new(big.Int).SetBytes(proofBytes[fp.Bytes*3 : fp.Bytes*4])
	proof.B[1][0] = new(big.Int).SetBytes(proofBytes[fp.Bytes*4 : fp.Bytes*5])
	proof.B[1][1] = new(big.Int).SetBytes(proofBytes[fp.Bytes*5 : fp.Bytes*6])
	proof.C[0] = new(big.Int).SetBytes(proofBytes[fp.Bytes*6 : fp.Bytes*7])
	proof.C[1] = new(big.Int).SetBytes(proofBytes[fp.Bytes*7 : fp.Bytes*8])

	return &proof, nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err) // return true if file does not exist
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

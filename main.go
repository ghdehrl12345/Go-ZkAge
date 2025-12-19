package main

import (
	"fmt"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// 증명 구조체 정의
type AgeCircuit struct {
	CurrentYear frontend.Variable `gnark:",public"`
	LimitAge    frontend.Variable `gnark:",public"`
	BirthYear   frontend.Variable
}

// 로직 설계
func (circuit *AgeCircuit) Define(api frontend.API) error {
	myAge := api.Sub(circuit.CurrentYear, circuit.BirthYear)
	api.AssertIsLessOrEqual(circuit.LimitAge, myAge)

	return nil
}

// main 함수
func main() {
	var circuit AgeCircuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)

	pk, vk, _ := groth16.Setup(ccs)

	pkFile, _ := os.Create("zage.pk")
	pk.WriteTo(pkFile)
	pkFile.Close()

	vkFile, _ := os.Create("zage.vk")
	vk.WriteTo(vkFile)
	vkFile.Close()

	fmt.Println("zage.pk, zage.vk 생성 완료")

	// Witness 생성
	assignment := AgeCircuit{
		CurrentYear: 2025,
		LimitAge:    20,
		BirthYear:   2005,
	}

	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())

	// 증명서 생성
	proof, _ := groth16.Prove(ccs, pk, witness)

	proofFile, _ := os.Create("zage.proof")
	proof.WriteTo(proofFile)
	proofFile.Close()
}

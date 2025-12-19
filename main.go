package main

import (
	"github.com/consensys/gnark/frontend"
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

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

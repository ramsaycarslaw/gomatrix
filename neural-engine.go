package gomatrix

import (
	"log"
	"math"
)

//Sigmoid is used in machine learning
func Sigmoid(r, c int, z float64) float64 {
	return 1.0 / (1 + math.Exp(-1*z))
}

//SigmoidPrime is used in neural operations
func SigmoidPrime(m Matrix) (Matrix, error) {
	rows, _ := m.Dims()
	o := make([]float64, rows)
	for i := range o {
		o[i] = 1
	}
	ones, err := FilledMatrix(rows, 1, o)
	if err != nil {
		log.Fatal(err)
	}
	return multiply(m, subtract(ones, m)), err // m * (1 - m)
}

func multiply(m, n Matrix) Matrix {
	r, c := m.Dims()
	o, _ := FilledMatrix(r, c, nil)
	o, _ = m.MulElm(n)
	return o
}

func subtract(m, n Matrix) Matrix {
	r, c := m.Dims()
	o := GenerateMatrix(r, c)
	o = m.Sub(n)
	return o
}

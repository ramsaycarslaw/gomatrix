package gomatrix

import (
	"fmt"
	"math/rand"
	"time"
)

//IndexBuf gets iteration size of matrix
func (matrix Matrix) IndexBuf(row, col int) int {
	return row*int(matrix[1]) + col + 2
}

func raiseError(s string) (err error) {
	err = fmt.Errorf(s)
	return
}

func randInt(min, max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return min + r1.Intn(max-min)
}

//PrintMat gets nested form and prints each slice on a newline
func PrintMat(mat Matrix) {
	matrix := NestedForm(mat)
	for _, v := range matrix {
		fmt.Println(v)
	}
}

//NestedForm returns type matrix as a nested Row type
func NestedForm(mat Matrix) []Row {
	//rows := int(mat[0])
	colums := int(mat[1])
	matrix := []Row{}
	var row Row
	for _, v := range mat[2:] {
		row = append(row, v)
		if len(row) == colums {
			matrix = append(matrix, row)
			row = []float64{}
		}
	}
	return matrix
}

//Dims returns the dimensions of a given type Matrix
func (matrix Matrix) Dims() (rows, cols int) {
	rows = int(matrix[0])
	cols = int(matrix[1])
	return
}

//At returns the value from the ith column and the jth row
func (matrix Matrix) At(i, j int) (val float64) {
	val = matrix[i*int(matrix[1])+j+2]
	return
}

//Set changes the value at the ith row and jth column
func (matrix Matrix) Set(i, j int, val float64) {
	matrix[matrix.IndexBuf(i, j)] = val
}

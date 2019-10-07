package gomatrix

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func raiseError(s string) (err error) {
	err = fmt.Errorf(s)
	return
}

func randInt(min, max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return min + r1.Intn(max-min)
}

//PrintMatLegacy gets nested form and prints each slice on a newline
func PrintMatLegacy(mat Matrix) {
	matrix := NestedForm(mat)
	for _, v := range matrix {
		fmt.Println(v)
	}
}

//PrintVec formats a vector nicley
func PrintVec(vector Vector) {
	for _, v := range vector {
		fmt.Println("[" + strconv.Itoa(int(v)) + "]")
	}
}

//NestedForm returns type matrix as a nested Row type
func NestedForm(mat Matrix) [][]float64 {
	//rows := int(mat[0])
	colums := mat.Cols
	matrix := [][]float64{}
	var row []float64
	for _, v := range mat.Data {
		row = append(row, v)
		if uint(len(row)) == colums {
			matrix = append(matrix, row)
			row = []float64{}
		}
	}
	return matrix
}

//Dims returns the dimensions of a given type Matrix
func (matrix Matrix) Dims() (rows, cols uint) {
	rows = matrix.Rows
	cols = matrix.Cols
	return
}

//At returns the value from the ith row and the jth column
func (matrix Matrix) At(i, j int) (val float64) {
	val = matrix.Data[i*int(matrix.Cols)+j]
	return
}

//Set changes the value at the ith row and jth column
func (matrix Matrix) Set(i, j int, val float64) {
	matrix.Data[i*int(matrix.Cols)+j] = val
	return
}

//PrintMat prints a matrix with an optional heading
func (matrix *Matrix) PrintMat(heading string) {
	if heading > "" {
		fmt.Print("\n", heading, "\n")
	}
	for e := 0; e < len(matrix.Data); e += int(matrix.Cols) {
		fmt.Println(matrix.Data[e : e+int(matrix.Cols)])
	}
}

//RemoveCurrent removes a row and column from a matrix and returns a new one of correct dimensions.
func (matrix Matrix) RemoveCurrent(i, j int) (result Matrix) {
	result = GenerateMatrix(matrix.Rows-1, matrix.Cols-1)
	h := int(matrix.Rows)
	stride := int(matrix.Cols)
	//rr := int(result.Rows)
	arr := matrix.Data
	var m int
	j = j - 1
	i = i - 1
	mod := 0
	for k := 0; k < stride; k++ {
		m = (k*stride + j) - mod
		arr = arr[:m+copy(arr[m:], arr[m+1:])]
		mod++
	}
	arr = append(arr[:i], arr[i+(h-1):]...)
	result.Data = arr
	return
}

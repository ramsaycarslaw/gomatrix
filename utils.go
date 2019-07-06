package gomatrix

import "fmt"

//IndexBuf gets iteration size of matrix
func (matrix Matrix) IndexBuf(row, col int) int {
	return row*int(matrix[1]) + col + 2
}

func raiseError(s string) (err error) {
	err = fmt.Errorf(s)
	return
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
func Dims(mat Matrix) []int {
	var dimensions []int
	dimensions[0] = int(mat[0])
	dimensions[1] = int(mat[1])
	return dimensions
}

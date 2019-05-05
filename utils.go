package gomatrix

import "fmt"

func (matrix Matrix) IndexBuf(row, col int) int {
	return row*int(matrix[1]) + col + 2
}

func raiseError(s string) (err error) {
	err = fmt.Errorf(s)
	return
}

func PrintMat(mat Matrix) {
	//rows := int(mat[0])
	colums := int(mat[1])
	matrix := [][]float64{}
	row := []float64{}
	for _, v := range mat[2:] {
		row = append(row, v)
		if len(row) == colums {
			matrix = append(matrix, row)
			row = []float64{}
		}
	}
	for _, v := range matrix {
		fmt.Println(v)
	}
}

package gomatrix

import "fmt"

func (matrix Matrix) IndexFor(row, col int) int {
	return row*int(matrix[1]) + col + 2
}

func raiseError(s string) (err error) {
	err = fmt.Errorf(s)
	return
}

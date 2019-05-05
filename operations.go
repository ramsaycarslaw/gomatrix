package gomatrix

import "fmt"

//Dot multiplys two matrices
func Dot(mat1 Matrix, mat2 Matrix) (result Matrix, err error) {
	if mat1[0] != mat2[1] {
		err = raiseError(fmt.Sprintf("Matrices not aligned: cannot dot %v rows with %v colums", mat1, mat2))
		return
	}
	return
}

//DotProduct multiplys a row and a column
func DotProduct(row Row, colum Row) (product float64, err error) {
	if len(row) != len(colum) {
		err = raiseError(fmt.Sprintf("Type rows not aligned: %v and %v", len(row), len(colum)))
		return
	}
	products := []float64{}
	for i, v := range row {
		products = append(products, v*colum[i])
	}
	for _, v := range products {
		product = product + v
	}
	return
}

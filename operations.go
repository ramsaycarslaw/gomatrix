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

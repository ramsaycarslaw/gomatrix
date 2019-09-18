package gomatrix

import "fmt"

//MatDotVec multiplys a matrix and a vector
func (matrix Matrix) MatDotVec(vector Vector) (result Vector) {
	if int(matrix[0]) != len(vector) {
		raiseError(fmt.Sprintf("Matrix and Vector not aligned"))
	}
	result = GenerateVector(int(matrix[0]))
	for i := 0; i < int(matrix[0]); i++ {
		for j := 0; j < len(vector); j++ {
			result[i] = matrix[matrix.IndexBuf(i, j)] * vector[j]
		}
	}
	return result
}

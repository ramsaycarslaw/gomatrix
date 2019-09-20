package gomatrix

import "fmt"

//MatDotVec multiplys a matrix and a vector
func (matrix Matrix) MatDotVec(vector Vector) (result Vector) {
	if int(matrix[0]) != len(vector) {
		//check alignment
		raiseError(fmt.Sprintf("Matrix and Vector not aligned"))
	}
	//create blank reciever
	result = GenerateVector(int(matrix[0]))
	//for every row
	for i := 0; i < int(matrix[0]); i++ {
		//for every column
		for j := 0; j < len(vector); j++ {
			//multiply the two
			result[i] = matrix[matrix.IndexBuf(i, j)] * vector[j]
		}
	}
	return result
}

//AddVec adds two vectors and returns an answer of type vector
func (vector1 Vector) AddVec(vector2 Vector) (result Vector) {
	//check alignment
	if len(vector1) != len(vector2) {
		raiseError(fmt.Sprintf("Vectors not aligned"))
	}
	//create blank vector
	result = GenerateVector(len(vector1))
	for i, v := range vector1 {
		//perform addition and save in result
		result[i] = v + vector2[i]
	}
	return
}

//SubVec adds two vectors and returns an answer of type vector
func (vector1 Vector) SubVec(vector2 Vector) (result Vector) {
	//check alignment
	if len(vector1) != len(vector2) {
		raiseError(fmt.Sprintf("Vectors not aligned"))
	}
	//create blank vector
	result = GenerateVector(len(vector1))
	for i, v := range vector1 {
		//perform addition and save in result
		result[i] = v - vector2[i]
	}
	return
}

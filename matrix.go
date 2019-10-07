package gomatrix

import "fmt"

//Matrix is the default type in gomatrix
type Matrix struct {
	Rows uint // ?Maybe not needed
	Cols uint
	Data []float64
}

//GenerateMatrix creates a blank matrix
func GenerateMatrix(rows, cols uint) (matrix Matrix) {
	data := make([]float64, rows*cols)
	matrix.Data = data
	matrix.Rows = rows
	matrix.Cols = cols
	return
}

//FilledMatrix creates a new matrix on a single thread
func FilledMatrix(rows, cols uint, data []float64) (matrix Matrix, err error) {
	if rows == 0 || cols == 0 {
		err = raiseError("Cannot make empty matrix")
		return
	} else if rows*cols != uint(len(data)) {
		err = raiseError("Incorrect number of elements")
	}
	matrix = GenerateMatrix(rows, cols)
	matrix.Data = data

	return
}

//RandFill creates and fills a matrix with random integers
func RandFill(rows, cols uint, min, max int) (mat Matrix) {
	mat = GenerateMatrix(rows, cols)
	for i := range mat.Data {
		random := randInt(min, max)
		mat.Data[i] = float64(random)
	}
	return
}

//Delete removes all items from a matrix
func (matrix Matrix) Delete() {
	matrix.Data = nil
}

// Augment augments a matrix which spliuts it into a column and a matrix
func (matrix *Matrix) Augment(B *Matrix) (C *Matrix, err error) {
	if matrix.Rows != B.Rows {
		raiseError(fmt.Sprintf("Dimensions Mismatched"))
		return
	}
	C = Zeros(matrix.Rows, matrix.Cols+B.Cols)
	err = matrix.AugmentFill(B, C)
	return
}

//AugmentFill is used to fill a matrix
func (matrix *Matrix) AugmentFill(B, C *Matrix) (err error) {
	if matrix.Rows != B.Rows || C.Rows != matrix.Rows || C.Cols != matrix.Cols+B.Cols {
		raiseError(fmt.Sprintf("Dimensions Mismatched"))
		return
	}
	C.SetMatrix(0, 0, matrix)
	C.SetMatrix(0, int(matrix.Cols), B)
	return
}

//SetMatrix puts B 0,0 in A i,j
func (matrix *Matrix) SetMatrix(i, j int, B *Matrix) {
	for r := 0; r < int(B.Rows); r++ {
		for c := 0; c < int(B.Cols); c++ {
			matrix.Set(i+r, j+c, B.At(r, c))
		}
	}
}

//Identity creates an identity matrix
func Identity(span int) *Matrix {
	A := Zeros(uint(span), uint(span))
	for i := 0; i < span; i++ {
		A.Set(i, i, 1)
	}
	return A
}

//SwapRows swaps 2 rows
func (matrix *Matrix) SwapRows(r1 int, r2 int) {
	index1 := r1 * int(matrix.Cols)
	index2 := r2 * int(matrix.Cols)
	for j := 0; j < int(matrix.Cols); j++ {
		matrix.Data[index1], matrix.Data[index2] = matrix.Data[index2], matrix.Data[index1]
		index1++
		index2++
	}
}

//ScaleRow scales a row by a float f
func (matrix *Matrix) ScaleRow(r int, f float64) {
	i := r * int(matrix.Cols)
	for j := 0; j < int(matrix.Cols); j++ {
		matrix.Data[i] *= f
		i++
	}
}

//ScaleAddRow adds and scales a row by a float
func (matrix *Matrix) ScaleAddRow(ri int, rj int, f float64) {
	indexd := ri * int(matrix.Cols)
	indexs := rj * int(matrix.Cols)
	for j := 0; j < int(matrix.Cols); j++ {
		matrix.Data[indexd] += f * matrix.Data[indexs]
		indexd++
		indexs++
	}
}

//GetMatrix is like GenerateMatrix but returns a pointer
func (matrix *Matrix) GetMatrix(i, j int, rows, cols uint) *Matrix {
	B := new(Matrix)
	B.Data = matrix.Data[i*int(matrix.Cols)+j : i*int(matrix.Cols)+j+(int(rows)-1)*int(matrix.Cols)+int(cols)]
	B.Rows = rows
	B.Cols = cols
	return B
}

//Zeros fills a matrix with zeros and returns a ponter
func Zeros(rows, cols uint) *Matrix {
	A := new(Matrix)
	A.Data = make([]float64, int(rows*cols))
	A.Rows = rows
	A.Cols = cols
	return A
}

package gomatrix

import "runtime"

//Matrix is the default type in gomatrix
type Matrix []float64

//Row is used for construction
type Row []float64

//Columns is a list of rows
type Columns []Row

//GenerateMatrix creates a blank matrix
func GenerateMatrix(rows, cols int) (mat Matrix) {
	mat = make(Matrix, rows*cols+2)
	mat[0] = float64(rows)
	mat[1] = float64(cols)

	return
}

//FilledMatrix chooses between concurrent and single threaded
func FilledMatrix(rows, cols int, data []float64) (matrix Matrix, err error) {
	var large bool
	if rows > 100 {
		large = true
	} else {
		large = false
	}
	switch large {
	case true:
		matrix, err = filledMatrixPar(rows, cols, data)
		return
	default:
		matrix, err = filledMatrixSingle(rows, cols, data)
		return
	}
}

//FilledMatrixSingle creates a new matrix on a single thread
func filledMatrixSingle(rows, cols int, data []float64) (matrix Matrix, err error) {
	if rows == 0 || cols == 0 {
		err = raiseError("Cannot make empty matrix")
		return
	} else if rows*cols != len(data) {
		err = raiseError("Incorrect number of elements")
	}
	matrix = GenerateMatrix(rows, cols)
	for i, v := range data {
		matrix[i+2] = v
	}
	return
}

func filledMatrixPar(rows, cols int, data []float64) (matrix Matrix, err error) {
	if rows == 0 || cols == 0 {
		err = raiseError("Cannot make empty matrix")
		return
	} else if rows*cols != len(data) {
		err = raiseError("Incorrect number of elements")
		return
	}
	matrix = GenerateMatrix(rows, cols)
	in := make(chan int)
	exit := make(chan bool)
	fill := func() {
		for {
			select {
			case i := <-in:
				matrix[i+2] = data[i]
			case <-exit:
				return
			}
		}

	}
	threads := runtime.GOMAXPROCS(0) + 2

	for i := 0; i < threads; i++ {
		go fill()
	}

	for i := range data {
		in <- i
	}

	for i := 0; i < threads; i++ {
		exit <- true
	}
	return
}

//FillMatrix adds values to an existing matrix
func (matrix Matrix) FillMatrix(data []float64) {
	for i := 2; i <= len(data); i++ {
		matrix[i] = data[i-2]
	}
}

//Zeros creates a matrix and fills it with zeros
func Zeros(rows, cols int) (mat Matrix) {
	mat = make(Matrix, rows*cols+2)
	var zero float64
	for i := range mat {
		mat[i] = zero
	}
	mat[0] = float64(rows)
	mat[1] = float64(cols)

	return
}

//RandFill creates and fills a matrix with random integers
func RandFill(rows, cols, min, max int) (mat Matrix) {
	mat = make(Matrix, rows*cols+2)
	for i := range mat {
		random := randInt(min, max)
		mat[i] = float64(random)
	}
	mat[0] = float64(rows)
	mat[1] = float64(cols)

	return
}

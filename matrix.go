package gomatrix

type Matrix []float64
type Row []float64
type Columns []Row

func GenerateMatrix(rows, cols int) (mat Matrix) {
	mat = make(Matrix, rows*cols+2)
	mat[0] = float64(rows)
	mat[1] = float64(cols)

	return
}

func FillMatrix(rows, cols int, data []float64) (matrix Matrix, err error) {
	if rows == 0 || cols == 0 {
		err = raiseError("Cannot make empty matrix")
		return
	} else if rows*cols != len(data) {
		err = raiseError("Incorrect number of elements")
	}
	matrix = GenerateMatrix(rows, cols)
	for i := 2; i <= len(data); i++ {
		matrix[i] = data[i-2]
	}
	return
}

//Zeros creates a matrix and fills it with zeros
func Zeros(rows, cols int) (mat Matrix) {
	mat = make(Matrix, rows*cols+2)
	var zero float64 = 0
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

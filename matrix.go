package gomatrix

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

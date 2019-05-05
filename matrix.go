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

func FillMatrix(build Columns) (matrix Matrix, err error) {
	if len(build) == 0 || len(build[0]) == 0 {
		err = raiseError("Cannot make empty matrix")
		return
	}
	matrix = GenerateMatrix(len(build), len(build[0]))
	for i, row := range build {
		for j, value := range row {
			matrix[matrix.IndexBuf(i, j)] = value
		}
	}
	return
}

package gomatrix

import "fmt"

//Vector is a 1D matrix
type Vector []float64

//GenerateVector creates a blank vector
func GenerateVector(rows int) (vector Vector) {
	vector = make(Vector, rows)
	return
}

//FilledVector creates and fills a vector with custom data
func FilledVector(rows int, data []float64) (vector Vector) {
	if rows == 1 {
		raiseError(fmt.Sprintf("Cannot make 1x1 vector"))
	}
	if len(data) != rows {
		raiseError(fmt.Sprintf("Incorrect number of elements"))
	}
	vector = GenerateVector(rows)
	for i, v := range data {
		vector[i] = v
	}
	return
}

//RandFillVector creates and fills a matrix with random integers
func RandFillVector(rows, min, max int) (vector Vector) {
	vector = make(Vector, rows)
	for i := range vector {
		random := randInt(min, max)
		vector[i] = float64(random)
	}
	return
}

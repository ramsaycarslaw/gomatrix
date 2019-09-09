package gomatrix

import (
	"fmt"
	"runtime"
)

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

//Add adds two matrices of the same dimensions
func Add(mat1 Matrix, mat2 Matrix) (result Matrix, err error) {
	if mat1[0] != mat2[0] && mat1[1] != mat2[1] {
		err = raiseError(fmt.Sprintf("Matrices not aligned: cannot perform addition"))
		return
	}
	matrix1 := mat1[2:]
	matrix2 := mat2[2:]
	for i := range matrix1 {
		result[i] = matrix1[i] + matrix2[i]
	}
	return
}

//Add adds two matrices of the same dimensions
func (mat1 Matrix) Sub(mat2 Matrix) (Matrix, err error) {
	if mat1[0] != mat2[0] && mat1[1] != mat2[1] {
		err = raiseError(fmt.Sprintf("Matrices not aligned: cannot perform addition"))
		return
	}
	result := GenerateMatrix(int(mat1[0]), int(mat2[1]))
	matrix1 := mat1[2:]
	matrix2 := mat2[2:]
	for i := range matrix1 {
		result[i+2] = matrix1[i] + matrix2[i]
	}
	return
}

//Dot multiplys concurrently
//Dot multiplys concurrently
func (mat1 Matrix) Dot(mat2 Matrix) (result Matrix, err error) {
	if mat1[0] != mat2[1] {
		//check matrix alignment
		err = raiseError(fmt.Sprintf("Matrices not aligned: cannot multiply"))
		return
	}
	//fill result with zeros
	//Dimensions are known as matrix multiplication is predictable
	result = GenerateMatrix(int(mat1[0]), int(mat2[1]))

	//create channels for concurrency
	in := make(chan int)
	exit := make(chan bool)

	//create function literal (lambda function) for a concurrent model
	dot := func() {
		for {
			//select lets a goroutine wait for a communication
			//it is used to dertermine if there is an input or exit control
			select {
			//in the case of input it performs the multiplication of
			//one row and one column
			case i := <-in:
				sums := make([]float64, int(mat2[1]))
				//create blank list to store values for new matrix
				for k := 0; k < int(mat1[1]); k++ {
					for j := 0; j < int(mat2[1]); j++ {
						//multiply row 1 by col 1, add to row 2 by col 2 etc
						sums[j] += mat1.At(i, k) * mat2.At(k, j)
					}
				}
				//once finished, update result
				for j := 0; j < int(mat2[1]); j++ {
					result.Set(i, j, sums[j])
				}
			// if the maths is finished, return
			case <-exit:
				return
			}
		}
	}
	//get the max thread count
	threads := runtime.GOMAXPROCS(0) + 2

	for i := 0; i < threads; i++ {
		//perform on every avliable thread
		go dot()
	}
	for i := 0; i < int(mat1[0]); i++ {
		//pass dot() the rows
		in <- i
	}

	for i := 0; i < threads; i++ {
		//exit the goroutine every time
		exit <- true
	}

	//main return in the case of error
	return

}

//Reshape changes the dimensions of a matrix
func (matrix Matrix) Reshape(i, j int) (err error) {
	if int(matrix[0])*int(matrix[1]) != i*j {
		err = raiseError(fmt.Sprintf("Matrices not aligned"))
		return
	}
	matrix[0] = float64(i)
	matrix[1] = float64(j)
	return
}

func (matrix1 Matrix) DotNaive(matrix2 Matrix) (result Matrix, err error) {
	if matrix1[0] != matrix2[1] {
		raiseError(fmt.Sprintf("Matrices not aligned: cannot multiply"))
		return
	}
	ar := int(matrix1[0])
	ac := int(matrix1[1])
	bc := int(matrix2[1])
	result = GenerateMatrix(ar, bc)
	for i := 0; i < ar; i++ {
		for j := 0; j < bc; j++ {
			var sum float64
			for k := 0; k < ac; k++ {
				sum = sum + matrix1.At(i, k)*matrix2.At(k, j)
			}
			result.Set(i, j, sum)
		}
	}
	return
}

func (matrix1 Matrix) Dot2(matrix2 Matrix) (result Matrix, err error) {
	ar := int(matrix1[0])
	ac := int(matrix1[1])
	bc := int(matrix2[1])
	in := make(chan int)
	exit := make(chan bool)
	dot := func() {
		for {
			select {
			case r := <-in:
				row := make([]float64, ac)
				for i := range row {
					row[i] = matrix1.At(r, i)
				}
				for c := 0; c < bc; c++ {
					var v float64
					for i, e := range row {
						v += e * matrix2.At(i, c)
					}
					result[r*int(result[1])+c+2] = v
				}
			case <-exit:
				return

			}
		}

	}

	threads := runtime.GOMAXPROCS(0) + 2

	for r := 0; r < threads; r++ {
		//perform on every avliable thread
		go dot()
	}
	for r := 0; r < ar; r++ {
		//pass dot() the rows
		in <- r
	}

	for r := 0; r < threads; r++ {
		//exit the goroutine every time
		exit <- true
	}
	return
}

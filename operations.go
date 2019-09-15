package gomatrix

import (
	"fmt"
	"runtime"
)

//Dot multiplys concurrently
func (matrix1 Matrix) Dot(matrix2 Matrix) (result Matrix, err error) {
	if matrix1[0] != matrix2[1] {
		//check matrix alignment
		err = raiseError(fmt.Sprintf("Matrices not aligned: cannot multiply"))
		return
	}
	//fill result with zeros
	//Dimensions are known as matrix multiplication is predictable
	result = GenerateMatrix(int(matrix1[0]), int(matrix2[1]))

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
				sums := make([]float64, int(matrix2[1]))
				//create blank list to store values for new matrix
				for k := 0; k < int(matrix1[1]); k++ {
					for j := 0; j < int(matrix2[1]); j++ {
						//multiply row 1 by col 1, add to row 2 by col 2 etc
						sums[j] += matrix1.At(i, k) * matrix2.At(k, j)
					}
				}
				//once finished, update result
				for j := 0; j < int(matrix2[1]); j++ {
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
	for i := 0; i < int(matrix1[0]); i++ {
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
func (matrix1 Matrix) Reshape(i, j int) (err error) {
	if int(matrix1[0])*int(matrix1[1]) != i*j {
		err = raiseError(fmt.Sprintf("Matrices not aligned"))
		return
	}
	matrix1[0] = float64(i)
	matrix1[1] = float64(j)
	return
}

//DotNaive uses only one core
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

//Times performs an element-wise multiplication on two matrices
func (matrix1 Matrix) Times(matrix2 Matrix) (result Matrix, err error) {
	if matrix1[0] != matrix2[0] || matrix1[1] != matrix2[1] {
		raiseError(fmt.Sprintf("Matrices must have the same dimensions"))
	}
	result = GenerateMatrix(int(matrix1[0]), int(matrix1[1]))
	for i, v := range matrix1 {
		result[i] = v * matrix2[i]
	}
	result[0], result[1] = matrix1[0], matrix1[1]
	return
}

package gomatrix

import (
	"fmt"
	"math"
	"runtime"
)

//Dot multiplys concurrently
func (matrix1 Matrix) Dot(matrix2 Matrix) (result Matrix, err error) {
	if matrix1.Cols != matrix2.Rows {
		//check matrix alignment
		err = raiseError(fmt.Sprintf("Matrices not aligned: cannot multiply"))
		return
	}
	//fill result with zeros
	//Dimensions are known as matrix multiplication is predictable
	result = GenerateMatrix(matrix1.Rows, matrix2.Cols)
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
				sums := make([]float64, int(matrix2.Cols))
				//create blank list to store values for new matrix
				for k := 0; k < int(matrix1.Cols); k++ {
					for j := 0; j < int(matrix2.Cols); j++ {
						//multiply row 1 by col 1, add to row 2 by col 2 etc
						sums[j] += matrix1.At(i, k) * matrix2.At(k, j)
					}
				}
				//once finished, update result
				for j := 0; j < int(matrix2.Cols); j++ {
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
	for i := 0; i < int(matrix1.Rows); i++ {
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
func (matrix1 Matrix) Reshape(i, j uint) (err error) {
	if matrix1.Rows*matrix1.Cols != i*j {
		err = raiseError(fmt.Sprintf("Matrices not aligned"))
		return
	}
	matrix1.Rows = i
	matrix1.Cols = j
	return
}

//DotNaive uses only one core
func (matrix1 Matrix) DotNaive(matrix2 Matrix) (result Matrix, err error) {
	if matrix1.Rows != matrix2.Cols {
		raiseError(fmt.Sprintf("Matrices not aligned: cannot multiply"))
		return
	}
	ar := int(matrix1.Rows)
	ac := int(matrix1.Cols)
	bc := int(matrix2.Cols)
	//get dimensions
	result = GenerateMatrix(uint(ar), uint(bc))
	//for i in rows
	for i := 0; i < ar; i++ {
		//for j in columns
		for j := 0; j < bc; j++ {
			var sum float64
			//multiply individual items
			for k := 0; k < ac; k++ {
				sum = sum + matrix1.At(i, k)*matrix2.At(k, j)
			}
			//set the value of result
			result.Set(i, j, sum)
		}
	}
	return
}

//MulElm performs an element-wise multiplication on two matrices
func (matrix1 Matrix) MulElm(matrix2 Matrix) (result Matrix, err error) {
	if matrix1.Rows != matrix2.Rows || matrix1.Cols != matrix2.Cols {
		raiseError(fmt.Sprintf("Matrices must have the same dimensions"))
	}
	result = GenerateMatrix(matrix1.Rows, matrix1.Cols)
	//create blank matrix
	for i, v := range matrix1.Data {
		result.Data[i] = v * matrix2.Data[i]
	}
	//set dimensions
	return
}

//Sub subtracts two matrices
func (matrix1 Matrix) Sub(matrix2 Matrix) (result Matrix) {
	ar, ac := matrix1.Dims()
	br, bc := matrix2.Dims()
	//get columns and rows
	if ar != br || ac != bc {
		raiseError(fmt.Sprintf("Cannot subratact, not aligned"))
	}
	//new blank matrix
	result = GenerateMatrix(ar, ac)
	for i, v := range matrix1.Data {
		result.Data[i] = v - matrix2.Data[i]
	}
	return
}

//Add adds two matrices
func (matrix1 Matrix) Add(matrix2 Matrix) (result Matrix) {
	ar, ac := matrix1.Dims()
	br, bc := matrix2.Dims()
	//get columns and rows
	if ar != br || ac != bc {
		raiseError(fmt.Sprintf("Cannot subratact, not aligned"))
	}
	//new blank matrix
	result = GenerateMatrix(ar, ac)
	for i, v := range matrix1.Data {
		result.Data[i] = v + matrix2.Data[i]
	}
	return
}

//T gets the transposed form of the matrix in the reciever
func (matrix *Matrix) T() {
	h := int(matrix.Rows)
	//get the index of the data
	for start := range matrix.Data {
		next := start
		i := 0
		//while next <= start, compute new position
		for {
			i++
			//find next value to update
			next = (next%h)*int(matrix.Cols) + next/h
			//base case
			if next <= start {
				break
			}
		}
		//make sure it doesnt exit early
		if next < start || i == 1 {
			continue
		}
		//reset next
		next = start
		//create a tempory variable
		tmp := matrix.Data[next]
		for {
			//find new i value
			i = (next%h)*int(matrix.Cols) + next/h
			if i == start {
				//set next to temp
				matrix.Data[next] = tmp
			} else {
				//set next to i
				matrix.Data[next] = matrix.Data[i]
			}
			next = i
			//base case
			if next <= start {
				break
			}
		}
	}
	//reset
	matrix.Rows = matrix.Cols
	matrix.Cols = uint(h)
}

//Cofactors applys a plus minus grid to a matrix
func (matrix1 *Matrix) Cofactors() {
	//set a flag variable
	plus := true
	//for index, value in elements of reciever matrix
	for i, v := range matrix1.Data {
		if plus == true {
			//just set false as adding + to number has no effect
			plus = false
			//continue stops overwiting
			continue
		} else if plus == false {
			//v - 2*v converts the string
			matrix1.Data[i] = v * -1
			plus = true
			//continue stops overwriting
			continue
		}
	}
}

//Determinant calculates the dterminant of any given square matrix
func (matrix1 Matrix) Determinant() (D float64) {
	if matrix1.Rows != matrix1.Cols {
		raiseError(fmt.Sprintf("Cannot find the determinant of non-square matrix"))
	} else {
		if matrix1.Rows == 2 {
			D = ((matrix1.At(1, 1) * matrix1.At(2, 2)) - (matrix1.At(1, 2) * matrix1.At(2, 1)))
		} else {
			var o float64
			var Dtemp Matrix
			for i := 0; i < int(matrix1.Rows); i++ {
				Dtemp = matrix1.RemoveCurrent(1, i)
				o = matrix1.At(1, i) * (math.Pow(-1, (float64(i) + 1))) * Dtemp.Determinant()
				if i == 1 {
					D = o
				} else {
					D = D + o
				}
			}
		}
	}
	return
}

//Inverse computes the inverse of a matrix
func (A Matrix) Inverse() (*Matrix, error) {
	if A.Rows != A.Cols {
		raiseError(fmt.Sprintf("Dimensions not aligned: mathematical error"))
	}
	aug, _ := A.Augment(Identity(int(A.Rows)))
	for i := 0; i < int(aug.Rows); i++ {
		j := i
		for k := i; k < int(aug.Rows); k++ {
			if math.Abs(aug.At(k, i)) > math.Abs(aug.At(j, i)) {
				j = k
			}
		}
		if j != i {
			aug.SwapRows(i, j)
		}
		if aug.At(i, i) == 0 {
			return nil, nil
		}
		aug.ScaleRow(i, 1.0/aug.At(i, i))
		for k := 0; k < int(aug.Rows); k++ {
			if k == i {
				continue
			}
			aug.ScaleAddRow(k, i, -aug.At(k, i))
		}
	}
	inv := aug.GetMatrix(0, int(A.Cols), A.Rows, A.Cols)
	return inv, nil
}

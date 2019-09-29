# gomatrix
Gomatrix is a library for the Go programming language. It features the key features for matrix arithmatic and linear algebra with a foucus on simplicity.

**Installation**

To install gomatrix on linux/mac paste the following command into your terminal:

`go get github.com/ramsaycarslaw/gomatrix`

**Dependencies**

Gomatrix has currently does not have dependencies.

**Examples**

To create two random matrices, multiply then print:

```
x := gomatrix.RandFill(10,10, 1, 100)
y := gomatrix.RandFill(10,10, 1, 100)

z, err := x.Dot(y)

if err != nil {
	log.Fatal(err)
}

gomatrix.PrintMat(z)
```

This will print to the terminal.

To fill a matrix with custom data from a slice of float64s (and print):
```
data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}

x := gomatrix.FilledMatrix(3, 3, data)

gomatrix.PrintMat(x)
```

Gomatrix can also be used to transform a point by a matrix:

```
//create vector
data := []float64{2, 1}
point := gomatrix.FilledVector(2, data)

//create matrix
matrixData := []float64{0, 0, 1, 0}
matrix := gomatrix.FilledMatrix(2, 2, matrixData)

//compute transformation
transform := matrix.MatDotVec(vector)

//show new point
fmt.Println(transform)
```

  

**Authors**

Ramsay Carslaw, 2019

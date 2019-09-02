# gomatrix
Gomatrix is a library for the Go programming language. It features the key features for matrix arithmatic and linear algebra with a foucus on simplicity.

**Installation**

To install gomatrix on linux/mac paste the following command into your terminal:

`go get github.com/ramsaycarslaw/gomatrix`

Gomatrix has currently does not have dependencies.

**Examples**

To create two random matrices, multiply then print:

```x := RandFill(10,10, 1, 100)
y := RandFill(10,10, 1, 100)

z, err := x.Dot(y)

if err != nil {
	log.Fatal(err)
}

PrintMat(z)
```

This will print to the terminal.
  

**Authors**

Ramsay Carslaw, 2019

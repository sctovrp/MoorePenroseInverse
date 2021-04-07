package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type Matrix [][]float64

func main() {
	// a := Matrix{
	// 	{0, 1, 2, 3}, /*  initializers for row indexed by 0 */
	// 	{4, 5, 6, 7}, /*  initializers for row indexed by 1 */
	// 	{8, 9, 10, 11},
	// }
	// Generate a 6×6 matrix of random values.
	data := make([]float64, 12)
	var s = -1
	for i := range data {
		s = s + 1
		data[i] = float64(s)
	}
	a := mat.NewDense(3, 4, data)
	matPrint(a)

	b := a.T()
	matPrint(b)

	mul1 := mat.NewDense(3, 3, nil)
	mul1.Mul(a, b)
	fmt.Println("multiplicacion:")
	matPrint(mul1)

	mul2 := mat.NewDense(3, 3, nil)
	mul2.Inverse(mul1)
	matPrint(mul1)

	fmt.Println(mul2.Dims())
	fmt.Println(b.Dims())
	final := mat.NewDense(4, 4, nil)
	final.Mul(mul2, b)
	matPrint(final)

	// var i, j int
	// for i = 0; i < 3; i++ {
	// 	for j = 0; j < 4; j++ {
	// 		fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
	// 	}
	// }
	// b := a
	// fmt.Println("b", b)
	// fmt.Println("a:")
	// for _, h := range a {
	// 	for _, n := range h {
	// 		fmt.Print(n, "\t")
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println(transpose(a))
}
func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

// func transpose(x [][]float64) Matrix {
// 	n := len(x[0])
// 	m := len(x)
// 	fmt.Println(n)
// 	fmt.Println(m)
// 	out := make([][]float64, m)
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			out[j] = append(out[j], x[i][j])
// 		}
// 	}
// 	return out
// }

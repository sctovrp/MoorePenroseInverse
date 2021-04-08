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
	// data := make([]float64, 12)
	// var s = -1
	// for i := range data {
	// 	s = s + 1
	// 	data[i] = float64(s)
	// }
	fmt.Println("Bienvenido al programa de Matrix inversa por Moore-Penrose")

	fmt.Println("Ingrese la dimension de su matrix en n (filas)")
	var n int
	fmt.Scanln(&n)
	fmt.Println("n:", n)

	fmt.Println("Ingrese la dimension de su matrix en m (Columnas)")
	var m int
	fmt.Scanln(&m)
	fmt.Println("n:", m)

	fmt.Println("Ahora puede ingresar los datos de la matrix")
	data := m * n
	datos_ingr := set(data)
	fmt.Println("datos ingresados:", datos_ingr)
	a := mat.NewDense(n, m, datos_ingr)
	matPrint(a)

	a_trans := a.T()
	matPrint(a_trans)
	// b_prim := a_prim.T()
	// matPrint(b_prim)

	mul1 := mat.NewDense(m, m, nil)
	mul1.Mul(a_trans, a)
	fmt.Println("multiplicacion:")
	matPrint(mul1)

	inver := mat.NewDense(m, m, nil)
	inver.Inverse(mul1)
	fmt.Println("inversa:")
	matPrint(inver)

	final_matrix := mat.NewDense(m, n, nil)
	final_matrix.Mul(inver, a_trans)
	fmt.Println("final matrix:")
	matPrint(final_matrix)

	checker := mat.NewDense(m, m, nil)
	checker.Mul(final_matrix, a)
	fmt.Println("checker:")
	matPrint(checker)
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func set(n int) []float64 {
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(a)
	return a
}

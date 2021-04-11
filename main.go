package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type Matrix [][]float64

func main() {

	//
	// Version de go: 1.16
	// Importar la libreta gonum, version v0.9.1: "gonum.org/v1/gonum/mat"
	// Web documentacion gonum: https://pkg.go.dev/gonum.org/v1/gonum/mat#section-documentation
	//
	// Datos provisionales de un ejemplo:
	// n = 3
	// m = 2
	// coeficientes de la matriz = 1 3 5 7 9 11 13
	//

	fmt.Println("Bienvenido al programa de Matrix inversa por Moore-Penrose")

	fmt.Println("Ingrese la dimension de su matrix en n (filas)")
	var n int
	fmt.Scanln(&n)
	fmt.Println("n:", n)

	fmt.Println("Ingrese la dimension de su matrix en m (Columnas)")
	var m int
	fmt.Scanln(&m)
	fmt.Println("n:", m)

	fmt.Println("Ahora puede ingresar los coeficientes de la matrix")
	data := m * n
	datos_ingr := set(data)
	fmt.Println("Coeficientes ingresados:", datos_ingr)
	a := mat.NewDense(n, m, datos_ingr)
	matPrint(a)

	a_trans := a.T()
	fmt.Println("Transpuesta:")
	matPrint(a_trans)

	mul1 := mat.NewDense(m, m, nil)
	mul1.Mul(a_trans, a)
	fmt.Println("Multiplicacion:")
	matPrint(mul1)

	inver := mat.NewDense(m, m, nil)
	inver.Inverse(mul1)
	fmt.Println("Inversa primera multiplicacion:")
	matPrint(inver)

	final_matrix := mat.NewDense(m, n, nil)
	final_matrix.Mul(inver, a_trans)
	fmt.Println("Matriz inversa por el metodo Moore-Penrose:")
	matPrint(final_matrix)

	checker := mat.NewDense(m, m, nil)
	checker.Mul(final_matrix, a)
	fmt.Println("Checker:")
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

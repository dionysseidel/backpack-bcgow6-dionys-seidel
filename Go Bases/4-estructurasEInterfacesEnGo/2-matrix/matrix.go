/* Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo */

package main

import "fmt"

type Matrix struct {
	Height      int
	Width       int
	IsCuadratic bool
	MaxValue    float64
	Values      []float64
}

func (m *Matrix) Set(values ...float64) {
	m.Values = values
}

func (m *Matrix) Print() {
	k := 0
	for i := 0; i < m.Height && k < len(m.Values); i++ {
		for j := 0; j < m.Width && k < len(m.Values); j++ {
			fmt.Printf("%f\t", m.Values[k])
			k++
		}
		fmt.Print("\n")
	}
}

func main() {
	m := Matrix{
		Height: 4,
		Width:  4,
	}
	m.Set(1, 2, 3, 4, 5, 6)
	m.Print()
}

/* Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le descontará además un 10% */

package main

import "fmt"

func calculateTax(salary float64) (tax float64) {

	if salary > 150000 {
		return salary * 0.27
	} else if salary > 50000 {
		return salary * 0.17
	}

	return

}

func main() {
	fmt.Println(calculateTax(296000))
	fmt.Println(calculateTax(100000))
	fmt.Println(calculateTax(40000))
}

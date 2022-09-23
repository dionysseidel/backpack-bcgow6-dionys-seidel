/* Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo
Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000. 
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes. */

package main

import "fmt"

var age int = 18
var isEmployed bool = true
var seniority int = 2
var salary float64 = 296000

func main()  {
	if age > 22 {
		switch {
		case isEmployed && seniority > 1:
			fmt.Println("We have granted you the loan!")
			fallthrough
		case salary > 100000: fmt.Println("You were also chosen to obtain the loan without charging interest")
		default: fmt.Println("You don't meet any of our requirements")
		}
	}
}

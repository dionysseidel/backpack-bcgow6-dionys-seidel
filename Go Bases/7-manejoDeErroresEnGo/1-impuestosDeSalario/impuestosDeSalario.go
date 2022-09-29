/* Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”
Crea un error personalizado con un struct que implemente “Error()” con el mensaje
“error: el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000
Caso contrario, imprime por consola el mensaje “Debe pagar impuesto” */

package main

import (
	"fmt"
	"os"
)

type myCustomError struct {
	salary int
	msg    string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.salary, e.msg)
}

func validateTax(salary int) error {
	if salary < 150_000 {
		return &myCustomError{
			salary: salary,
			msg:    "error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}

func main() {
	const salary int = 150_000

	err := validateTax(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")
}

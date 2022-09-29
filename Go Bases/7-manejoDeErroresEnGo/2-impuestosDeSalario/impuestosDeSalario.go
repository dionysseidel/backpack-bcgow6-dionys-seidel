/* Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,
se implemente “errors.New()” */

package main

import (
	"errors"
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
	const salary int = 149_000

	err := validateTax(salary)
	if err != nil {
		fmt.Println(errors.New(err.Error()))
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")
}

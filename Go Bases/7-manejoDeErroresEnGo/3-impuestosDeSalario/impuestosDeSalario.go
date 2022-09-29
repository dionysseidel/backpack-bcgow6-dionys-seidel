/* Ejercicio 3 - Impuestos de salario #3
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”,
para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible
(el mensaje mostrado por consola deberá decir:
“error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”,
siendo [salary] el valor de tipo int pasado por parámetro) */

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
		// return fmt.Errorf("%v y el salario ingresado es de: %v", err, salary)
	}
	return nil
}

func main() {
	const salary int = 149_000

	err := validateTax(salary)
	if err != nil {
		errorr := fmt.Errorf("%w y el salario ingresado es de: %v", err, salary)
		fmt.Println(errorr)
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")
}

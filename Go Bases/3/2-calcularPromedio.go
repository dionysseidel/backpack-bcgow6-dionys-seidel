/* Ejercicio 2 - Calcular promedio
Un colegio necesita calcular el promedio (por alumno) de sus calificaciones
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros
y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo */

package main

import (
	"errors"
	"fmt"
)

func calculateAverage(studentGrades ...int) (average float64, err error) {
	sum, count := 0, 0
	for _, v := range studentGrades {
		count++
		if v < 0 {
			return 0, errors.New("One of the numers entered is negative")
		} else {
			sum += v
		}
	}
	average = float64(sum) / float64(count) // leng(studentGrades)
	return average, nil
}

func main() {
	fmt.Println(calculateAverage(9, 8, 7, 6))
	fmt.Println(calculateAverage(9, 8, -7, 6))
}

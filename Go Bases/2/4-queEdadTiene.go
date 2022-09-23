/* Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.
Por otro lado también es necesario: 
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa. */

package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	fmt.Println(employees["Benjamin"])

	count := 0
	for _, employee := range employees {
		if employee > 21 {
			count++
		}
	}
	fmt.Println("Employees over 21:", count)
}
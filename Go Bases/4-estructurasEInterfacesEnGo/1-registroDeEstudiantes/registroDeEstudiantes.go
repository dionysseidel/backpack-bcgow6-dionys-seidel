/* Una universidad necesita registrar a los/as estudiantes
y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha
y que tenga un método detalle */

package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumnos) Detalle(nombre, apellido, fecha string, dNI int) {
	a.Nombre = nombre
	a.Apellido = apellido
	a.DNI = dNI
	a.Fecha = fecha
	fmt.Println(a.Nombre)
	fmt.Println(a.Apellido)
	fmt.Println(a.DNI)
	fmt.Println(a.Fecha)
}

// func printStudentDetails() {
// 	fmt.Println(s1.Nombre)
// 	fmt.Println(s1.Apellido)
// 	fmt.Println(s1.DNI)
// 	fmt.Println(s1.Fecha)
// }

func main() {

}

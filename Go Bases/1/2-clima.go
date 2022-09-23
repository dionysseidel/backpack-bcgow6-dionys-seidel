/* Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares. 
Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
Imprime los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables? */

package main

import "fmt"

var (
	temperature float64 = 17
	humidity float64 = 44
	atmosphericPressure float64 = 1021
)

func main() {
	fmt.Printf("Temperature: %v, humidity: %v, atmospheric pressure: %v\n", temperature, humidity, atmosphericPressure)
}
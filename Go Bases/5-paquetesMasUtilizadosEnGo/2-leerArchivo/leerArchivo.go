/* Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado, para ello requiere que:
se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio,
la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50				*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./Go Bases/5-paquetesMasUtilizadosEnGo/1-guardarArchivo/purchasedProducts.txt")
	if err != nil {
		errMessage := fmt.Errorf("failed to read the file\n")
		fmt.Print(errMessage)
	}
	values := strings.Split(string(data), "\n")
	fmt.Printf("ID\tPrecio\tCantidad\n")
	for _, line := range values {
		items := strings.Split(line, ";")
		for _, i := range items {
			fmt.Printf("\t%v", i)
		}
		fmt.Print("\n")
	}
	// fmt.Printf("%s\n", data)
}

/* Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados,
separados por punto y coma (csv)
Debe tener el id del producto, precio y la cantidad
Estos valores pueden ser hardcodeados o escritos en duro en una variable */

package main

import (
	"fmt"
	"os"
)

type Product struct {
	ID       int
	Price    float64
	Quantity int
}

func (p Product) details() string {
	return fmt.Sprintf("%d;\n%f;\n%d\n", p.ID, p.Price, p.Quantity)
}

func saveToTxtFile(purchasedProduct Product) {

	fmt.Println(purchasedProduct.details())
	data := []byte(purchasedProduct.details())

	file, err := os.OpenFile("./Go Bases/5-paquetesMasUtilizadosEnGo/1-guardarArchivo/purchasedProducts.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	if _, err := file.Write(data); err != nil {
		panic(err)
	}

}

func main() {
	product1 := Product{
		ID:       1,
		Price:    89999,
		Quantity: 1,
	}

	product2 := Product{
		ID:       2,
		Price:    10499,
		Quantity: 1,
	}

	saveToTxtFile(product1)
	saveToTxtFile(product2)
	// Podría agrupar products e ir iterando
}

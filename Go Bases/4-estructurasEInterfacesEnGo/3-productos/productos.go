/* Ejercicio 3 - Productos
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto
Se requiere una función “nuevaTienda” que devuelva un Ecommerce
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda */

package main

import "fmt"

// Funcionalidad para administrar productos y retornar el valor del precio total

type tienda struct {
	products []Producto
}

func (t *tienda) Total() (totalPrice float64) {
	for _, product := range t.products {
		totalPrice += product.CalcularCosto()
	}
	return
}

func (t *tienda) Agregar(product Producto) {
	t.products = append(t.products, product)
}

type producto struct {
	productType string
	name        string
	price       float64
}

func (p producto) CalcularCosto() (totalPrice float64) {
	switch p.productType {
	case "Grande":
		return p.price*1.06 + 2500
	case "Mediano":
		return p.price * 1.03
	default:
		return p.price
	}
}

type Producto interface {
	CalcularCosto() (totalPrice float64)
}

type Ecommerce interface {
	Total() (totalPrice float64)
	Agregar(product Producto)
}

func nuevoProducto(productType, name string, price float64) producto {
	return producto{
		productType: productType,
		name:        name,
		price:       price,
	}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func main() {
	products := []producto{nuevoProducto("Pequeño", "Joystick ACCO Brands PowerA Enhanced Wired Controller for Xbox Series X|S zen purple", 21_099),
		nuevoProducto("Grande", "Bicicleta Ghepard Nordic X1.0 By Slp 29 21v Disco Suspension", 67_754.43),
		nuevoProducto("Pequeño",
			"Notebook EXO Smart XQH-S3182 plateada 15.6\", Intel Core i3 1115G4 8GB de RAM 256GB SSD, Intel UHD Graphics Xe G4 48EUs 1920x1080px Windows 11 Home",
			144_979),
		nuevoProducto("Mediano", "Freidora De Aire Digital Gadnic 1400w 4lts Conveccion + Piza", 31_299)}

	ecommerce := nuevaTienda()

	for _, product := range products {
		ecommerce.Agregar(product)
	}

	fmt.Printf("%.2f\n", ecommerce.Total())
}

/* Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios
Para ello requieren que tanto los usuarios como los productos
tengan la misma dirección de memoria en el main del programa como en las funciones
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos)
Producto: Nombre, precio, cantidad
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario
Borrar productos: recibe un usuario, borra los productos del usuario */

package main

import "fmt"

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func nuvevoProducto(nombre string, precio float64) (producto *Producto) {
	fmt.Println("Producto antes de la creación:", producto)
	return &Producto{
		Nombre: nombre,
		Precio: precio,
	}
}

func agregarProducto(usuario *Usuario, producto Producto, cantidad int) {
	fmt.Println("Usuario en la función, antes de la operación:", usuario)
	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, producto)
	fmt.Println("Usuario en la función, después de la operación:", usuario)
}

func borrarProductos(usuario *Usuario) {
	usuario.Productos = []Producto{}
}

func main() {
	usuario1 := &Usuario{
		Nombre:    "Mario Teodoro",
		Apellido:  "Escudero",
		Correo:    "marioteodoro.escudero@mercadolibre.com",
		Productos: []Producto{},
	}

	producto1 := nuvevoProducto("Cerveza Andes Origen Negra Lata X473cc", 170)

	agregarProducto(usuario1, *producto1, 4)
	fmt.Println("Producto después de la creación:", *producto1)

	fmt.Println("Productos que tiene el usuario en main:", *usuario1)

	borrarProductos(usuario1)
	fmt.Println("Productos que tiene el usuario en main al terminar la operatoria:", *usuario1)
}

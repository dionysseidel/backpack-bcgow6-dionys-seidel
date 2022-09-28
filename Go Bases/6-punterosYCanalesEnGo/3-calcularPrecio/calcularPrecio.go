/* Ejercicio 3 - Calcular Precio
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos
Debido a la fuerte demanda y para optimizar la velocidad requieren
que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad
Servicios: nombre, precio, minutos trabajados
Mantenimiento: nombre, precio

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad)
Sumar Servicios: recibe un array de servicio y devuelve el precio total
(precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora)
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3) */

package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type Service struct {
	Name          string
	Price         float64
	MinutesWorked int
}

type Maintenance struct {
	Name  string
	Price float64
}

// Calcular el precio total de Products, Services, y Maintenance

func sumProducts(products []Product, c chan float64) {
	totalPrice := 0.0
	for _, product := range products {
		totalPrice += product.Price * float64(product.Quantity)
	}
	c <- totalPrice
}

func sumServices(services []Service, c chan float64) {
	totalPrice := 0.0
	for _, service := range services {
		periodsWorked := 1
		if service.MinutesWorked > 30 {
			periodsWorked = service.MinutesWorked / 30
		}
		totalPrice += service.Price * float64(periodsWorked)
	}
	c <- totalPrice
}

func sumMaintenances(maintenances []Maintenance, c chan float64) {
	totalPrice := 0.0
	for _, maintenance := range maintenances {
		totalPrice += maintenance.Price
	}
	c <- totalPrice
}

func main() {
	products := []Product{{
		Name:     "Vino Malbec Raza Argentina Bicentenario - La Riojana",
		Price:    1_500,
		Quantity: 1,
	}, {
		Name:     "Yerba Mate Grapia Milenaria 1 Kg. Establecimiento Natural",
		Price:    890,
		Quantity: 1,
	}, {
		Name:     "Vaso Termico Jarro Acero Inoxidable Waterdog 500 Ml Ta500bg",
		Price:    10_326.67,
		Quantity: 2,
	}}

	services := []Service{{
		Name:          "Clases Manejo Moto Alquiler 110/150/200/400 Examen /ba/caba",
		Price:         500,
		MinutesWorked: 60,
	}, {
		Name:          "Clases De Inglés Y Español - Clases Particulares",
		Price:         900,
		MinutesWorked: 60,
	}, {
		Name:          "Alquiler Estudio Fotográfico Por Hora",
		Price:         1_250,
		MinutesWorked: 20,
	}}

	maintenances := []Maintenance{{
		Name:  "Pitch Technics Sl1200mk2 Servicio De Limpieza/mantenimiento",
		Price: 4_800,
	}, {
		Name:  "Alineación De Luces Lea Hasta El Final Trabajos C/garantía",
		Price: 2_500,
	}}

	c := make(chan float64)
	totalAmount := 0.0
	go sumProducts(products, c)
	go sumServices(services, c)
	go sumMaintenances(maintenances, c)
	for i := 0; i < 3; i++ {
		totalAmount += <-c
	}
	fmt.Println("Total amount:", totalAmount)
}

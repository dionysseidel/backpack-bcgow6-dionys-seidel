/* Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario
con funciones que vayan agregando información a la estructura
Para optimizar y ahorrar memoria requieren
que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña. */

package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

// Queremos que main y funciones hagan referencia al misme struct en memoria

func (u *Usuario) cambiarNombre(nombre, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
	fmt.Println("La dirección de memoria de u.Nombre es: ", &u.Nombre)
}

func (u *Usuario) cambiarEdad(edad int) {
	u.Edad = edad
}

func (u *Usuario) cambiarCorreo(correo string) {
	u.Correo = correo
}

func (u *Usuario) cambiarContraseña(contraseña string) {
	u.Contraseña = contraseña
}

func main() {
	usuario1 := &Usuario{
		Nombre:     "Dionys",
		Apellido:   "Seidel",
		Edad:       25,
		Correo:     "dionys.seidel@mercadolibre.com",
		Contraseña: "1234Abcd",
	}

	fmt.Println("La dirección de memoria de usuario1.Nombre es: ", &usuario1.Nombre)

	usuario1.cambiarNombre("Daniel", "Abila")
	usuario1.cambiarEdad(18)
	usuario1.cambiarCorreo("daniel.abila@mercadolibre.com")
	usuario1.cambiarContraseña("5678Efgh")
}

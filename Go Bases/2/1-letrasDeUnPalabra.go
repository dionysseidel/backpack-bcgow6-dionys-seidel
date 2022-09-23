/* Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. 
Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
Luego imprimí cada una de las letras. */

package main

import "fmt"

var word string = "arbotante"

func main() {
	var wordLength int = len(word)
	fmt.Println(wordLength)
	for i := 0; i < wordLength; i++ {
		fmt.Printf("%c\n", word[i])
	}	
}

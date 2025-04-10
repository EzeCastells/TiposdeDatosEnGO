/*
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE 754 32-bit floating-point numbers
float64     the set of all IEEE 754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
*/
package main

import (
	"fmt"
)

func main() {
	/*var nombrePersona string = "Eze"
	fmt.Println("Hola soy " + nombrePersona)
	var apellidoPersona = "Castells"
	fmt.Println("Mi apellido es " + apellidoPersona)
	segundoNombre := "Gastón"
	fmt.Println("Mi segundo nombre es " + segundoNombre)

	fmt.Println("------------ Trabajo con número ----------")

	var anioActual int16 = 2025
	fmt.Println("El año es", anioActual)

	var edad int8 = 40
	fmt.Println("Mi edad es", edad)*/

	//Uso arrays o slice (array sin definir una dimension)
	var array = [4]string{"Banana", "Manzana", "Naranja", "limon"}
	fmt.Println(array)

	array[0] = "Pera"
	fmt.Println(array)

	var arraySlice []string // Slice dinámico, sin tamaño fijo

	arraySlice = append(arraySlice, "Argentina")
	fmt.Println(arraySlice)

}

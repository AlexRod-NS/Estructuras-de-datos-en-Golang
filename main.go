package main

import (
	"./Estructuras"
)

func main() {
	lista := new(Estructuras.Lista)
	lista.Agregar(1)
	cola := new(Estructuras.Cola)
	cola.Encolar(2)
	pila := new(Estructuras.Pila)
	pila.Apilar(3)
}

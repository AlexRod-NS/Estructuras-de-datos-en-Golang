package Estructuras

import (
	"fmt"
)

type Nodo struct {
	dato      int
	siguiente *Nodo
}

type Lista struct {
	ptr *Nodo
}

func (self *Lista) Agregar(n int) {
	new_nodo := new(Nodo)
	new_nodo.dato = n
	aux := self.ptr
	if aux == self.ptr {
		self.ptr = new_nodo
	}
	new_nodo.siguiente = aux
}

func (self *Lista) Remover(val int) {
	aux := []*Nodo{nil, self.ptr}
	if self.ptr.dato == val {
		self.ptr = self.ptr.siguiente
	} else {
		for aux[1] != nil {
			if aux[1].dato == val {
				aux[0].siguiente = aux[1].siguiente
				aux[1].siguiente = nil
			}
			aux[0] = aux[1]
			aux[1] = aux[1].siguiente
		}
	}
}

func (self Lista) Indice(val int) int {
	aux := self.ptr
	cont := 0
	for aux != nil {
		if aux.dato == val {
			return cont
		}
		cont++
		aux = aux.siguiente
	}
	return -1
}

func (self Lista) Pop(indx int) {
	aux := []*Nodo{nil, self.ptr}
	cont := 0
	if indx == 0 {
		self.ptr = self.ptr.siguiente
	} else {
		for aux[1] != nil {
			if cont == indx {
				aux[0].siguiente = aux[1].siguiente
				aux[1].siguiente = nil
			}
			cont++
			aux[0] = aux[1]
			aux[1] = aux[1].siguiente
		}
	}
}

func (self *Lista) Elementos() {
	ptr := self.ptr
	for ptr != nil {
		fmt.Print(ptr.dato, " -> ")
		ptr = ptr.siguiente
	}
	fmt.Println("nil")
}

type Cola struct {
	inicio *Nodo
	fin    *Nodo
}

func (self *Cola) Desencolar() {
	aux := self.inicio
	self.inicio = aux.siguiente
	if self.inicio == nil {
		self.fin = nil
	}
	aux = nil
}

func (self *Cola) Encolar(n int) {
	new_nodo := new(Nodo)
	new_nodo.dato = n
	new_nodo.siguiente = nil
	if self.inicio == nil {
		self.inicio = new_nodo
	} else {
		self.fin.siguiente = new_nodo
	}
	self.fin = new_nodo
}

func (self Cola) Elementos() {
	aux := self.inicio
	for aux != nil {
		fmt.Print(aux.dato, "->")
		aux = aux.siguiente
	}
	fmt.Println("nil")
}

type Pila struct {
	ptr *Nodo
}

func (self *Pila) Apilar(n int) {
	new_nodo := new(Nodo)
	new_nodo.dato = n
	new_nodo.siguiente = self.ptr
	self.ptr = new_nodo
}

func (self *Pila) Desapilar() {
	aux := self.ptr
	self.ptr = aux.siguiente
	aux = nil
}

func (self *Pila) Vaciar() {
	for self.ptr != nil {
		self.Desapilar()
	}
}

func (self Pila) Elementos() {
	aux := self.ptr
	for aux != nil {
		fmt.Print(aux.dato, "->")
		aux = aux.siguiente
	}
	fmt.Println("nil")
}

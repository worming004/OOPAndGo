package main

import "fmt"

type idContainer struct {
	myValue int
}

func (cont idContainer) increment() {
	cont.myValue++ // N'a aucun effet en dehors de ce scope
}

func (cont *idContainer) ptrIncrement() {
	cont.myValue++ // Met à jour la référence derrière le pointer
}

func main() {
	// Instantiation d'une variable de type idContainer
	myContainer := idContainer{myValue: 0}

	myContainer.increment()
	fmt.Println(myContainer.myValue) // -> 0

	myContainer.ptrIncrement()
	fmt.Println(myContainer.myValue) // -> 1
}

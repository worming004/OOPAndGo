package main

import "fmt"

func main() {
	integer1 := 0         // type int
	integer2 := integer1  // type int
	fmt.Println(integer2) // -> 0

	pInteger := &integer1  // type *int
	(*pInteger)++          // increment de la valeur pointé
	fmt.Println(*pInteger) // -> 1
	fmt.Println(integer1)  // -> 1

	fmt.Print("\r\n")
	moreMain() // Pour approfondir les explications
}

func moreMain() {
	integer := 0        // signifie une variable dont la valeur est 1
	var pInteger *int   // signifie création d'un pointer d'un type int
	pInteger = &integer // signifie assignation de l'addresse d'integer dans le pointer

	fmt.Println(*pInteger) // imprime la valeur a l'adresse mémoire stocké par pInteger
	fmt.Println(pInteger)  // imprime la valeur de pInteger, autrement dit une adresse mémoire
	fmt.Println(&integer)  // imprime l'adresse mémoire d'integer
}

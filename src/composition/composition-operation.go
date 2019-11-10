package main

import "fmt"

type Operation interface {
	ApplyOperation(a int, b int) int
}

type PlusOperation struct{}

func (plus PlusOperation) ApplyOperation(a int, b int) int {
	return a + b
}

type MinusOperation struct{}

func (minus MinusOperation) ApplyOperation(a int, b int) int {
	return a - b
}

type OperationFacade struct {
	Plus  Operation
	Minus Operation
}

func NewOperationFacade() OperationFacade {
	return OperationFacade{
		Plus:  PlusOperation{},
		Minus: MinusOperation{},
	}
}

func main() {
	facade := NewOperationFacade()
	fmt.Println(facade.Plus.ApplyOperation(3, 1))
	fmt.Println(facade.Minus.ApplyOperation(3, 1))

	mainIoc()
}

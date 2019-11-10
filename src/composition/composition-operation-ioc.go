package main

import "fmt"

func NewOperationFacadeInversion(plus Operation, minus Operation) OperationFacade {
	return OperationFacade{
		Plus:  plus,
		Minus: minus,
	}
}

func mainIoc() {
	facade := NewOperationFacadeInversion(PlusOperation{}, MinusOperation{})
	fmt.Println(facade.Plus.ApplyOperation(3, 1))
	fmt.Println(facade.Minus.ApplyOperation(3, 1))
}

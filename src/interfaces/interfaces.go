package main

import "fmt"

// Implement Stringer interface
func (noise Noise) String() string {
	return noise.Name
}

type Noise struct {
	Name string
}

type IMakeNoise interface {
	MakeNoise() Noise
}

type Motor struct{}

// Implémente l'interface IMakeNoise
func (motor Motor) MakeNoise() Noise {
	return Noise{Name: "Vroom"}
}

type Car struct {
	motor Motor
}

func (car Car) MakeNoise() Noise {
	return car.motor.MakeNoise()
}

type Dog struct{}

func (dog Dog) MakeNoise() Noise {
	return Noise{Name: "Woof"}
}

func main() {
	var myNoiseMaker IMakeNoise = nil
	myNoiseMaker = Dog{}
	fmt.Println(myNoiseMaker.MakeNoise()) // -> Woof
	myNoiseMaker = Car{}
	fmt.Println(myNoiseMaker.MakeNoise()) // -> Vroom
}

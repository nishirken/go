package main

import (
	"fmt"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func getAnimal(animalName string) Animal {
	var cow = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	var bird = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	var snake = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	switch animalName {
	case "cow":
		return cow
	case "bird":
		return bird
	case "snake":
		return snake
	}
	return cow
}

func callMethod(animal *Animal, method string) {
	switch method {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func main() {
	var animalName string
	var methodName string

	for {
		fmt.Println(">")
		fmt.Scan(&animalName, &methodName)
		var animal = getAnimal(animalName)
		callMethod(&animal, methodName)
	}
}

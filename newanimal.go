package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var firstWord string
	var secondWord string
	var thirdWord string
	var animals = make(map[string]Animal)

	for {
		fmt.Println(">")
		fmt.Scan(&firstWord, &secondWord, &thirdWord)

		switch firstWord {
		case "newanimal":
			switch thirdWord {
			case "cow":
				animals[secondWord] = Cow{}
			case "bird":
				animals[secondWord] = Bird{}
			case "snake":
				animals[secondWord] = Snake{}
			}
		case "query":
			switch thirdWord {
			case "eat":
				animals[secondWord].Eat()
			case "move":
				animals[secondWord].Move()
			case "speak":
				animals[secondWord].Speak()
			}
		}
	}
}

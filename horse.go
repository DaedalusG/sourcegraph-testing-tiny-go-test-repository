package main

import "fmt"

type Animal interface {
	Sound() string
}

var _ Animal = Cat{}

type Cat struct{}

func (c Cat) Sound() string { return "i am a cat" }

type Dog struct{}

func (d Dog) Sound() string { return "it is i, the dog" }

func animalFarm() {
	makeSound(Cat{})
	makeSound(Dog{})
}

func makeSound(a Animal) {
	fmt.Printf("animal made a sound: %s", a.Sound())

	foo := a

	fmt.Printf("another animal: %s", foo.Sound())
}

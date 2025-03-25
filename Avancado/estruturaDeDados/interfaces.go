package estruturadedados

import (
	"fmt"
	bar "goAvancado/Bar"
	foo "goAvancado/Foo"
)

type Animal interface {
	Sound() string
}

type Dog struct{}

func (Dog) Sound() string {
	return "Au! Au!"
}

type Cat struct{}

func (d *Cat) Sound() string {
	//gatinho mineiro
	return "Miar! Miar!"
}

func takeAnimal(a Animal) {
	switch t := a.(type) {
	case *Dog:
		t.Sound()
	case *Cat:
		t.Sound()
	}
}

func (Dog) Interface() {
	fmt.Println("dog interface called")
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
}

func whatDoesThisInterface(i foo.Interface) {
	fmt.Println("Ne que funciona mesmo")
}

type Kayo struct{}

// que loucura realmente funciona
func (Kayo) String() string {
	return "esse e um teste"
}

func interfaces() {
	dog := Dog{}
	whatDoesThisAnimalSay(dog)

	whatDoesThisInterface(dog)
	// que loucura cara slckk
	bar.TakeFoo(dog)
}

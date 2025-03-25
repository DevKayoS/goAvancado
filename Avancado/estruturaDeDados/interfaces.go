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

func (Dog) Interface() {
	fmt.Println("dog interface called")
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Sound())
}

func whatDoesThisInterface(i foo.Interface) {
	fmt.Println("Ne que funciona mesmo")
}

func interfaces() {
	dog := Dog{}
	whatDoesThisAnimalSay(dog)

	whatDoesThisInterface(dog)
	// que loucura cara slckk
	bar.TakeFoo(dog)
}

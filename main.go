package main

type Animal interface {
	Sound() string
}

type Dog struct{}

func (d *Dog) Sound() string {
	return "Au! Au!"
}

func takeEmptyInterface(a any) {

}

func main() {
	var a Animal
	takeEmptyInterface("")
	takeEmptyInterface(a)
}

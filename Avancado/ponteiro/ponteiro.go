package ponteiro

import "fmt"

func ponteiroMain() {
	x := create()
	fmt.Println(*x)
}

func take(x *int) {
	*x = 100
}

func create() *int {
	x := 10
	return &x
}

func foo(x *int) {
	*x = 100
}

func main() {
	var x *int
	x = nil
	take(x)
	fmt.Println(x)
}

func takeTwo(x *int) {
	*x = 100
}

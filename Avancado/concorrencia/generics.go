package concorrencia

import "fmt"

type MeuTipo1 string

func (MeuTipo) Foo() {}

func concorrencia() {
	var mt MeuTipo = ""
	foo(mt)
}

type MyConstraint2 interface {
	Foo()
}

func foo2[T MyConstraint](arg T) {
	fmt.Println(arg)
}

func main() {
	foo(123)
	foo("")
	foo([]int{1, 2})

	var mt MeuTipo = ""

	foo(mt)

	// ms := MyStruct[string]{}

}

type MeuTipo string

type MyConstraint interface {
	int | ~string | []int
}

func foo[T MyConstraint](arg T) {
	fmt.Println(arg)
}

type MyStruct[T any] struct {
	Foo T
}

func (MyStruct[T]) foo(a T) {}

type Foo struct{}

func (Foo) fazer() {}

type Bar struct{}

func (Bar) fazer() {}

type MinhaInterface interface {
	Bar | Foo
	fazer()
}

func foo3[T MinhaInterface](arg T) {
	arg.fazer()
}

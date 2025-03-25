package foo

import "fmt"

type Foo struct {
	Name string
}

func (f Foo) Bar() {
	fmt.Println("vc esta no  bar")
}

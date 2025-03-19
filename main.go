package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	bar(slice)
	fmt.Println(slice)

}

func bar(slice []int) {
	slice[0] = 123
}

// fazendo validacao para ter uma perfomance melhor
func foo(slice []int) {
	_ = slice[3] //bounds check
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
}

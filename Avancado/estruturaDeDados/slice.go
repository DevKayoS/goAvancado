package estruturadedados

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	arr := [5]int{1, 2, 3, 4, 5}
	secondSlice := arr[:]
	fmt.Println(secondSlice, len(secondSlice), cap(secondSlice))

}

func arr() {
	arr := [5]int{1, 2, 3, 4, 5}

	slice := arr[1:4]
	arr[2] = 15
	slice[0] = 123
}

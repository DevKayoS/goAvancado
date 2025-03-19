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

func sliceV2() {
	movies := []string{
		"Inception",
		"Interstellar",
		"The Dark Knight",
		"Pulp Fiction",
		"Fight Club",
		"Forrest Gump",
		"The Matrix",
		"The Lord of the Rings",
		"The Godfather",
		"Back to the Future",
	}
	resultFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// matrix := [][]int{}
	// matrix3D := [][][]int{}

	filmes := make([]string, 0, len(resultFromApi))

	for _, id := range resultFromApi {
		filme := movies[id-1]
		fmt.Println(len(filmes), cap(filmes))
		filmes = append(filmes, filme)
	}
	fmt.Println(filmes)
}

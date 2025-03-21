package estruturadedados

import (
	"fmt"
	"math"
)

func maps() {
	var m map[string]string
	m1 := make(map[string]string)
	m2 := map[string]string{
		"Kayo":  "Pessoa",
		"Julia": "Pessoa",
	}

	m3 := map[string][]int{
		"Kayo": {1, 2, 3, 4},
	}

	fmt.Println(m3)

	fmt.Println(m, m1, m2)
}

func maps2() {
	m := make(map[string]string)
	m["Kayo"] = "Pessoa"
	valor, ok := m["Kayo"]

	fmt.Println(valor, ok)
	delete(m, "Kayo")
	valorRemovido, ok := m["Kayo"]
	fmt.Println(valorRemovido, ok)

	valor2, ok := m["foo"]

	fmt.Println(valor2, ok)

}

func limpar() {
	m := map[string]string{
		"Kayo":  "Dev",
		"Julia": "Professora",
	}

	clear(m)
	fmt.Println(m)
}

func mapFloat() {
	var f float64
	f = math.NaN()
	var f2 float64
	f2 = math.NaN()
	m := map[float64]string{
		f:  "Pessoa",
		f2: "Animal",
	}

	fmt.Println(f == f2)
	fmt.Println(m)

	delete(m, f)
	fmt.Println(m, f)

	clear(m)
	fmt.Println(m)
}

func iteracoes() {
	m := map[string]string{
		"Kayo":  "Dev",
		"Julia": "Professora",
	}

	for key := range m {
		if key == "Kayo" {
			delete(m, key)
		}
		fmt.Println(key, key)
	}
}

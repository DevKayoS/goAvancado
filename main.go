package main

func main() {
}

func Contains[T comparable](s []T, cmp T) bool {
	// slices.Contains()

	for _, str := range s {
		if str == cmp {
			return true
		}
	}
	return false
}

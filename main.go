package main

import (
	"io"
	"strings"
)

func main() {
	str := "hello world"
	reader := strings.NewReader(str)

	io.ReadAll(reader)
	io.ReadFull(reader, nil)

}

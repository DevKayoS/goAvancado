package readerwriter

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func mainV2() {
	str := "hello world"
	reader := strings.NewReader(str)
	writer := MyWritter{}

	buffer := make([]byte, 2)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		writer.Write(buffer[:n])
	}

}

type MyWritter struct {
	r io.Writer
}

func (MyWritter) Write(b []byte) (int, error) {
	fmt.Print(string(b))
	return len(b), nil
}

func (mw MyWritter) Write21(b []byte) (int, error) {
	for i, bb := range b {
		b[i] = bb + 10
	}
	return mw.r.Write(b)
}

func readerExample1() {
	str := "hello world"
	reader := strings.NewReader(str)
	buffer := make([]byte, 2)
	n, err := reader.Read(buffer)

	if err != nil {
		panic(err)
	}

	fmt.Println(n)
	fmt.Println(string(buffer[:n]))

	fmt.Println(buffer)
}

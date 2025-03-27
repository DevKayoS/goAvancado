package main

import (
	"errors"
	"fmt"
)

func zelo() {
	user, err := NewUser(true)

	if err != nil {
		fmt.Println("algum erro na hora de criar novo usuarui")
		return
	}

	user.Foo()
}

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("nao pode dividir por zero")
	}

	return a / b, nil
}

type User struct {
	foo string
}

func (u User) Foo() {
	fmt.Println(u.foo)
}

func NewUser(wantErr bool) (*User, error) {
	if wantErr {
		return nil, errors.New("um novo erro")
	}

	return &User{}, nil
}

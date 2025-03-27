package main

import (
	"errors"
	"fmt"
	"math"
)

func zelo() {
	user, err := NewUser(true)

	if err != nil {
		fmt.Println("algum erro na hora de criar novo usuarui")
		return
	}

	user.Foo()
}

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string { return s.msg }

var ErrNotFound = errors.New("not found")

func raizQuadrada(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{"Nao existe raiz quadrada de numero negativo"}
	}
	resultado := math.Sqrt(x)
	return resultado, nil
}

func foo() error { return SqrtError{msg: "teste"} }

func mainReserva() {
	x := -10

	res, err := raizQuadrada(float64(x))

	if err != nil {
		fmt.Println(err)
		return
	}

	err = foo()

	if err != nil && errors.Is(err, ErrNotFound) {
		fmt.Println("deu erro not found melhore apenas")
		return
	}
	fmt.Println("Ih o erro nem entrou")

	fmt.Println(res)

	var sqtrError SqrtError

	if err != nil && errors.As(err, &sqtrError) {
		fmt.Println(sqtrError.msg)
		return
	}

	fmt.Println("meteu marcha")
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

func a() error { return Errqualquer }
func b() error { return ErrQualquer2 }

func f() error {
	var erroResult error

	if err := a(); err != nil {
		erroResult = errors.Join(erroResult, err)
	}

	if err := b(); err != nil {
		erroResult = errors.Join(erroResult, err)
	}

	return erroResult
}

var (
	Errqualquer  = errors.New("error")
	ErrQualquer2 = errors.New("error 2")
)

func mainzadaReservada() {
	err := foo()

	if err != nil && errors.Is(err, Errqualquer) {
		fmt.Println("deu erro: ", err)
		return
	}
}

func foo21() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("deu erro em foo: %w", err)
	}

	return nil
}

func bar() error {
	return Errqualquer
}

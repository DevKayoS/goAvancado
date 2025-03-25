package estruturadedados

import (
	"encoding/json"
	"fmt"
	foo "goAvancado/Foo"
)

type MinhaString string

type User struct {
	foo.Foo
	Name string
	ID   uint64
}

func (u User) PrintName() {
	fmt.Println(u.Name)
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func UpdateName(u *User, newName string) {
	u.Name = newName
}

// nao ta liberado essa bagunca
// func (f foo.Foo) GetFoo() {

// }

func structs() {
	user := User{Name: "kayo vinicius",
		ID: 10958,
	}
	empUser := User{
		Name: "Usuario empresa",
	}

	user.Bar()

	fmt.Println(user.ID)
	empUser.PrintName()
	empUser.UpdateName("ze da manga")
	empUser.PrintName()
	UpdateName(&empUser, "Ai caliquinh")
	empUser.PrintName()

	res, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

}

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name  string `json:"name" db:"user_name"`
	Email string `json:"email,omitempty" db:"email"`
	age   int
}

func main() {
	u := User{Name: "Ada"}

	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("field=%s exported=%v json=%q db=%q\n",
			f.Name, f.IsExported(),
			f.Tag.Get("json"), f.Tag.Get("db"),
		)
	}

	data, _ := json.Marshal(u)
	fmt.Println("json:", string(data)) // {"name":"Ada"}  (Email omitted, age unexported)
}

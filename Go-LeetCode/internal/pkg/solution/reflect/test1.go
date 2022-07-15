package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("Eat")
}

func main() {
	a := Animal{}
	reflect.ValueOf(&a).MethodByName("Eat").Call([]reflect.Value{})
}

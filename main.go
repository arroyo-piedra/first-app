package main

import (
	"fmt"
	struct "github.com/arroyo-piedra/first-app/Struct"
)

func main() {
	fmt.Println("Hello World")
	p := struct.newPerson("Juan")
	fmt.Println(p.name)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type character interface {
	attack() string
	defense() string
}

type knight struct {
	swordType   string
	swordWeight int
}

func (k knight) attack() string {
	return "Use a sword"
}
func (k knight) defense() string {
	return "Use a shield"
}

type archer struct {
	bow string
}

func (a archer) attack() string {
	return "Use a bow"
}
func (a archer) defense() string {
	return "Dodge"
}

func characterMoves(c character) {
	fmt.Println(c)
	fmt.Println(c.attack())
	fmt.Println(c.defense())
}

func main() {
	knight := knight{swordType: "normal", swordWeight: 5}
	archer := archer{bow: "wood"}

	fmt.Println("This knight: ")
	characterMoves(knight)
	fmt.Println("This archer: ")
	characterMoves(archer)

	count("cat")
	count("dog")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheeps")
		wg.Done()
	}()

	wg.Wait()

}

func count(animal string) {
	for i := 1; true; i++ {
		fmt.Println(i, animal)
		time.Sleep(time.Millisecond * 500)
	}
}

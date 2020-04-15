/**                   _
 *  _             _ _| |_
 * | |           | |_   _|
 * | |___  _   _ | | |_|
 * | '_  \| | | || | | |
 * | | | || |_| || | | |
 * |_| |_|\___,_||_| |_|
 *
 * (c) Huli Inc
 */

package main

import (
	"fmt"
	"sync"
	"time"

	character "github.com/arroyo-piedra/first-app/src/struct_example"
)

func main() {
	knight := character.Knight{swordType: "normal", swordWeight: 5}
	archer := character.Archer{bow: "wood"}
	fmt.Print(example())
	fmt.Println("This knight: ")
	characterMoves(knight)
	fmt.Println("This archer: ")
	CharacterMoves(archer)

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

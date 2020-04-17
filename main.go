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
)

func main() {

	// Var declaration, arrays  and slices

	var number int = 5

	fmt.Println(number)

	/*
		arrayExample := [6]int{2, 3, 5, 7, 11, 13}

		for index, element := range arrayExample {
			fmt.Println(index, "=>", element)
		}
	*/
	/*
		var s []int = arrayExample[1:4]
		fmt.Println(s)
	*/
	// Struct examples

	/*knight := character.Knight{SwordType: "normal", SwordWeight: 5}
	archer := character.Archer{Bow: "wood"}
	fmt.Println("This knight: ")
	character.CharacterMoves(knight)
	fmt.Println("This archer: ")
	character.CharacterMoves(archer)*/

	// Gorotuines examples

	/*go count("cat")
	count("dog")

		// go rutines counter
		var wg sync.WaitGroup
		wg.Add(1)

		// anonimus function
		go func() {
			count("sheeps")
			wg.Done()
		}()

		wg.Wait()*/

	/*array := []int{7, 2, 8, -9, 4, 0}

	channel := make(chan int)
	go sum(array[:len(array)/2], channel)
	go sum(array[len(array)/2:], channel)
	last3, first3 := <-channel, <-channel // receive from c

	fmt.Println(last3, first3, last3+first3)*/

}

func count(animal string) {
	for i := 1; true; i++ {
		fmt.Println(i, animal)
		//time.Sleep(time.Millisecond * 500)
	}
}

func sum(array []int, channel chan int) {
	sum := 0
	for _, value := range array {
		sum += value
	}
	channel <- sum // send sum to c
}

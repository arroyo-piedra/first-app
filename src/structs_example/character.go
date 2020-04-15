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

package structs_example

import "fmt"

// interface that define two "methods"
type character interface {
	attack() string
	defense() string
}

type Knight struct {
	swordType   string
	swordWeight int
}

func (k Knight) Attack() string {
	return "Use a sword"
}
func (k Knight) Defense() string {
	return "Use a shield"
}

type Archer struct {
	bow string
}

func (a Archer) Attack() string {
	return "Use a bow"
}
func (a Archer) Defense() string {
	return "Dodge"
}

// With this function you only need to pass a character
// and the interface is in charge of use the corresponding function
func CharacterMoves(c character) {
	fmt.Println(c)
	fmt.Println(c.attack())
	fmt.Println(c.defense())
}

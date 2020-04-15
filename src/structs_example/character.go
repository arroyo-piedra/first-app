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

package struct_example

import "fmt"

// interface that define two "methods"
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

// With this function you only need to pass a character
// and the interface is in charge of use the corresponding function
func characterMoves(c character) {
	fmt.Println(c)
	fmt.Println(c.attack())
	fmt.Println(c.defense())
}

package structs

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	k := new(Knight)
	fmt.Println(k.Attack())
}

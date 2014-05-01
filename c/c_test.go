package c

import (
	"fmt"
	"github.com/fudanchii/a/b"
)

func TestBbb() {
	a := true
	b.Bbb()
	fmt.Println("Ccc")
	if !a {
		println("false assertion")
	}
}

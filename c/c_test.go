package c

import (
	"fmt"
	"github.com/fudanchii/A/b"

	"testing"
)

func TestBbb(t *testing.T) {
	a := true
	b.Bbb()
	fmt.Println("Ccc")
	if !a {
		t.Error("false assertion")
	}
}

package b_test

import (
	"fmt"
	"github.com/fudanchii/a/b"
	"testing"
)

func TestBbb(t *testing.T) {
	a := true
	b.Bbb()
	fmt.Println("Ccc")
	if !a {
		println("false assertion")
		t.Errorf("false assertion")
	}
}

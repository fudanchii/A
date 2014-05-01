package b_test

import (
	"testing"
	"github.com/fudanchii/a/b"
)

func TestBbb(t *testing.T) {
	a := true
	b.Bbb()
	println("Ccc")
	if (!a) {
		t.Errorf("false assertion")
	}
}

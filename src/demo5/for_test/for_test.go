package for_test

import "testing"

func TestFor(t *testing.T) {
	n := 0
	for n<5 {
		t.Log(n)
		n++
	}
}

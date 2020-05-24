package constant_test

import "testing"

const (
	a = 1 + iota
	b
	c
)

func TestConstantTry(t *testing.T) {
	t.Log(a, b, c)
}
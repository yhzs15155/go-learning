package string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)

	s = "hello"
	t.Log(len(s))

	for _,v := range s {
		t.Logf("%c", v)
	}

}
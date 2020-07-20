package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	s := "A,B,C"

	parts := strings.Split(s, ",")

	t.Log(parts)
	for _, v := range parts {
		t.Log(v)
	}

	t.Log(strings.Join(parts, "-"))
}


func TestStringConv(t *testing.T) {
	t.Log(strconv.Atoi("10"))
}
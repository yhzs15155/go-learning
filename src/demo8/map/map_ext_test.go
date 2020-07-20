package map_ext

import "testing"


func TestMapWithFunValue(t *testing.T) {
	m := map[string] func (op int) int {}

	m["a"] = func(op int) int { return op }
	m["b"] = func(op int) int { return op * op }
	m["d"] = func(op int) int { return op * op * op }

	t.Log(m["a"](2), m["b"](2), m["d"](2))
}


func TestMapForSet(t *testing.T) {
	mySet := map[int] bool {}

	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}

	mySet[3] = true
	t.Log(len(mySet))

	delete(mySet, 1)

	t.Log(len(mySet))
}

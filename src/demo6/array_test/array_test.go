package array_test

import "testing"

func TestArray(t *testing.T) {
	var arr [3] int
	arr[0] = 1

	arr1 := [4] int {1,2,3,4}
	arr2 := [...] int {1,2,3,4,5,6}

	t.Log(arr[0], arr[1])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [...] int {1,2,3,4,5,6}

	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for idx, e := range arr {
		t.Log(idx, e)
	}

	for _, e := range arr {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...] int {1,2,3,4,5,6}

	arr_sec := arr[1:3]
	arr_sec = arr[:3]
	arr_sec = arr[3:]
	arr_sec = arr[3:len(arr)]

	t.Log(arr_sec)
}
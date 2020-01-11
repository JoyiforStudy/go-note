package v1

import "testing"

func Add(x, y int) int {
	return x + y
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	excepted := 4
	if sum != excepted {
		t.Errorf("excepted '%d' but god '%d' ", excepted, sum)
	}
}

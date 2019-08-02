package dollar

import "testing"

func TestMultiplication(t *testing.T) {
	five := Dollar{amount: 5}
	five.Times(2)
	if expected, actual := 10, five.amount; expected != actual {
		t.Errorf("five.Times(2).amount wont %d but got %d", expected, actual)
	}
}

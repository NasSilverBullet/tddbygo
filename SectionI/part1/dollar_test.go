package dollar

import "testing"

func TestMultiplication(t *testing.T) {
	five := Dollar{
		ammount: 5,
	}
	five.times(2)
	if expected, actual := 10, five.ammount; expected != actual {
		t.Errorf("five.amount wont %d but got %d", expected, actual)
	}
}

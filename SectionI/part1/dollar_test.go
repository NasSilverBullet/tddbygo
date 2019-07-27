package dollar

import "testing"

func TestMultiplication(t *testing.T) {
	five := Dollar{
		Ammount: 5,
	}
	five.times(2)
	if expected, actual := 10, five.ammount; expected != actual {
		t.Errorf("five.amount wont %s but got %s", expected, actual)
	}
}

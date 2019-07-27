package dollar

import "testing"

func TestMultiplication(t *testing.T) {
	five := Dollar{
		ammount: 5,
	}
	product := five.times(2)
	if expected, actual := 10, product.ammount; expected != actual {
		t.Errorf("five.amount wont %d but got %d", expected, actual)
	}
	product = five.times(3)
	if expected, actual := 15, product.ammount; expected != actual {
		t.Errorf("five.amount wont %d but got %d", expected, actual)
	}

}

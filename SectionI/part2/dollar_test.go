package dollar

import "testing"

func TestMultiplication(t *testing.T) {
	five := Dollar{amount: 5}
	product := five.Times(2)
	if expected, actual := 10, product.amount; expected != actual {
		t.Errorf("five.Times(2).amount wont %d but got %d", expected, actual)
	}
	product = five.Times(3)
	if expected, actual := 15, product.amount; expected != actual {
		t.Errorf("five.Times(3).amount wont %d but got %d", expected, actual)
	}

}

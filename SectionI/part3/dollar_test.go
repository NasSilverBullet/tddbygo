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

func TestEquality(t *testing.T) {
	d := &Dollar{
		ammount: 5,
	}
	if expected, actual := true, d.equals(&Dollar{ammount: 5}); expected != actual {
		t.Errorf("Dollar{5}.equals(Dollar{5} wont %t but got %t", expected, actual)
	}
	if expected, actual := false, d.equals(&Dollar{ammount: 6}); expected != actual {
		t.Errorf("Dollar{5}.equals(Dollar{5} wont %t but got %t", expected, actual)
	}
}

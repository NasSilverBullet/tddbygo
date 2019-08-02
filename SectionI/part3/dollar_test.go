package dollar

import "testing"

func TestMultiplication(t *testing.T) {
	five := &Dollar{amount: 5}
	product := five.Times(2)
	if expected, actual := 10, product.amount; expected != actual {
		t.Errorf("five.Times(2).amount wont %d but got %d", expected, actual)
	}
	product = five.Times(3)
	if expected, actual := 15, product.amount; expected != actual {
		t.Errorf("five.Times(2).amount wont %d but got %d", expected, actual)
	}

}

func TestEquality(t *testing.T) {
	d := &Dollar{amount: 5}
	if expected, actual := true, d.Equals(&Dollar{amount: 5}); expected != actual {
		t.Errorf("&Dollar{amount: 5}.Equals(&Dollar{amount: 5}) wont %t but got %t", expected, actual)
	}
	if expected, actual := false, d.Equals(&Dollar{amount: 6}); expected != actual {
		t.Errorf("&Dollar{amount: 5}.Equals(&Dollar{amount: 6}) wont %t but got %t", expected, actual)
	}
}

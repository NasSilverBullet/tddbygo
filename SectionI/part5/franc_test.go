package dollar

import (
	"testing"
)

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	if expected, actual := NewFranc(10), five.times(2); *expected != *actual {
		t.Errorf("product wont %v but got %v", expected, actual)
	}
	if expected, actual := NewFranc(15), five.times(3); *expected != *actual {
		t.Errorf("product wont %v but got %v", expected, actual)
	}

}

func TestFrancEquality(t *testing.T) {
	d := NewFranc(5)
	if expected, actual := true, d.equals(NewFranc(5)); expected != actual {
		t.Errorf("Dollar{5}.equals(Dollar{5} wont %t but got %t", expected, actual)
	}
	if expected, actual := false, d.equals(NewFranc(6)); expected != actual {
		t.Errorf("Dollar{5}.equals(Dollar{5} wont %t but got %t", expected, actual)
	}
}

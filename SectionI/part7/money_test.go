package money

import "testing"

func TestEquality(t *testing.T) {
	if expected, actual := true, NewDollar(5).Equals(NewDollar(5)); expected != actual {
		t.Errorf("NewDollar(5).Equals(NewDollar(5)) wont %t but got %t", expected, actual)
	}
	if expected, actual := false, NewDollar(5).Equals(NewDollar(6)); expected != actual {
		t.Errorf("NewDollar(5).Equals(NewDollar(6)) wont %t but got %t", expected, actual)
	}
	if expected, actual := true, NewFranc(5).Equals(NewFranc(5)); expected != actual {
		t.Errorf("NewFranc(5).Equals(NewFranc(5)) wont %t but got %t", expected, actual)
	}
	if expected, actual := false, NewFranc(5).Equals(NewFranc(6)); expected != actual {
		t.Errorf("NewFranc(5).Equals(NewFranc(6)) wont %t but got %t", expected, actual)
	}
	// 現状 fail になる
	if expected, actual := false, NewFranc(5).Equals(NewDollar(5)); expected != actual {
		t.Errorf("NewFranc(5).Equals(NewDollar(5)) wont %t but got %t", expected, actual)
	}
}

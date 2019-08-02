package money

import "testing"

func TestPairEquality(t *testing.T) {
	if expected, actual := true, makePair("USD", "USD").equals(makePair("USD", "USD")); expected != actual {
		t.Errorf("makePair(\"USD\", \"USD\").equals(makePair(\"USD\", \"USD\")) wont %t but got %t", expected, actual)
	}
	if expected, actual := false, makePair("USD", "USD").equals(makePair("USD", "CHF")); expected != actual {
		t.Errorf("makePair(\"USD\", \"USD\").equals(makePair(\"USD\", \"CHF\")) wont %t but got %t", expected, actual)
	}
}

func TestIdentityRate(t *testing.T) {
	if expected, actual := 1, NewBank().Rate("USD", "USD"); expected != actual {
		t.Errorf("NewDollar(5).Equals(NewDollar(5)) wont %d but got %d", expected, actual)
	}
}

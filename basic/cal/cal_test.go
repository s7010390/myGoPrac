package cal

import (
	"testing"
)

func TestSumOfTest(t *testing.T) {
	given := 5
	want := 15

	get := SumOffirst(given)
	if get != want {
		t.Errorf("given %d want %d but %d", given, want, get)
	}
}

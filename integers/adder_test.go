package integers_test

import (
	"testing"

	"github.com/infernalslam/TDD_Golang/integers"
)

func TestAdder(t *testing.T) {
	sum := integers.Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestAdderIsEqual5(t *testing.T) {
	sum := integers.Add(3, 2)
	expected := 5

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

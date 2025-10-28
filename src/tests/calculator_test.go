package tests

import (
	"testing"
	"workshop-template-go/src/utils"
)

func TestAdd(t *testing.T) {
	// Case 1 : numbers > 0
	result := utils.Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; attendu %d", result, expected)
	}

	// Case 2 : numbers < 0
	result = utils.Add(-1, -1)
	expected = -2
	if result != expected {
		t.Errorf("Add(-1, -1) = %d; attendu %d", result, expected)
	}

	// Case 3 : numbers with a 0
	result = utils.Add(10, 0)
	expected = 11
	if result != expected {
		t.Errorf("Add(10, 0) = %d; attendu %d", result, expected)
	}
}
package main

import "testing"

func TestIsContainsDuplicate(t *testing.T) {
	slice1 := []string{"a", "b", "c", "a"}
	shouldReturnTrue := isContainsDuplicate(slice1)
	if shouldReturnTrue != true {
		t.Errorf("isContainsDuplicate(slice1) = %v, want true", shouldReturnTrue)
	}

	slice2 := []string{"a", "b", "c", "d"}
	shouldReturnFalse := isContainsDuplicate(slice2)
	if shouldReturnFalse != false {
		t.Errorf("isContainsDuplicate(slice2) = %v, want false", shouldReturnFalse)
	}
}


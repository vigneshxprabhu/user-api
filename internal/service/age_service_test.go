package service

import "testing"

func TestCalculateAge(t *testing.T) {

	age, err := CalculateAge("2000-01-01")

	if err != nil {
		t.Fatal(err)
	}

	if age <= 0 {
		t.Errorf("expected positive age, got %d", age)
	}
}

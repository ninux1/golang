package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("expected pointer to singleton after calling GetInstance(), not nil")
	}
	expectedCounter := counter1
}
currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After Calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
}

package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance()
	if instance1 == nil {
		t.Error("expected pointer to singleton after calling GetInstance(), not nil")
	}

	calculateadd(instance1)
	fmt.Println("The value of the counter of first instance is ", instance1.count)

	instance2 := GetInstance()
	if instance2 == nil {
		t.Error("expected pointer to singleton after calling GetInstance(), not nil")
	}
	calculateadd(instance2)
	fmt.Println("The value of the counter of second instance is ", instance2.count)

}

func calculateadd(i Singleton) {
	i.AddOne()
}

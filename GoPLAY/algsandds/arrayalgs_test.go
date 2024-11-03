package algsandds

import "testing"

func TestIsUnique(t *testing.T) {
	text := "hello"

	isUnique := IsUniqueString(&text)

	if isUnique == true {
		t.Errorf("String '%v' does not have unique characters. Expected false but got %v", text, true)
	}
}

func TestUniqueForUnique(t *testing.T) {

	textUnique := "help"
	isUniqueDos := IsUniqueString(&textUnique)
	if isUniqueDos != true {
		t.FailNow()
	}

}

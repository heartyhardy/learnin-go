package linkedlist

import "testing"

//Test Append method of LinkedList
func TestAppend(t *testing.T) {
	ll := new(LinkedList)
	expected := &LinkedList{end: &Node{Value: 1, next: nil}}
	expected.Length = 1
	actual := ll.Append(1)

	if actual.end.Value != actual.start.Value {
		t.Error("Head and Start does not point to the same Node")
	}
	if (*expected.end != *actual.end) || (expected.Length != actual.Length) {
		t.Error("Elements not Equal")
	}
}

//Test Prepend method of LinkedList
func TestPrepend(t *testing.T) {
	ll := new(LinkedList)
	expected := &LinkedList{end: nil, start: &Node{Value: 2, next: nil}}
	expected.Length = 1
	actual := ll.Prepend(2)

	if actual.end.Value != actual.start.Value {
		t.Error("Head and Start does not point to the same Node")
	}
	if (*expected.start != *actual.start) || (expected.Length != actual.Length) {
		t.Error("Elements not Equal")
	}
}

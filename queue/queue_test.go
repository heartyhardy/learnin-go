package queue

import (
	"fmt"
	"testing"
)

func TestEnqueue(t *testing.T) {
	n := Node{Value: "Hello"}
	q := &Queue{}
	//expected := &Queue{End: &Node{Value: "Hello"}, Length: 1}
	q.Enqueue(n)

	tx := q.Dequeue()
	fmt.Println((tx.Value).(string)[:2])

	//fmt.Printf("%v %T", fmt.Sprintf("%v", actual.end.Value), fmt.Sprintf("%v", actual.end.Value))
	// if expected.Length != actual.Length {
	// 	t.Error("Lens are different!")
	// }

	// if actual.end.Value != expected.end.Value {
	// 	t.Error("Ends are not equal!")
	// }

	// if actual.start == nil {
	// 	t.Error("Start pointer not pointing correctly!")
	// }

}

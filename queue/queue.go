package queue

//Node struct
type Node struct {
	Value interface{}
	Next  *Node
}

//Queue struct
type Queue struct {
	Start  *Node
	End    *Node
	Length int
}

//Enqueuer interface
type Enqueuer interface {
	Enqueue(interface{}) *Queue
}

//Dequeuer interface
type Dequeuer interface {
	Dequeue(interface{}) *Queue
}

//Enqueue an element to the Queue
func (q *Queue) Enqueue(val interface{}) *Queue {
	node := new(Node)
	node.Value = val
	node.Next = nil
	if q.End != nil {
		t := q.End
		t.Next = node
		q.End = node
	} else {
		q.End = node
		q.Start = node
	}
	q.Length++
	return q
}

//Dequeue an item from queue
func (q *Queue) Dequeue() Node {
	if q.Start != nil {
		t := q.Start
		q.Start = t.Next
		q.Length--
		t1 := (t.Value).(Node)
		return t1
	}
	return Node{}
}

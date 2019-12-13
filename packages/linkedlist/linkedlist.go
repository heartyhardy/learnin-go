package linkedlist

//Node struct
type Node struct {
	Value interface{}
	next  *Node
}

//LinkedList struct
type LinkedList struct {
	start  *Node
	end    *Node
	Length int
}

//Appender interface
type Appender interface {
	Append(interface{}) *LinkedList
}

//Prepender interface
type Prepender interface {
	Prepend(interface{}) *LinkedList
}

//Append an element to the LinkedList
func (l *LinkedList) Append(val interface{}) *LinkedList {
	node := new(Node)
	node.Value = val
	node.next = nil
	if l.end != nil {
		t := l.end
		t.next = node
		l.end = node
	} else {
		l.end = node
		l.start = node
	}
	l.Length++
	return l
}

//Prepend an element to the LinkedList
func (l *LinkedList) Prepend(val interface{}) *LinkedList {
	node := new(Node)
	node.Value = val
	node.next = nil
	if l.start != nil {
		node.next = l.start
		l.start = node
	} else {
		l.start = node
		l.end = node
	}
	l.Length++
	return l
}

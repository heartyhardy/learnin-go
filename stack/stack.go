package stack

//Node holds the values in the LinkedList
type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

//Stack ...
type Stack struct {
	Start  *Node
	End    *Node
	length int
}

//Push ...
func (s *Stack) Push(value interface{}) *Stack {
	node := new(Node)
	node.Value = value
	node.Next = nil
	node.Prev = nil

	if s.length != 0 {
		node.Next = s.Start
		s.Start = node
	}
	s.Start = node
	s.End = node
	s.length++
	return s
}

//Peek the Start of the Stack
func (s *Stack) Peek() *Node {
	return s.Start
}

//PeekVal Peeks the value of the Start
func (s *Stack) PeekVal() Node {
	return (s.Start.Value).(Node)
}

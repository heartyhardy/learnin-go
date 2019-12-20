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

	if s.length > 0 {
		t := s.Start
		node.Next = s.Start
		s.Start = node
		t.Prev = node

		if s.End.Prev == nil {
			s.End.Prev = t
		}
	} else {
		s.Start = node
		s.End = node
		s.End.Prev = nil
	}
	s.length++
	return s
}

//Peek the Start of the Stack
func (s *Stack) Peek() *Node {
	return s.Start
}

//PeekVal Peeks the value of the Start
func (s *Stack) PeekVal() *Node {
	return (s.Start.Value).(*Node)
}

//PeekEnd returns the End of the Stack
func (s *Stack) PeekEnd() *Node {
	return s.End
}

//PeekEndVal returns the value of the End
func (s *Stack) PeekEndVal() *Node {
	return (s.End.Value).(*Node)
}

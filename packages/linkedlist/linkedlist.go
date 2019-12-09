package linkedlist

//Node works as an element for linkedlist
type Node struct {
	Value interface{}
	Next  *Node
}

//LinkedList : Basic structure
type LinkedList struct {
	start  *Node
	end    *Node
	length int
}

//LinkedLister interface forthe linkedlist
type LinkedLister interface {
	Append(*Node)
	Prepend(*Node)
	Peek() *Node
}

//Append Node into LinkedList
func (list *LinkedList) Append(newNode *Node) {
	if list.length == 0 {
		list.end = newNode
	} else {
		current := list.end
		current.Next = newNode
		list.end = newNode
	}
	list.length++
}

//Peek at the end element
func (list *LinkedList) Peek() *Node {
	if list.length > 0 {
		return list.end
	}
	return nil
}

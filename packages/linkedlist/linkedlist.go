package linkedlist

//Node works as an element for linkedlist
type Node struct {
	Value interface{}
	Next  *Node
}

//LinkedList : Basic structure
type LinkedList struct {
	end    *Node
	length int
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

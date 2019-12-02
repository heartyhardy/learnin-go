package linkedlist

//Robot ...
type Robot struct {
	ID   int
	Name string
	next *Robot
}

//Next Robot
func (robot *Robot) Next() *Robot {
	return robot.next
}

//Robots ...
type Robots struct {
	length int
	head   *Robot
	tail   *Robot
}

//Append a robot
func (rb *Robots) Append(newrobot *Robot) {
	if rb.length == 0 {
		rb.head = newrobot
		rb.tail = newrobot
	} else {
		lastRobot := rb.tail
		lastRobot.next = newrobot
		rb.tail = newrobot

	}
	rb.length++
}

//Get Robot
func (rb *Robots) Get() *Robot {
	return rb.head
}

package main

import (
	"testing"
)

func TestTurning(t *testing.T) {

	rb := new(Robot)
	rb.direction = new(Vector)
	rb.location = new(Vector)
	panel := new(Panel)
	rb.panels = make(map[Panel]int)
	rb.messages = make([]int, 0)
	rb.location.x = 0
	rb.location.y = 0
	rb.direction.x = 1
	rb.direction.y = 0
	panel.x = rb.location.x
	panel.y = rb.location.y
	rb.panels[*panel] = 0

	turn := 0
	expected := &Vector{0, 1}
	done := true
	actual, ok := rb.Turn(turn)

	if (expected.x != actual.x || expected.y != actual.y) || !done || !ok {
		t.Errorf("Expected direction to be %d but instead got %v!", expected, actual)
	}
}

func TestMoving(t *testing.T) {

	rb := new(Robot)
	rb.direction = new(Vector)
	rb.location = new(Vector)
	panel := new(Panel)
	rb.panels = make(map[Panel]int)
	rb.messages = make([]int, 0)
	rb.location.x = 0
	rb.location.y = 0
	rb.direction.x = 0
	rb.direction.y = 1
	panel.x = rb.location.x
	panel.y = rb.location.y
	rb.panels[*panel] = 0

	rb.Move(0)
	expectedDirection := &Vector{-1, 0}
	expectedLocation := &Vector{-1, 0}
	rb.Move(0)
	expectedDirection = &Vector{0, -1}
	expectedLocation = &Vector{-1, -1}
	rb.Move(0)
	expectedDirection = &Vector{1, 0}
	expectedLocation = &Vector{0, -1}
	rb.Move(1)
	expectedDirection = &Vector{0, -1}
	expectedLocation = &Vector{0, -2}
	rb.Move(1)
	expectedDirection = &Vector{-1, 0}
	expectedLocation = &Vector{-1, -2}

	if (*rb.direction != *expectedDirection) ||
		(*rb.location != *expectedLocation) || (len(rb.panels) != 6) {
		t.Error("Test Failed")
	}

}

func TestPainting(t *testing.T) {

	rb := new(Robot)
	rb.direction = new(Vector)
	rb.location = new(Vector)
	panel := new(Panel)
	rb.panels = make(map[Panel]int)
	rb.messages = make([]int, 0)
	rb.location.x = 0
	rb.location.y = 0
	rb.direction.x = 0
	rb.direction.y = 1
	panel.x = rb.location.x
	panel.y = rb.location.y
	rb.panels[*panel] = 0

	p := &Panel{0, 0}
	rb.Paint(1)

	if rb.panels[*p] != 1 || rb.state != moveReady {
		t.Error("Test Failed")
	}

}

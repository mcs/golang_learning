package adt

import "testing"

func TestEmpty(t *testing.T) {
	s := NewStack()

	if s.Empty() != true {
		t.Error("Stack was not empty")
	}
}

func TestNotEmpty(t *testing.T) {
	s := NewStack()
	s.Push("Bob")

	if s.Empty() != false {
		t.Error("Stack was empty")
	}
}

func TestSizeNone(t *testing.T) {
	s := NewStack()

	if s.Size() != 0 {
		t.Errorf("Expected 0, got: %d", s.Size())
	}
}

func TestSizeOne(t *testing.T) {
	s := NewStack()
	s.Push("Bob")

	if s.Size() != 1 {
		t.Errorf("Expected 1, got: %d", s.Size())
	}
}

func TestSizeSome(t *testing.T) {
	s := NewStack()
	s.Push("Bob")
	s.Push("Alice")

	if s.Size() != 2 {
		t.Errorf("Expected 2, got: %d", s.Size())
	}
}

func TestPopOneFromOne(t *testing.T) {
	s := NewStack()
	s.Push("Bob")

	item, err := s.Pop()
	if item != "Bob" {
		t.Errorf("Expected Bob, got: %s", item)
	}
	if err != nil {
		t.Errorf("Expected nil, got: %s", err)
	}
	if s.Size() != 0 {
		t.Errorf("Expected 0, got: %d", s.Size())
	}
}

func TestPopTwoFromTwo(t *testing.T) {
	s := NewStack()
	s.Push("Alice")
	s.Push("Bob")

	item, err := s.Pop()
	if item != "Bob" {
		t.Errorf("Expected Bob, got: %s", item)
	}
	if err != nil {
		t.Errorf("Expected nil, got: %s", err)
	}
	if s.Size() != 1 {
		t.Errorf("Expected 1, got: %d", s.Size())
	}

	item, _ = s.Pop()
	if item != "Alice" {
		t.Errorf("Expected Alice, got: %s", item)
	}
	if err != nil {
		t.Errorf("Expected nil, got: %s", err)
	}
	if s.Size() != 0 {
		t.Errorf("Expected 0, got: %d", s.Size())
	}
}

func TestPushPopPushPop(t *testing.T) {
	s := NewStack()
	s.Push("Alice")
	_, _ = s.Pop()
	s.Push("Bob")
	item, err := s.Pop()
	if item != "Bob" {
		t.Errorf("Expected Bob, got: %s", item)
	}
	if err != nil {
		t.Errorf("Got unexpected error: %s", err)
	}
}

func TestPopFromEmpty(t *testing.T) {
	s := NewStack()

	_, err := s.Pop()
	if err == nil {
		t.Error("Expected error, got: nil")
	}
	if err.Error() != "cannot pop empty stack" {
		t.Error("Got unexpected error: ", err)
	}

}

package stack

import "sync"

// Stack is a simple stack implementation
type Stack struct {
	lock sync.Mutex
	s    []interface{}
}

// NewStack creates a new stack
func NewStack() *Stack {
	return &Stack{s: make([]interface{}, 0)}
}

// Push pushes a new element onto the stack
func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

// Pop pops the top element from the stack
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.s) == 0 {
		return nil
	}

	last := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]

	return last
}

// Peek returns the top element from the stack without removing it
func (s *Stack) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.s) == 0 {
		return nil
	}

	return s.s[len(s.s)-1]
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.s)
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.s) == 0
}

// Clear clears the stack
func (s *Stack) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = make([]interface{}, 0)
}

// ToSlice returns the stack as a slice
func (s *Stack) ToSlice() []interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.s
}

// FromSlice sets the stack from a slice
func (s *Stack) FromSlice(slice []interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = slice
}

// Copy returns a copy of the stack
func (s *Stack) Copy() *Stack {
	s.lock.Lock()
	defer s.lock.Unlock()

	newStack := NewStack()
	newStack.s = append(newStack.s, s.s...)

	return newStack
}

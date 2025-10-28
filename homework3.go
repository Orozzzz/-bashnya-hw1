package main

import "fmt"

func main() {
	defer fmt.Println("== COMPLETED ==")
	defer fmt.Println()
	
	stack := NewStack()

	
	stack.Push(10)   //add elements
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	fmt.Println("Size of stack: ", stack.Size())
    fmt.Println("Is the stack empty: ",stack.IsEmpty())

	fmt.Println("Pop: ", stack.Pop())
	fmt.Println("Pop: ", stack.Pop())
	fmt.Println("Pop: ", stack.Pop())
	//fmt.Println("Pop: ", stack.Pop())

	fmt.Println("Size after Pop: ", stack.Size())

	stack.Clear()   //clear stack
	fmt.Println("is the empty after Clear: ", stack.IsEmpty())
	
}

type Stack struct {
	items []interface{} 
}

func NewStack() *Stack {
	return  &Stack{
		items: make([]interface{}, 0),
	}
}

func (s *Stack) Push(value interface{}) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return value
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = make([]interface{},0)

}
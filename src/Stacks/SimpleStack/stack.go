package main

import "fmt"

const MAX_SIZE int = 3;

type stack struct {
	top int
	array [MAX_SIZE]int
}

// Takes the value at the end of the array out
func (s *stack) pop() int {

	if s.top > 0 {
		s.top--;
		temp := s.array[s.top];
		s.array[s.top] = 0;
		return temp;
	} else {
		return 0;
	}
}

// Puts the values of the at the position of top
func (s *stack) push(num int) {
	if s.top > MAX_SIZE-1 {
		return;
	}

	s.array[s.top] = num;
	s.top++;
	return;
}

func createStack() *stack {

	s := new(stack);

	s.top = 0;

	return s;
}

func main() {

	stack := createStack();

	stack.push(10);
	stack.push(20);
	stack.push(30);
	stack.push(40);
	fmt.Println(stack.pop());
}

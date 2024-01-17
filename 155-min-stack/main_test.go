package min_stack

import "testing"

func TestMinStackGetMin(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	t.Log(minStack.GetMin())
	t.Log(minStack)
	minStack.Pop()
	t.Log(minStack)
	minStack.Pop()
	t.Log(minStack)
	t.Log(minStack.GetMin())
}

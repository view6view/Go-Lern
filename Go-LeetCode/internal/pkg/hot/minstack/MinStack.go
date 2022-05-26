package minstack

import "container/list"

type MinStack struct {
	stack    *list.List
	minStack *list.List
}

func Constructor() MinStack {
	return MinStack{
		stack:    list.New(),
		minStack: list.New(),
	}
}

func (this *MinStack) Push(val int) {
	this.stack.PushBack(val)
	if this.minStack.Len() == 0 || val <= this.minStack.Back().Value.(int) {
		this.minStack.PushBack(val)
	}
}

func (this *MinStack) Pop() {
	if this.stack.Remove(this.stack.Back()).(int) == this.minStack.Back().Value.(int) {
		this.minStack.Remove(this.minStack.Back())
	}
}

func (this *MinStack) Top() int {
	return this.stack.Back().Value.(int)
}

func (this *MinStack) GetMin() int {
	return this.minStack.Back().Value.(int)
}

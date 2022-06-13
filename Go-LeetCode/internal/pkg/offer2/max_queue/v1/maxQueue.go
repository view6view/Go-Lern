package v1

import "container/list"

type MaxQueue struct {
	queue *list.List
	deque *list.List
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue: list.New(),
		deque: list.New(),
	}
}

func (this *MaxQueue) Max_value() int {
	if this.deque.Len() > 0 {
		return this.deque.Front().Value.(int)
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	this.queue.PushBack(value)
	for this.deque.Len() > 0 && this.deque.Back().Value.(int) < value {
		this.deque.Remove(this.deque.Back())
	}
	this.deque.PushBack(value)
}

func (this *MaxQueue) Pop_front() int {
	if this.queue.Len() == 0 {
		return -1
	}
	if this.queue.Front().Value.(int) == this.deque.Front().Value.(int) {
		this.deque.Remove(this.deque.Front())
	}
	return this.queue.Remove(this.queue.Front()).(int)
}

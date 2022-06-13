package v2

type MaxQueue struct {
	queue []int
	deque []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.deque) == 0 {
		return -1
	}
	return this.deque[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	for len(this.deque) > 0 && this.deque[len(this.deque)-1] < value {
		this.deque = this.deque[:len(this.deque)-1]
	}
	this.deque = append(this.deque, value)
}

func (this *MaxQueue) Pop_front() int {
	l1 := len(this.queue)
	if l1 == 0 {
		return -1
	}
	peek := this.queue[0]
	this.queue = this.queue[1:l1]
	if l2 := len(this.deque); l2 > 0 && peek == this.deque[0] {
		this.deque = this.deque[1:l2]
	}
	return peek
}

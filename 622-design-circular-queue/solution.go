package design_circular_queue

type Node struct {
	Val  int
	Next *Node
}

type MyCircularQueue struct {
	front, rear *Node
	size        int
	capacity    int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{capacity: k}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.size++

	node := &Node{Val: value}

	if q.IsEmpty() {
		q.front, q.rear = node, node
		node.Next = q.front
		return true
	}

	q.rear.Next = node
	node.Next = q.front
	q.rear = node
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.size--

	if q.front == q.rear {
		q.front, q.rear = nil, nil
		return true
	}

	q.front = q.front.Next
	q.rear.Next = q.front
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.front == nil {
		return -1
	}
	return q.front.Val
}

func (q *MyCircularQueue) Rear() int {
	if q.rear == nil {
		return -1
	}
	return q.rear.Val
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.front == nil
}

func (q *MyCircularQueue) IsFull() bool {
	return q.size == q.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

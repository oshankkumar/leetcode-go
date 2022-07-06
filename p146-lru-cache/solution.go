package p146

type Node struct {
	Key, Val   int
	Next, Prev *Node
}

type List struct {
	Head, Tail *Node
}

func (l *List) MoveToFront(node *Node) {
	if node == l.Head {
		return
	}

	if node == l.Tail {
		l.Tail = node.Prev
	}

	node.Prev.Next = node.Next
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Prev = nil
	node.Next = l.Head
	l.Head.Prev = node
	l.Head = node
}

func (l *List) PushToFront(node *Node) {
	if l.Head == nil {
		l.Head, l.Tail = node, node
		return
	}
	node.Next = l.Head
	l.Head.Prev = node
	l.Head = node
}

func (l *List) DelLastNode() *Node {
	tail := l.Tail
	if l.Head == l.Tail {
		l.Head, l.Tail = nil, nil
		return tail
	}

	l.Tail.Prev.Next = nil
	l.Tail = l.Tail.Prev
	tail.Prev = nil
	return tail
}

type LRUCache struct {
	keyAddr  map[int]*Node
	capacity int
	list     List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{keyAddr: make(map[int]*Node), capacity: capacity}
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.keyAddr[key]
	if !ok {
		return -1
	}

	l.list.MoveToFront(node)
	return node.Val
}

func (l *LRUCache) Put(key int, value int) {
	node, ok := l.keyAddr[key]
	if ok {
		node.Val = value
		l.list.MoveToFront(node)
		return
	}

	if len(l.keyAddr) >= l.capacity {
		node := l.list.DelLastNode()
		delete(l.keyAddr, node.Key)
	}

	node = &Node{Key: key, Val: value}
	l.keyAddr[node.Key] = node
	l.list.PushToFront(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

package lfu_cache

import (
	"container/heap"
)

type Item struct {
	Key          int
	Val          int
	Index        int
	Freq         int
	LastAccessAt int
}

type ItemHeap []*Item

func (h ItemHeap) Len() int {
	return len(h)
}

func (h ItemHeap) Less(i, j int) bool {
	if h[i].Freq == h[j].Freq {
		return h[i].LastAccessAt < h[j].LastAccessAt
	}
	return h[i].Freq < h[j].Freq
}

func (h ItemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

func (h *ItemHeap) Push(x any) {
	item, _ := x.(*Item)
	item.Index = len(*h)
	*h = append(*h, item)
}

func (h *ItemHeap) Pop() any {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}

type LFUCache struct {
	cap           int
	store         map[int]*Item
	heap          *ItemHeap
	timeStampFunc func() int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:   capacity,
		store: make(map[int]*Item),
		heap:  new(ItemHeap),
		timeStampFunc: func() func() int {
			var a int
			return func() int {
				a++
				return a
			}
		}(),
	}
}

func (l *LFUCache) Get(key int) int {
	item, ok := l.store[key]
	if !ok {
		return -1
	}

	item.Freq++
	item.LastAccessAt = l.timeStampFunc()
	heap.Fix(l.heap, item.Index)
	return item.Val
}

func (l *LFUCache) Put(key int, value int) {
	item, ok := l.store[key]
	if ok {
		item.Freq++
		item.LastAccessAt = l.timeStampFunc()
		item.Val = value
		l.store[key] = item
		heap.Fix(l.heap, item.Index)
		return
	}

	if len(l.store) == l.cap {
		item, _ := heap.Pop(l.heap).(*Item)
		delete(l.store, item.Key)
	}

	item = &Item{
		Key:          key,
		Val:          value,
		Freq:         1,
		LastAccessAt: l.timeStampFunc(),
	}
	heap.Push(l.heap, item)
	l.store[item.Key] = item
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

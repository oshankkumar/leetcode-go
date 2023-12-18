package maximumproductoftwoelementsinanarray

import "container/heap"

var _ = maxProduct

func leftChildIdx(i int) int {
	return 2*i + 1
}

func rightChildIdx(i int) int {
	return 2*i + 2
}

func parentIdx(i int) int {
	return (i - 1) / 2
}

func newMinHeap(c int) *minHeap {
	return &minHeap{cap: 2, nums: make([]int, c)}
}

type minHeap struct {
	nums []int
	size int
	cap  int
}

func (m *minHeap) heapifyUp(i int) {
	pIdx := parentIdx(i)
	if m.nums[i] < m.nums[pIdx] {
		m.nums[i], m.nums[pIdx] = m.nums[pIdx], m.nums[i]
		m.heapifyUp(pIdx)
	}
}

func (m *minHeap) heapify(i int) {
	smallest := i
	l, r := leftChildIdx(i), rightChildIdx(i)

	if l < m.size && m.nums[l] < m.nums[smallest] {
		smallest = l
	}

	if r < m.size && m.nums[r] < m.nums[smallest] {
		smallest = r
	}

	if smallest != i {
		m.nums[i], m.nums[smallest] = m.nums[smallest], m.nums[i]
		m.heapify(smallest)
	}
}

func (m *minHeap) deleteRoot() {
	m.nums[0], m.nums[m.size-1] = m.nums[m.size-1], m.nums[0]
	m.size--
	m.heapify(0)
}

func (m *minHeap) insert(n int) {
	if m.size == m.cap {
		if n < m.nums[0] {
			return
		}
		m.deleteRoot()
		m.insert(n)
		return
	}

	m.nums[m.size] = n
	m.heapifyUp(m.size)
	m.size++
}

func maxProduct(nums []int) int {
	h := newMinHeap(2)
	for _, n := range nums {
		h.insert(n)
	}
	return (h.nums[0] - 1) * (h.nums[1] - 1)
}

type MaxHeap []int

func (m MaxHeap) Len() int { return len(m) }

func (m MaxHeap) Less(i, j int) bool { return m[i] > m[j] }

func (m MaxHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *MaxHeap) Push(x any) { *m = append(*m, x.(int)) }

func (m *MaxHeap) Pop() any {
	old := *m
	v := old[len(old)-1]
	*m = old[:len(old)-1]
	return v
}

func MaxProduct(nums []int) int {
	data := MaxHeap(nums)
	heap.Init(&data)
	return (heap.Pop(&data).(int) - 1) * (heap.Pop(&data).(int) - 1)
}

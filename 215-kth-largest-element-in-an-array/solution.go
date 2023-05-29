package kthlargestelementinanarray

func findKthLargest(nums []int, k int) int {
	hp := newMinHeap(k)
	for i := 0; i < k; i++ {
		hp.insert(nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] < hp.root() {
			continue
		}

		hp.deleteRoot()
		hp.insert(nums[i])
	}

	return hp.root()
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func newMinHeap(l int) *minHeap {
	return &minHeap{heap: make([]int, l), maxLen: l}
}

type minHeap struct {
	heap   []int
	maxLen int
	len    int
}

func (m *minHeap) root() int {
	return m.heap[0]
}

func (m *minHeap) insert(n int) {
	if m.len == m.maxLen {
		return
	}

	m.len++
	m.heap[m.len-1] = n
	m.heapifyUp(m.len - 1)
}

func (m *minHeap) heapifyUp(i int) {
	if i > 0 && m.heap[i] < m.heap[parent(i)] {
		m.heap[parent(i)], m.heap[i] = m.heap[i], m.heap[parent(i)]
		m.heapifyUp(parent(i))
	}
}

func (m *minHeap) heapifyDown(i int) {
	if i < 0 || i >= m.len {
		return
	}
	l, r := leftChild(i), rightChild(i)

	minIdx := i
	if l < m.len && m.heap[l] < m.heap[minIdx] {
		minIdx = l
	}

	if r < m.len && m.heap[r] < m.heap[minIdx] {
		minIdx = r
	}

	if minIdx != i {
		m.heap[i], m.heap[minIdx] = m.heap[minIdx], m.heap[i]
		m.heapifyDown(minIdx)
	}

}

func (m *minHeap) deleteRoot() {
	if m.len <= 0 {
		return
	}
	m.heap[m.len-1], m.heap[0] = m.heap[0], m.heap[m.len-1]
	m.len--
	m.heapifyDown(0)
}

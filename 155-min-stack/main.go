package min_stack

type MinStack struct {
	arr  []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		arr:  make([]int, 0),
		mins: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, val)
		return
	}
	this.mins = append(this.mins, Min(val, this.mins[len(this.mins)-1]))
}

func (this *MinStack) Pop() {
	if len(this.arr) == 0 {
		return
	}
	this.arr = this.arr[:len(this.arr)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

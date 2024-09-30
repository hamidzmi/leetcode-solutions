package main

import "fmt"

type CustomStack struct {
	items []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{make([]int, 0, maxSize)}
}

func (this *CustomStack) Push(x int) {
	if len(this.items) == cap(this.items) {
		return
	}
	this.items = append(this.items, x)
}

func (this *CustomStack) Pop() int {
	if len(this.items) == 0 {
		return -1
	}

	value := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]

	return value
}

func (this *CustomStack) Increment(k int, val int) {
	if k > len(this.items) {
		k = len(this.items)
	}
	for i := 0; i < k; i++ {
		this.items[i] += val
	}
}

func main() {
	obj := Constructor(3)
	fmt.Println(obj)

	obj.Push(1)
	fmt.Println(obj)

	obj.Push(2)
	fmt.Println(obj)

	obj.Pop()
	fmt.Println(obj)

	obj.Push(2)
	fmt.Println(obj)

	obj.Push(3)
	fmt.Println(obj)

	obj.Push(4)
	fmt.Println(obj)

	obj.Increment(5, 100)
	fmt.Println(obj)

	obj.Increment(2, 100)
	fmt.Println(obj)
}

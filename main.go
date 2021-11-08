package main

import (
	"container/heap"
	"math"
)

type MaxHeap struct {
	maxHeap []int
	target  float64
}

func (h MaxHeap) Len() int {
	return len(h.maxHeap)
}

// Max Heap
// ABS(i - target) > ABS(j - target)
func (h MaxHeap) Less(i, j int) bool {
	return math.Abs(float64(h.maxHeap[i])-h.target) > math.Abs(float64(h.maxHeap[j])-h.target)
}

func (h MaxHeap) Swap(i, j int) {
	h.maxHeap[i], h.maxHeap[j] = h.maxHeap[j], h.maxHeap[i]
}

func (h *MaxHeap) Push(i interface{}) {
	h.maxHeap = append(h.maxHeap, i.(int))
}

func (h *MaxHeap) Pop() interface{} {
	last := h.maxHeap[len(h.maxHeap)-1]
	h.maxHeap = h.maxHeap[:len(h.maxHeap)-1]
	return last
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	cur := root

	// stack is used to track work
	stack := []*TreeNode{}

	// heap tracks result
	// initialzie it with the target
	hp := MaxHeap{target: target}

	// loop until the stack is empty and cur is nil
	for cur != nil || len(stack) > 0 {
		// this is an in order traversal
		if cur == nil {
			cur = stack[len(stack)-1]

			// work
			heap.Push(&hp, cur.Val)
			if hp.Len() > k {
				heap.Pop(&hp)
			}

			cur = cur.Right

			// pop stack item
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, cur)
			cur = cur.Left
		}
	}

	// return the max heap
	return hp.maxHeap
}

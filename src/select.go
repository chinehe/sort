package src

import (
	"fmt"
)

// Select 选择排序
type Select struct {
}

func (s *Select) Sort(arr []int) {
	// 长度小于2直接返回
	if len(arr) < 2 {
		return
	}
	// 选择排序，共需要N-1轮
	var maxIndex int
	for i := 0; i < len(arr)-1; i++ {
		// 找出剩余元素集中最大的元素
		maxIndex = 0
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		// 将最大元素与剩余元素集中的最大元素交换
		arr[maxIndex], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[maxIndex]
		fmt.Printf("第%d轮 -> %v\n", i, arr)
	}
}

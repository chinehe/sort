package src

import (
	"fmt"
)

// Bubble 冒泡排序
type Bubble struct {
}

func (b *Bubble) Sort(arr []int) {
	// 长度小于2直接返回
	if len(arr) < 2 {
		return
	}
	// 冒泡排序，共需要N-1轮
	for i := 0; i < len(arr)-1; i++ {
		// 每轮需要比较所有剩余元素
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Printf("%d -> %v\n", i, arr)
	}
}

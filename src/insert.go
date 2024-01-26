package src

import (
	"fmt"
)

// Insert 插入排序
type Insert struct {
}

func (i *Insert) Sort(arr []int) {
	// 长度小于2直接返回
	if len(arr) < 2 {
		return
	}
	// 插入排序，共需要N-1轮
	for i := 0; i < len(arr)-1; i++ {
		// 从后往前遍历已排序的部分
		for j := i; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				// 如果当前元素比要排序的元素大，就交换位置
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				// 如果当前元素小于等于要排序的元素，就跳出循环
				break
			}
		}

		fmt.Printf("第%d轮 -> %v\n", i, arr)
	}
}

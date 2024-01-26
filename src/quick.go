package src

import (
	"fmt"
)

// Quick 快速排序
type Quick struct {
}

func (q *Quick) Sort(arr []int) {
	// 长度小于2，直接返回
	if len(arr) < 2 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

// quickSort 快速排序
func quickSort(arr []int, leftIndex, rightIndex int) {
	if leftIndex >= rightIndex {
		fmt.Printf("执行区间[%d %d] 直接返回\n", leftIndex, rightIndex)
		return
	}
	// 取第一个数字为Key
	key := arr[leftIndex]
	keyIndex := leftIndex
	// 开始快速排序
	left, right := leftIndex, rightIndex
	for left < right {
		// 从后往前找第一个小于Key的数
		for left < right {
			if arr[right] >= key {
				right--
			} else {
				// 找到后,将找到为数放到Key的位置
				arr[keyIndex] = arr[right]
				keyIndex = right
				break
			}
		}
		// 从前往后找第一个大于Key的数
		for left < right {
			if arr[left] <= key {
				left++
			} else {
				// 找到后,将找到为数放到Key的位置
				arr[keyIndex] = arr[left]
				keyIndex = left
				break
			}
		}
	}
	// 将Key放入最后的位置
	arr[keyIndex] = key
	fmt.Printf("执行区间[%d %d]:%v\n", leftIndex, rightIndex, arr)
	// 将数组拆成两部分，分别排序
	quickSort(arr, leftIndex, keyIndex-1)
	quickSort(arr, keyIndex+1, rightIndex)
}

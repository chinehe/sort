package src

import (
	"fmt"
)

// Shell 希尔排序
type Shell struct {
}

func (s *Shell) Sort(arr []int) {
	// 长度小于2，直接返回
	if len(arr) < 2 {
		return
	}
	// 计算起始增量
	var increment = len(arr)
	for increment > 1 {
		// 计算增量
		increment /= 2
		// 根据增量将数组拆分成子序列,i表示每个子序列的起始元素索引
		for i := 0; i < increment; i++ {
			// 从子序列的第一个元素遍历到倒数第二个元素
			for j := i; j < len(arr)-increment; j += increment {
				// 从后往前遍历已经排序的子序列，将子序列的下一个元素插入有序部分
				for k := j; k >= 0; k -= increment {
					if arr[k] > arr[k+increment] {
						// 如果当前元素大于要插入的元素，就交换位置
						arr[k], arr[k+increment] = arr[k+increment], arr[k]
					} else {
						// 如果当前元素小于要插入的元素，就跳出循环
						break
					}
				}
			}
			fmt.Printf("增量=%d,第%d轮 -> %v\n", increment, i, arr)
		}
	}
}

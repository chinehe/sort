package src

import (
	"math"
)

// Radix 基数排序
type Radix struct {
}

func (r *Radix) Sort(arr []int) {
	res := make([]int, len(arr))
	index := 0
	// 找到数组中的最大值
	max := arr[0]
	for _, a := range arr {
		if a > max {
			max = a
		}
	}

	// 计算最大值的位数
	maxDigits := int(math.Log10(float64(max))) + 1

	// 根据最大位数进行循环排序
	for i := 1; i <= maxDigits; i++ {
		// 创建十个桶
		buckets := make([][]int, 10)
		for b, _ := range buckets {
			buckets[b] = make([]int, 0)
		}

		// 将元素分配到桶中
		for _, a := range arr {
			digit := int(math.Log10(float64(a))) + 1
			if digit != i {
				continue
			}
			buckets[digit] = append(buckets[digit], a)
		}

		// 对每个桶中的元素进行排序
		// 此处采用直接插入排序
		for b, _ := range buckets {
			new(Insert).Sort(buckets[b])
		}

		// 合并桶中的元素
		for ii, _ := range buckets {
			for jj, _ := range buckets[ii] {
				res[index] = buckets[ii][jj]
				index++
			}
		}
	}
	copy(arr, res)
}

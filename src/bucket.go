package src

import (
	"math"
)

// Bucket 桶排序
type Bucket struct {
}

func (b *Bucket) Sort(arr []int) {
	if len(arr) < 2 {
		return
	}
	// 计算最大值和最小值
	max, min := arr[0], arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}

	// 桶的数量，此处取平方根
	bucketSize := int(math.Sqrt(float64(len(arr))))
	// 初始化桶
	buckets := make([][]int, bucketSize)
	for i, _ := range buckets {
		buckets[i] = make([]int, 0)
	}

	// 计算每个桶需要存储的范围
	bucketRange := (max-min)/bucketSize + 1

	// 将元素分配到桶中
	for _, a := range arr {
		index := (a - min) / bucketRange
		buckets[index] = append(buckets[index], a)
	}

	// 对每个桶的元素进行排序，这里采用直接插入排序
	sorter := new(Insert)
	for i, _ := range buckets {
		sorter.Sort(buckets[i])
	}

	// 合并桶的元素
	index := 0
	for i, _ := range buckets {
		for j, _ := range buckets[i] {
			arr[index] = buckets[i][j]
			index++
		}
	}
}

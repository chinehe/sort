package src

// Heap 堆排序
type Heap struct {
}

func (h *Heap) Sort(arr []int) {
	// 构建大顶堆
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr))
	}
	// 调整堆结构+交换堆顶元素与末尾元素
	for j := len(arr) - 1; j > 0; j-- {
		// 交换对堆顶元素与末尾元素
		arr[0], arr[j] = arr[j], arr[0]
		adjustHeap(arr, 0, j)
	}
}

// 调整大顶堆
func adjustHeap(arr []int, index, length int) {
	// 取出当前元素
	curr := arr[index]
	// 从当前节点的左子节点开始
	for i := index*2 + 1; i < length; i = i*2 + 1 {
		// 如果左子节点小于右子节点
		if i+1 < length && arr[i] < arr[i+1] {
			// 指向右子节点
			i++
		}
		// 如果子节点大于父节点
		if arr[i] > curr {
			// 将子节点值赋值给父节点
			arr[index] = arr[i]
			index = i
		} else {
			break
		}
	}
	// 将当前值放入最终的位置
	arr[index] = curr
}

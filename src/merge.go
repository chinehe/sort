package src

// Merge 归并排序
type Merge struct {
}

func (m *Merge) Sort(arr []int) {
	if len(arr) < 2 {
		return
	}
	mid := len(arr) / 2
	m.Sort(arr[:mid])
	m.Sort(arr[mid:])
	merge(arr[:mid], arr[mid:])
}

// merge 归并
func merge(arrA, arrB []int) {
	res := make([]int, len(arrA)+len(arrB))
	indexA, indexB, indexR := 0, 0, 0

	for indexA < len(arrA) && indexB < len(arrB) {
		if arrA[indexA] > arrB[indexB] {
			res[indexR] = arrB[indexB]
			indexB++
			indexR++
		} else {
			res[indexR] = arrA[indexA]
			indexA++
			indexR++
		}
	}

	for i := indexA; i < len(arrA); i++ {
		res[indexR] = arrA[i]
		indexR++
	}

	for i := indexB; i < len(arrB); i++ {
		res[indexR] = arrB[i]
		indexR++
	}

	copy(arrA, res[0:len(arrA)])
	copy(arrB, res[len(arrA):])
}

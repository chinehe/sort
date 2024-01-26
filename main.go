package main

import (
	"fmt"
	"github.com/chinehe/sort/src"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().Unix())
	N := 8
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(200)
	}

	fmt.Println(arr)

	new(src.Radix).Sort(arr)

	fmt.Println(arr)
}

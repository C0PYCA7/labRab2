package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//linear(10)
	parallel(10, 2)
}

func createArr(n int) ([]int, []int) {
	arr := make([]int, n)
	arrFind := make([]int, n, n)

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 2
	}
	return arr, arrFind
}

func linear(n int) {
	t := time.Now()
	arr, arrFind := createArr(n)
	//fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]%arr[i] == 0 {
				arrFind[j] = 1
			} else if arr[j]%arr[i] != 0 {
				if arrFind[j] == 1 {
					continue
				}
				arrFind[j] = 0
			}
		}
	}
	fmt.Println(arrFind)
	for idx, value := range arrFind {
		if value == 0 {
			fmt.Print(arr[idx], " ")
		}
	}
	fmt.Println()
	elapsedTime := time.Since(t)
	fmt.Println(elapsedTime)
}

func Parallel(n, numWorkers int) {
	t := time.Now()
	arr, arrFind := createArr(n)

	var wg sync.WaitGroup
	partSize := n / numWorkers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j < end; j++ {
				for k := 0; k < len(arr); k++ {
					if arr[k]%arr[j] == 0 {
						arrFind[k] = 1
					} else {
						arrFind[k] = 0
					}
				}
			}
		}(i*partSize, (i+1)*partSize)
	}

	wg.Wait()

	elapsedTime := time.Since(t)
	fmt.Println("Прошло времени:", elapsedTime)
}

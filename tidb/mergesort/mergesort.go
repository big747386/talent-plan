package main

import (
	"fmt"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	buffer := make([]int64, len(src))
	Merge_Sort(src, buffer, nil)
}

func Merge_Sort(src []int64, buffer []int64, myChan chan int) {
	n := len(src)
	mid := n >> 1
	if mid > 5 {
		recv := make(chan int)

		go Merge_Sort(src[:mid], buffer[:mid], recv )
		Merge_Sort(src[mid:], buffer[mid:], nil)
		<- recv
		Merge(src[:mid], src[mid:], buffer)
		copy(src, buffer)
	} else {
		for i, j := 1, 0; i < n; i++  {
			temp := src[i]
			for j = i-1; j >= 0 ; j-- {
				if src[j] > temp {
					src[j+1] = src[j]
				} else {
					break
				}
			}
			src[j+1] = temp
		}
	}
	if myChan != nil {
		myChan <- 0
	}
}

func Merge(ll []int64, rr[]int64, buffer []int64) {
	m := len(ll)
	n := len(rr)
	i := 0
	j := 0
	for k := 0; k < m+n; k++ {
		if i >= m {
			buffer[k] = rr[j]
			j++
			continue
		}
		if j >= n {
			buffer[k] = ll[i]
			i++
			continue
		}
		if ll[i] < rr[j]  {
			buffer[k] = ll[i]
			i++
		} else {
				buffer[k] = rr[j]
				j++
		}
	}
}
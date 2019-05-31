package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	Merge_Sort(src, 0, len(src)-1, nil)
}

func Merge_Sort(src []int64, p int, r int, myChan chan int) {
	if p < r {
		q := (p + r) >> 1
		recv := make(chan int)
		go Merge_Sort(src, p, q,recv )
		Merge_Sort(src, q+1, r, nil)
		<- recv
		Merge(src, p, q, r)
	}
	if myChan != nil {
		myChan <- 0
	}
}

const INT_MAX = int64(^uint64(0) >> 1)

func Merge(src []int64, p, q, r int) {
		m := q - p + 1
		n := r - q
		ml := make([]int64, m+1)
		nl := make([]int64, n+1)

		for i := 0; i < m; i++ {
			ml[i] = src[p+i]
		}
		for i := 0; i < n; i++ {
			nl[i] = src[q+i+1]
		}
		ml[m] = INT_MAX
		nl[n] = INT_MAX

		i := 0
		j := 0
		for k := p; k <= r; k++ {
			if ml[i] < nl[j] {
				src[k] = ml[i]
				i++
			} else {
				src[k] = nl[j]
				j++
			}
		}
	}
//func  main() {
//	s := []int64{2,3,4,6,2,6,3,6,3,6,4,67,25632,32,252,235,235}
//	MergeSort(s)
//}
package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	Merge_Sort(src, 0, len(src)-1)
}

func Merge_Sort(src []int64, p int, r int) {
	if p < r {
		q := (p + r)/2
		//fmt.Println(q)
		Merge_Sort(src, p, q)
		Merge_Sort(src, q+1, r)
		Merge(src, p, q, r)
	}
}

const INT_MAX = int64(^uint64(0) >> 1)

func Merge(src []int64, p int, q int, r int) {
	//fmt.Println("##",p,q,r)
	m := q - p + 1
	n := r - q
	ml := make([]int64, m+1)
	nl := make([]int64, n+1)
	//fmt.Println(len(ml),len(nl))
	for i := 0; i < m; i++ {
		ml[i] = src[p+i]
	}
	for i := 0; i < n; i++ {
		nl[i] = src[q+i+1]
	}
	ml[m] = INT_MAX
	nl[n] = INT_MAX
	//fmt.Println("opppp")
	i := 0
	j := 0
	for k := p; k <= r; k++ {
		//fmt.Println("??",i,j)
		if ml[i] < nl[j] {
			src[k] = ml[i]
			i++
		} else {
			src[k] = nl[j]
				j++
		}
	}
}



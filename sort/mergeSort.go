package main

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid := len(arr) / 2
	mergeSort(arr[0:mid])
	mergeSort(arr[mid:len(arr)])
	merge(arr, 0, mid, len(arr))
}

//合并两个有序数组arr[s:m],arr[m:e]
func merge(arr []int, s, m, e int) {
	tmp1, tmp2 := make([]int, m-s), make([]int, e-m)
	copy(tmp1, arr[s:m])
	copy(tmp2, arr[m:e])
	p1, p2 := s, m
	for i := s; i < e; i++ {
		if p1 == m {
			arr[i] = tmp2[p2-m]
			p2++
		} else if p2 == e {
			arr[i] = tmp1[p1-s]
			p1++
		} else if tmp1[p1-s] <= tmp2[p2-m] {
			arr[i] = tmp1[p1-s]
			p1++
		} else {
			arr[i] = tmp2[p2-m]
			p2++
		}
	}
}
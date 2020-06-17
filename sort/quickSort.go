package main

func quickSort(arr []int)  {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, p, r int)  {
	if p >= r {
		return
	}
	q := partition(arr, p, r)
	sort(arr, p, q-1)
	sort(arr, q+1, r)
}

//前后指针法
//func partition(arr []int, p, r int) int {
//	tmp := arr[p]
//	i, j := p+1, r
//	for {
//		for arr[j] > tmp && i <= j  {
//			j--
//		}
//		for arr[i] <= tmp && i < j {
//			i++
//		}
//		if i < j {
//			swap(arr, i, j)
//		} else {
//			break
//		}
//	}
//	swap(arr, p, j)
//	return j
//}

//快慢指针法
func partition(arr []int, p, r int) int {
	tmp := arr[p]
	i, j := p+1, p+1
	for i <= r {
		if arr[i] > tmp {
			j = i+1
			for  j <= r && arr[j] > tmp {
				j++
			}
			if j <= r {
				swap(arr, i, j)
			} else {
				break
			}
		} else {
			i++
		}
	}
	swap(arr, p, i-1)
	return i-1
}
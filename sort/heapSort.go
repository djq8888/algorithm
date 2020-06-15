package main

func heapSort(arr []int) {
	//建堆
	for i := len(arr)/2-1; i >= 0; i-- {
		buildHeap(arr, i, len(arr))
	}

	//将堆顶与数组最后的元素换位，然后调整堆
	for j := len(arr)-1; j >= 0; j-- {
		swap(arr, 0, j)
		buildHeap(arr, 0, j)
	}
}

//构建最大堆
func buildHeap(arr []int, i, j int)  {
	tmp := arr[i]
	for k := 2*i+1; k < j; k = 2*k+1 {
		if k+1 < j && arr[k] < arr[k+1] {
			k = k+1
		}
		if tmp < arr[k] {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = tmp
}

//交换数组里的两个元素
func swap(arr []int, i, j int)  {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
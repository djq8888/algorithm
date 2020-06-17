package main

import "fmt"

func main()  {
	arr := genArray(5,20)
	//arr := []int{17,19,4,0,11}
	fmt.Println("array before sort is:", arr)

	//insertionSort(arr)
	//mergeSort(arr)
	//heapSort(arr)
	quickSort(arr)

	fmt.Println("array after sort is:", arr)
	checkSort(arr)
}
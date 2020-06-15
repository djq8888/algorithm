package main

import "fmt"

func main()  {
	arr := genArray(5,20)
	fmt.Println("array before sort is:", arr)

	//insertionSort(arr)
	//mergeSort(arr)
	heapSort(arr)

	fmt.Println("array after sort is:", arr)
	checkSort(arr)
}
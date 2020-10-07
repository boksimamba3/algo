package main

import "fmt"

// O(logN)
func sort(arr []int) {
	var mergeSort func(arr []int, start int, end int)
	var merge func(arr []int, start int, mid int, end int)

	merge = func(arr []int, start int, mid int, end int) {
		temp := make([]int, end-start+1)
		i := start
		j := mid + 1
		k := 0

		for i <= mid && j <= end {
			if arr[i] < arr[j] {
				temp[k] = arr[i]
				i++
			} else {
				temp[k] = arr[j]
				j++
			}
			k++
		}

		for i <= mid {
			temp[k] = arr[i]
			i++
			k++
		}

		for j <= end {
			temp[k] = arr[j]
			j++
			k++
		}

		for i := start; i <= end; i++ {
			arr[i] = temp[i-start]
		}
	}

	mergeSort = func(arr []int, start int, end int) {
		if start < end {
			mid := (start + end) / 2
			mergeSort(arr, 0, mid)
			mergeSort(arr, mid+1, end)
			merge(arr, start, mid, end)
		}
	}

	mergeSort(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{4, 4, 1, 8, 2, 3, 0, 4}
	sort(arr)
	fmt.Println(arr)
}

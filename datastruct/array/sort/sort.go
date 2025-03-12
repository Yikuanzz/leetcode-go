package array

import (
	"math/rand"
	"slices"
)

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				flag = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i
		for j > 0 && arr[j-1] > tmp {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = tmp
	}
	return arr
}

func shellSort(arr []int) []int {
	size := len(arr)
	gap := size / 2
	for gap > 0 {
		for i := gap; i < size; i++ {
			tmp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > tmp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = tmp
		}
		gap /= 2
	}
	return arr
}

func merge(left_arr, right_arr []int) []int {
	arr := make([]int, 0)
	left_i, right_i := 0, 0
	for left_i < len(left_arr) && right_i < len(right_arr) {
		if left_arr[left_i] < right_arr[right_i] {
			arr = append(arr, left_arr[left_i])
			left_i++
		} else {
			arr = append(arr, right_arr[right_i])
			right_i++
		}
	}

	for left_i < len(left_arr) {
		arr = append(arr, left_arr[left_i])
		left_i++
	}

	for right_i < len(right_arr) {
		arr = append(arr, right_arr[right_i])
		right_i++
	}
	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) >> 1
	left_arr := mergeSort(arr[:mid])
	right_arr := mergeSort(arr[mid:])
	return merge(left_arr, right_arr)
}

func randomPartition(arr []int, low, high int) int {
	i := low + rand.Intn(high-low+1)
	arr[i], arr[low] = arr[low], arr[i]
	return partition(arr, low, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i, j := low, high
	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[low], arr[i] = arr[i], arr[low]
	return i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		pivot := randomPartition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
	return arr
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int) []int {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
	return arr
}

func countingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	arr_min, arr_max := arr[0], arr[0]
	for _, value := range arr {
		if value < arr_min {
			arr_min = value
		} else if value > arr_max {
			arr_max = value
		}
	}

	countRange := arr_max - arr_min + 1
	count := make([]int, countRange)

	for _, value := range arr {
		count[value-arr_min]++
	}

	index := 0
	for i := 0; i < countRange; i++ {
		for count[i] > 0 {
			arr[index] = i + arr_min
			index++
			count[i]--
		}
	}
	return arr
}

func bucketSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	arr_min, arr_max := arr[0], arr[0]
	for _, value := range arr {
		if value < arr_min {
			arr_min = value
		} else if value > arr_max {
			arr_max = value
		}
	}

	bucket_count := len(arr)
	bucket := make([][]int, bucket_count)

	for _, value := range arr {
		bucket_index := (value - arr_min) * bucket_count / (arr_max - arr_min + 1)
		bucket[bucket_index] = append(bucket[bucket_index], value)
	}

	sortedArr := make([]int, 0, len(arr))
	for _, bucket_item := range bucket {
		slices.Sort(bucket_item)
		sortedArr = append(sortedArr, bucket_item...)
	}
	return sortedArr
}

func radixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	arr_max := arr[0]
	for _, value := range arr {
		if value > arr_max {
			arr_max = value
		}
	}

	exp := 1
	for arr_max/exp > 0 {
		countingSortByDigit(arr, exp)
		exp *= 10
	}
	return arr
}

func countingSortByDigit(arr []int, exp int) {
	n := len(arr)

	count := make([]int, 10)
	output := make([]int, n)

	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

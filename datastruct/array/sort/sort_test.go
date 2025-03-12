package array

import (
	"reflect"
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		actual := bubbleSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}

func Test_SelectionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		actual := selectionSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}

func Test_InsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		actual := insertionSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}

func Test_ShellSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		actual := shellSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}

func Test_MergeSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		actual := mergeSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}

func Test_QuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		low, high := 0, len(test.input)-1
		actual := quickSort(test.input, low, high)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}

func Test_HeapSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		actual := heapSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}

func Test_CountingSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		actual := countingSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}

func Test_BucketSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		actual := bucketSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}

func Test_RadixSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, test := range tests {
		actual := radixSort(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v, but got %v", test.expected, actual)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		bubbleSort(input)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		selectionSort(input)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		insertionSort(input)
	}
}

func BenchmarkShellSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		shellSort(input)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		mergeSort(input)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		quickSort(input, 0, len(input)-1)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		heapSort(input)
	}
}

func BenchmarkCountingSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		countingSort(input)
	}
}

func BenchmarkBucketSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		bucketSort(input)
	}
}

func BenchmarkRadixSort(b *testing.B) {
	input := []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		radixSort(input)
	}
}

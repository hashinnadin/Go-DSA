package main

func BubbbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func SelectionSort(arr []int) {

	n := len(arr)

	for i := 0; i < n-1; i++ {
		min := i

		for j := i + 1; i < n; j++ {
			if arr[j] > arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]

	}
}

func Insert(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := i
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {

}

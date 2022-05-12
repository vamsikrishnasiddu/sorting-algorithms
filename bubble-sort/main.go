package main

import "fmt"

func bubleSort(arr []int) {

	n := len(arr)

	for k := 1; k <= n-1; k++ {
		flag := 0
		for i := 0; i <= n-k-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = 1
			}

		}

		if flag == 0 {
			break
		}

	}

}

func main() {
	var arr []int = []int{2, 7, 4, 1, 5, 3}

	bubleSort(arr)

	fmt.Println(arr)
}

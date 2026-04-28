package internal

import "fmt"

// Service 函数用于输出问候信息
func Service() {
	// 打印 "Hello, world!" 到标准输出
	fmt.Println("domain Service")
}

func BubbleSort(arr []int) {
	// 实现冒泡排序算法
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				// 交换 arr[j] 和 arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

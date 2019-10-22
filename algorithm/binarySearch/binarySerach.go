package main

import (
	"flag"
	"fmt"
	"os"
)

//os包 用来接收命令行参数 os.Args
//flag包 用来解析命令行参数  flag.Int() flag.Parse

func BinarySearch(array []int, num int) int {
	left := 0
	right := len(array) - 1
	mid := 0
	for left <= right {
		mid = left + (right-left+1)/2
		if array[mid] == num {
			return mid
		} else if array[mid] < num {
			left = mid + 1
		} else if array[mid] > num {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println("start")

	fmt.Println(os.Args)

	num := flag.Int("num", 0, "num")
	flag.Parse()
	fmt.Println("num", *num)

	array := []int{1, 3, 5, 7, 9}

	rst := BinarySearch(array, *num)
	fmt.Println(rst)

	fmt.Println("end")
}

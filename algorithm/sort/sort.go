package _sort

import (
	"fmt"
)

//冒泡排序、插入排序、选择排序、归并排序、快速排序、计数排序、基数排序、桶排序

//时间复杂度： 最好 最坏  平均
//空间复杂度： 原地排序（特指：空间复杂度为O(1)）
//稳定性：

//Sort 最简单的方法  o(n^2)
func Sort(array []int, num int) {
	for i := 0; i < num; i++ {
		for j := i; j < num; j++ {
			if array[i] > array[j] {
				tmp := array[i]
				array[i] = array[j]
				array[j] = tmp
				//fmt.Println(i, j, array)
			}
			//fmt.Println(i, j, array)
		}
		fmt.Println(array)
	}
}

//BubbleSort 冒泡排序 array数组  num数组的大小
//空间复杂度O(1)   最优时间复杂度O(n)   最差时间复杂度O(n^2)
func BubbleSort(array []int, num int) {
	if array == nil || num < 0 {
		return
	}
	// fmt.Println("11", unsafe.Pointer(&array))
	for i := 0; i < num; i++ {
		change := false
		for j := 0; j < num-i-1; j++ {
			if array[j] > array[j+1] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
				change = true
			}
		}
		if !change {
			break
		}
		fmt.Println(array)
	}
	// fmt.Println("22", unsafe.Pointer(&array))
	return
}

//InsertionSort 插入排序
//空间复杂度O(1)   最优时间复杂度O(n)   最差时间复杂度O(n^2)
func InsertionSort(array []int, num int) {
	if array == nil || num < 0 {
		return
	}

	for i := 1; i < num; i++ {
		value := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if array[j] > value {
				array[i] = array[j]
			} else {
				break
			}
		}
		array[j+1] = value
		fmt.Println(array)
	}
}

//SelectionSort 选择排序
//空间复杂度O(1)   最优时间复杂度O(n^2)   最差时间复杂度O(n^2)
func SelectionSort(array []int, num int) {
	if array == nil || num < 0 {
		return
	}

	for i := 0; i < num; i++ {

		j := i
		min := j
		for ; j < num; j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		tmp := array[i]
		array[i] = array[min]
		array[min] = tmp
	}
}

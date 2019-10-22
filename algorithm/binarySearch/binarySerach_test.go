package main

import "fmt"

//测试类型 1.test 2.benchmark 3.example
//1.test Test
//2.
//3.

// func TestBinarySearch(t *testing.T) {
// 	array := []int{1, 3, 5, 7, 9}

// 	if rst := BinarySearch(array, 1); rst == -1 {
// 		fmt.Println("return -1, num = 1")
// 	} else {
// 		fmt.Println("return , num = 1", rst)
// 	}

// 	if rst := BinarySearch(array, 2); rst == -1 {
// 		fmt.Println("return -1, num = 2")
// 	} else {
// 		fmt.Println("return , num = 3", rst)
// 	}

// 	if rst := BinarySearch(array, 3); rst == -1 {
// 		fmt.Println("return -1, num = 3")
// 	} else {
// 		fmt.Println("return , num = 3", rst)
// 	}

// 	if rst := BinarySearch(array, 4); rst == -1 {
// 		fmt.Println("return -1, num = 4")
// 	} else {
// 		fmt.Println("return , num = 4", rst)
// 	}

// 	if rst := BinarySearch(array, 5); rst == -1 {
// 		fmt.Println("return -1, num = 5")
// 	} else {
// 		fmt.Println("return , num = 5", rst)
// 	}

// 	if rst := BinarySearch(array, 6); rst == -1 {
// 		fmt.Println("return -1, num = 6")
// 	} else {
// 		fmt.Println("return , num = 6", rst)
// 	}

// 	if rst := BinarySearch(array, 7); rst == -1 {
// 		fmt.Println("return -1, num = 7")
// 	} else {
// 		fmt.Println("return , num = 7", rst)
// 	}

// }

// func BenchmarkBinarySearch(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		array := []int{1, 3, 5, 7, 9}

// 		if rst := BinarySearch(array, 1); rst == -1 {
// 			//	fmt.Println("return -1, num = 1")
// 		} else {
// 			//	fmt.Println("return , num = 1", rst)
// 		}
// 	}
// }

func ExampleBinarySearch() {
	array := []int{1, 3, 5, 7, 9}

	fmt.Println(BinarySearch(array, 2))
	fmt.Println(BinarySearch(array, 1))
	fmt.Println(BinarySearch(array, 3))
	// Output:
	// -1
	// 0
	// 1
}

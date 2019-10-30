package a

//任意给定一个32位无符号整数n，求n的二进制表示中1的个数，比如n = 5（0101）时，返回2，n = 15（1111）时，返回4

import "fmt"

func BItCount(num uint32) uint32 {
	var count uint32 = 0
	for num > 0 {
		if (num & 1) == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

func BItCOunt1(num uint32) uint32 {
	var count uint32
	for count = 0; num > 0; num >>= 1 {
		count += num & 1
	}
	return count
}

func BItCount2(num uint32) uint32 {
	var count uint32
	for count = 0; num > 0; count++ {
		num &= num - 1
	}
	return count
}

func BitCount3(num uint32) uint32 {
	// 建表
	var BitsSetTable256 = [256]int{}

	// 初始化表
	for i := 0; i < 256; i++ {
		BitsSetTable256[i] = (i & 1) + BitsSetTable256[i/2]
	}

	fmt.Println(BitsSetTable256)
	//var c uint32 = 0

	// 查表
	// unsigned char* p = (unsigned char*) &n ;

	// c = BitsSetTable256[p[0]] +
	//     BitsSetTable256[p[1]] +
	//     BitsSetTable256[p[2]] +
	//     BitsSetTable256[p[3]];

	return 0
}

func BitCount4(n uint32) uint32 {

	fmt.Println(n)
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
	fmt.Println(n)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	fmt.Println(n)
	n = (n & 0x0f0f0f0f) + ((n >> 4) & 0x0f0f0f0f)
	fmt.Println(n)
	n = (n & 0x00ff00ff) + ((n >> 8) & 0x00ff00ff)
	fmt.Println(n)
	n = (n & 0x0000ffff) + ((n >> 16) & 0x0000ffff)
	fmt.Println(n)

	return n
}

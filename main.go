package main

import (
	"fmt"
)

const str1 = "Hello World1"
const str2 = `Hello World2`
const bt byte = 's'

const a, b, c int = 0, 1, 2
const d, e = 0, 1

func learnBasicType() {
	//var blank_string = "hello"
	//var va, vb, vc int = 1, 2, 3
	//vd, ve, vString := 1, 2, "hello"
	fmt.Print("不换行")
	fmt.Println("----")
	fmt.Printf(str1 + str2)
	fmt.Printf("%s: age is %d", "Tom", "19")
	var s = fmt.Sprintf("%d,%d,%d,%d", a, b, c, d)
	fmt.Println(s)
}

func learnArray() {
	var arr1 [5]int
	fmt.Println(arr1)

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	arr4 := [...]int{1, 2, 3, 4, 5, 5}
	fmt.Println(arr4)

	arr5 := [5]int{0: 3, 2: 4, 4: 6}
	fmt.Println(arr5)

	// 二维数组
	arr6 := [...][5]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr6)
	fmt.Println(arr6[1][4])

	var arr7 = [3][5]int{{1, 2, 3, 4, 5}, {6, 7}, {8, 9}}
	fmt.Println(arr7)

	arr8 := [3][5]int{{1, 2, 3, 4, 5}, {6, 7}, {8, 9}}
	fmt.Println(arr8)

	modifyArray(&arr8[0])
	fmt.Println(arr8)

	arr7 = arr8 // 相同长度和类型的才可以直接赋值
	fmt.Println(arr7)

	var arr9 [3][5]uint8
	//arr9 = arr8
	fmt.Println(arr9)
}

func modifyArray(a *[5]int) {
	a[0] = 100
}

func main() {
	fmt.Println("Hello World!")
	learnBasicType()
	learnArray()
}

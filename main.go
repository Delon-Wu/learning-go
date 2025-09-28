package main

import (
	"fmt"
	"sync"
)

const str1 = "Hello World1"
const str2 = `Hello World2`

//const bt byte = 's'
//
//const a, b, c int = 0, 1, 2
//const d, e = 0, 1

func learnBasicType() {
	//var a int = 0xf
	//var b int = 0xF
	//
	//var aa uint8 = uint8(a)
	//
	//var c uint8 = 0b11001100
	//
	//var d uint8 = 0o17 // 八进制
	//
	//var f float64 = 10

	//f2 := 10.0

	//var c1 complex64 = a + bi
	//var c1 complex64
	//c1 = 1.10 + 0.1i
	//c2 := 1.10 + 0.1i

	var bytes []byte = []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	var s string = string(bytes)

	//var blank_string = "hello"
	//var va, vb, vc int = 1, 2, 3
	//vd, ve, vString := 1, 2, "hello"
	fmt.Print("不换行")
	fmt.Println("----")
	fmt.Printf(str1 + str2)
	fmt.Printf("%s: age is %d", "Tom", 19)
	//var s = fmt.Sprintf("%d,%d,%d,%d", a, b, c, d)
	fmt.Println(s)

	var r1 rune = 'a'
	var r2 rune = '爱'

	fmt.Println(r1, r2)

	var ss string = "abc, 你好，世界"
	var runes []rune = []rune(ss)
	fmt.Println(runes, len(runes))

	var s1, s2 string = "Foo", `Foo`
	fmt.Println("s1 == s2: %d", s1 == s2)

	var (
		group1 = 1
		group2 = "Hello"
		group3 = []int{1, 2, 3}
	)
	var m map[string]string
	var b byte = 2
	var slice []byte
	var p *int
	fmt.Println(group1)
	fmt.Println(group2)
	fmt.Println(group3)
	//m["Hello"] = "World"
	//fmt.Println(m)
	m = make(map[string]string)
	m["Hello"] = "World"
	fmt.Println(m)
	fmt.Println(b)
	fmt.Println(slice)
	fmt.Println(p)

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

func method1() {
	var a = 1
	var b int
	var c string = "我是局部变量"
	fmt.Println(a, b, c)
}

func method2() (a int, b int, c string) {
	return a, 3, "Method 2"
}
func method3() (a int, b *map[string]int, c string) {
	a = 1
	m := make(map[string]int)
	b = &m
	c = "Method 3"
	//return a, b, c // 自动返回
	return
}

func method4() (s string, i int) {
	return
}

func learnVariable() {
	method1()
	fmt.Println(method2())
	var a, m, c = method3()
	(*m)["a"] = 1
	fmt.Println(a, *m, c)

	fmt.Println("method4() return: ")
	fmt.Println(method4())
}

func modifyArray(a *[5]int) {
	a[0] = 100
}

func learnPointer() {
	//var p *<type>
	var p *[5]int
	arr := [5]int{1, 2, 3, 4, 5}
	p = &arr
	// p++ // *go中指针不能偏移 运算
	var temp int = arr[0]
	(*p)[0] = arr[4]
	(*p)[4] = temp
	fmt.Println(*p)

	var p1 **string
	//var strP *string = &"Hello" 不能这样，要通过变量获取地址
	var strP *string
	var s = "Hello Pointer"
	strP = &s
	p1 = &strP
	fmt.Println(p1)
	fmt.Println(strP == (*p1))
	fmt.Println(**p1)

	// 使用 unsafe 的时候要特别小心
	//up1 := unsafe.Pointer(p)
	//uinp1 := uintptr(up1)
	//uinp1 += 1
	//p2 := (*int)(unsafe.Pointer(uinp1))
	//fmt.Println(*p2)

}

func learnStruct() {
	// go当中没有类的概念
	// 首字母大写是公开的，小写是非公开的
	type Persion struct {
		Weight, Age, Height float64
		Name                string      `json:"name"` // 这是标签
		Eat                 func() byte `json:"eat"`
		Sleep               func()
		Map                 map[string]string
		Slice               []interface{}
		Arr                 [3]int
		Ch                  chan string
		once                sync.Once
		Ptr                 *int
	}
	// 匿名结构体
	var Tom = struct {
		Weight, Age, Height float64
		Name                string `json:"name"` // 这是标签
		Eat                 func() `json:"eat"`
		Sleep               func()
		Map                 map[string]string
		Slice               []interface{}
		Arr                 [3]int
		Ch                  chan string
		once                sync.Once
		Ptr                 *int
		BestFriend          Persion
	}{
		Weight: 70.1,
		Age:    20,
		Height: 190,
		Name:   "Tom",
		Eat: func() {
			fmt.Println("Yummy~~")
		},
		Sleep: func() {
			fmt.Println("Good Night!")
		},
	}
	fmt.Println(Tom.Name)

	type StructA struct {
		Name string
		Age  int
	}
	type StructB struct {
		Field1 string
		StructA
	}
	var b StructB = StructB{
		Field1: "Hello",
		StructA: StructA{
			Name: "StructA",
			Age:  20,
		},
	}
	fmt.Println(b)
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println("----------------------learnBasicType---------------------")
	learnBasicType()
	fmt.Println("----------------------learnArray---------------------")
	learnArray()
	fmt.Println("----------------------learnVariable---------------------")
	learnVariable()
	fmt.Println("----------------------learnPointer---------------------")
	learnPointer()
	fmt.Println("---------------------learnStruct----------------------")
	learnStruct()
}

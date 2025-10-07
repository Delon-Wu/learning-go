package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
	"time"
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

	//var s1, s2 string = "Foo", `Foo`
	//fmt.Println("s1 == s2: %d", s1 == s2)

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

func learnSlice() {
	var blankSlice []int
	//blankSlice[0] = 1 // 错误用法
	fmt.Println(blankSlice, cap(blankSlice))
	var blankSlice2 []int = []int{}
	fmt.Println(blankSlice2)

	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	var slice2 []int = make([]int, 3, 5)
	slice3 := make([]int, 3, 5)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	var slice4 = slice1[1:]
	var slice5 = slice1[1:3:7]
	fmt.Println(slice4, cap(slice4))
	fmt.Println(slice5, cap(slice5))
	fmt.Println(slice5 == nil) // 只能和nil比较

	var slice6 = []int{1, 2, 3}
	slice6 = append(slice6, 4)
	fmt.Println(slice6, cap(slice6))
	slice6 = append(slice6, 5, 6, 7, 8)
	fmt.Println(slice6, cap(slice6))

	var slice7 = []int{1, 2, 3, 4, 5, 6, 7}
	slice7 = slice7[:len(slice7)-1]
	fmt.Println(slice7)
	slice7 = slice7[1:]
	fmt.Println(slice7)
	slice7 = append(slice7[:2], slice7[2+2:]...)
	fmt.Println(slice7)
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

	type A struct {
		a string
	}
	type B struct {
		b string
	}
	type C struct {
		A
		B
		a string
		b string
	}
	type D struct {
		da A
		b  B
		c  C
	}

	a1 := A{a: "I'm A"}
	b1 := B{b: "I'm B"}
	c1 := C{A: a1, B: b1, a: "I'm cA", b: "I'm cB"}
	d1 := D{da: a1, b: b1, c: c1}
	fmt.Println(d1.da)

	// * 想通过结构体方法改变结构体的属性，那就要通过传入指针来修改 *

	copyA1 := a1 // 深拷贝
	copyA1.a = "I'm copyA1"
	fmt.Println(a1, copyA1)
	pA1 := &a1
	(*pA1).a = "I'm changed A1"
	fmt.Println(a1, *pA1)
}

func learnMap() {
	var m1 map[string]int
	m1 = make(map[string]int)
	fmt.Println(m1)
	m1["foo"] = 256
	fmt.Println(m1)

	var m2 map[string]string = map[string]string{"foo": "m1_bar"}
	m3 := map[string]string{"foo": "m2_bar"}
	fmt.Println(m2)
	fmt.Println(m3)

	res := make(map[string]interface{})
	res["code"] = 200
	res["data"] = map[string]interface{}{"foo": "res", "value": 1000000000000}
	res["msg"] = "success"
	fmt.Println(res)

	jsonStr, errs := json.Marshal(res)
	if errs == nil {
		fmt.Printf("序列化之后：%s\n", jsonStr)
	}

	res2 := map[string]interface{}{}
	errs = json.Unmarshal(jsonStr, &res2)
	if errs != nil {
		return
	}
	fmt.Println(res2)

	res["code"] = 400
	delete(res2, "code")
	fmt.Println(res, res2)
}

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func (gender *Gender) IsMale() bool {
	return *gender == "male"
}
func (gender *Gender) IsFemale() bool {
	return *gender == "female"
}

type WorkDay int

const (
	Monday  WorkDay = -3
	Tuesday WorkDay = iota // 自动推断下面的值
	Wednesday
	Thursday
	Friday
)
const isolateGroup int = iota

func learnEnum() {
	// 类型别名
	type AliasInt int
	const constA AliasInt = 1
	const constB = AliasInt(2)
	fmt.Println(constA, constB)

	// 分组声明
	const (
		groupA = 1
		groupB = 2
		groupC = "3"
	)
	fmt.Println(groupA, groupB, groupC)

	gender := Gender("male")
	fmt.Println(gender)
	fmt.Printf("gender.IsMale(): %t, gender.IsFemale(): %t", gender.IsMale(), gender.IsFemale())

	const thirdDay = Wednesday
	fmt.Println(thirdDay)
}

func learnLoop() {
	person := [3]string{"Tom", "Jack", "Jill"}
	fmt.Println(person)

	for k, v := range person {
		fmt.Printf("persion[%d]: %s\n", k, v)
	}

	for i := range person {
		fmt.Printf("persion[%d]: %s\n", i, person[i])
	}

	for i := 0; i < len(person); i++ {
		fmt.Printf("persion[%d]: %s\n", i, person[i])
	}

	for _, name := range person {
		fmt.Println(name)
	}

	vegetables := []string{"potato", "tomato", "onion"}
	fmt.Println(vegetables)
	for k, v := range vegetables {
		fmt.Printf("vegetables[%d]: %s\n", k, v)
	}

	m := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(m, len(m))
	for k, v := range m {
		fmt.Printf("m[%d]: %s\n", k, v)
	}
	m[4] = "d"
	fmt.Println(m)

	fmt.Println("遇到c就跳出循环：")
	for _, v := range m {
		if v == "c" {
			break
		}
		fmt.Println(v)
	}

	fmt.Println("不要 a：")
	for _, v := range m {
		if v == "a" {
			continue
		}
		fmt.Println(v)
	}

	// goto
	for _, v := range person {
		if v == "Jill" {
			goto STAGE
		}
	}

	fmt.Println("-----------嘿！看看我被跳过了吗------------")

STAGE:
	fmt.Println("Jill, 你被跳过了")

	type Month int
	const (
		January Month = iota
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)
	month := October + 10
	fmt.Printf("month: %d\n", month+1)
	// 默认每个case都有break
	switch month {
	case March, April, May:
		fmt.Println("这是春天的月份")
	case June, July, August:
		fmt.Println("这是夏天的月份")
	case September, October, November:
		fmt.Println("这是秋天的月份")
	case January, February, December:
		fmt.Println("这是冬天的月份")
	default:
		fmt.Println("没有这个月份")
	}

	num := byte(1)
	var d interface{}
	d = &num
	switch t := d.(type) {
	case *int:
		fmt.Printf("It's %T\n", t)
	default:
		fmt.Printf("It's %T\n", t)
	}
}

type StructA struct {
	Name string
}

func changeName(s *StructA, newName string) {
	(*s).Name = newName
}

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func getTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func getTimeInt() int64 {
	return time.Now().Unix()
}

func creatSign(params map[string]interface{}) string {
	var keys []string
	var str = ""
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i := 0; i < len(keys); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", keys[i], params[keys[i]])
		} else {
			str += fmt.Sprintf("&xl_%v=%v", keys[i], params[keys[i]])
		}
	}
	var secret = "123123"
	sign := MD5(MD5(str) + MD5(secret))
	return sign
}

func learnFun() {
	//传递参数时，go默认复制一份参数出来作为入参
	var structA = StructA{
		Name: "a",
	}
	changeName(&structA, "Tom")
	fmt.Println(structA)

	fmt.Printf("MD5(\"Hello world!\"): %s\n", MD5("Hello world!"))

	fmt.Printf("Current time is: %s\n", getTimeStr())

	fmt.Println(getTimeInt())

	m := map[string]interface{}{
		"userName": "Tom2025",
		"age":      30,
		"Password": "123456",
	}
	sign := creatSign(m)
	fmt.Println("creatSign:\n", m, sign)
}

func producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}

func readChan(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func learnChan() {
	fmt.Println("START")

	//go func() {
	//	fmt.Println("Hello Channel")
	//}()
	ch := make(chan string, 3)
	go producer(ch)
	go readChan(ch) // 不及时读取就会堵塞

	//time.Sleep(1 * time.Second)
	fmt.Println("END")
}

func learnOperator() {
	var a, b = 10, 5
	v1 := a + b
	v2 := a - b
	v3 := a * b
	v4 := a / b
	v5 := a % b
	fmt.Println(v1, v2, v3, v4, v5)

	a++
	fmt.Println(a)

	//++a // 没有前置的自增用法
	//fmt.Println(a)

	//c := a++ + b // 不能在运算中使用
	//c := (a++) + b // 不能在运算中使用
	//fmt.Println(c)

	//fmt.Println(a++) // 不能在运算中使用

	a += 1
	b -= 1
	fmt.Println(a, b)

	//a = 10 + 0.1 // 只能是相同数据类型进行运算
	//fmt.Println(a)
	f := float64(10) + 0.1
	fmt.Println(f)
	b1 := byte(1) + 1
	fmt.Println(b1)
	num := int(f)
	fmt.Println(num) // 小数位会被强制截除掉

	//运算符
	// > < == >= <= !=
	// 逻辑运算符：&& || !
	// 位运算符： & | ^
	fmt.Println(a == (-1))

	num = 1
	num <<= 1
	fmt.Println(num)
	num += 6 * 2
	fmt.Println(num)

	if c := 100; a < c {
		fmt.Println("a < c", a < c, a, c)
	}
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
	fmt.Println("----------------------learnSlice---------------------")
	learnSlice()
	fmt.Println("---------------------learnStruct----------------------")
	learnStruct()
	fmt.Println("---------------------learnMap----------------------")
	learnMap()
	fmt.Println("---------------------learnEnum----------------------")
	learnEnum()
	fmt.Println("---------------------learnLoop----------------------")
	learnLoop()
	fmt.Println("---------------------learnFun----------------------")
	learnFun()
	fmt.Println("---------------------learnChan----------------------")
	learnChan()
	fmt.Println("---------------------learnOperator----------------------")
	learnOperator()
}

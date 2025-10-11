package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
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

	var arr10 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr10)

	// 三维数组，遍历多维数组时，有x维就嵌套x层遍历
	arr11 := [3][2][2]int{
		{{1, 2}, {3, 4}},
		{{5, 2}, {6, 4}},
	}
	fmt.Println(arr11, "arr11[0][1][1]: ", arr11[0][1][1])

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

	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	var slice2 []int = make([]int, 3, 5)
	slice3 := make([]int, 3, 5)
	fmt.Println(arr)
	fmt.Println(slice2)
	fmt.Println(slice3)

	var slice4 = arr[1:]
	var slice5 = arr[1:3:7]
	fmt.Println(slice4, cap(slice4))
	fmt.Println(slice5, cap(slice5))
	fmt.Println(slice5 == nil) // 只能和nil比较

	fmt.Println("arr: ", arr)
	arr[4] = 99
	fmt.Println("slice4: ", slice4) // 同样被改为99，
	// 所以，切片的元素是数组元素的一个引用，切片实际上是特殊的数组

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

	//复制
	src := []int{1, 2, 3, 4, 5, 6, 7}
	dst := make([]int, 3, 6)
	fmt.Println("before copy, src, dst = ", src, dst)
	copy(dst, src)
	fmt.Println("after copy, src, dst = ", src, dst)
	src1 := []int{1, 2}
	dst1 := make([]int, 3, 6)
	fmt.Println("before copy, src1, dst1 = ", src1, dst1)
	copy(dst1, src1)
	fmt.Println("after copy, src1, dst1 = ", src1, dst1)

	slice8 := make([]int, 3, 6)
	slice9 := append(slice8, 4, 2)
	fmt.Println("---------\n", slice8, len(slice8), cap(slice8))
	fmt.Println("---------\n", slice9, len(slice9), cap(slice9))
	slice8[1] = 1
	slice9[1] = 2
	fmt.Println(slice8)
	fmt.Println(slice9)
	// slice9在函数调用中作为参数使用也一样会修改引用指向的值
	slice10 := append(slice9, 6, 6, 6)
	fmt.Println(slice10)
	slice10[1] = -1 // 扩容后会将元素复制到一个新的切片当中，就不会再影响原来的切片
	fmt.Println(slice7, slice8, slice9, slice10)
}

func learnStruct() {
	// go当中没有类的概念
	// 首字母大写是公开的，小写是非公开的
	type Person struct {
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
		BestFriend          Person
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

	noneExitValue, exists := m1["foo1"]
	fmt.Println(exists, noneExitValue)

	var m2 = map[string]string{"foo": "m1_bar"}
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
	res2["newFoo"] = "new foo"
	fmt.Println(res, res2, len(res), len(res2))
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

	i := 10
	for {
		if i == 0 {
			break
		}
		i--
		fmt.Println(i)
	}

	for k, v := range person {
		fmt.Printf("Person[%d]: %s\n", k, v)
	}

	for i := range person {
		fmt.Printf("Person[%d]: %s\n", i, person[i])
	}

	for i := 0; i < len(person); i++ {
		fmt.Printf("Person[%d]: %s\n", i, person[i])
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
			goto STAGE // 不建议使用goto，代码会变得难以维护
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

	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
		for j := 5; j > 0; j-- {
			fmt.Printf("j = %d\n", j)
			if j == 3 {
				break // 只退出最近的for循环
			}
		}
	}
	fmt.Println("---------------")
outer:
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
		for j := 5; j > 0; j-- {
			fmt.Printf("j = %d\n", j)
			if j == 3 {
				break outer // 退出到指定位置
			}
		}
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

type StructB struct {
	b string
}

func (b *StructB) change(str string) {
	b.b = str
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

	//闭包
	sb := StructB{
		b: "default",
	}
	functionChange := sb.change
	returnFunc := func() func(int, string) (int, string) {
		fmt.Println("我是一个匿名函数")
		return func(n int, s string) (int, string) {
			return n, s
		}
	}()
	fmt.Println(sb)
	functionChange("Changed sb")
	fmt.Printf("执行functiohnChange之后的 sb.b: %s\n", sb.b)
	res1, res2 := returnFunc(10, "Hello world")
	fmt.Println(res1, res2)
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

func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("发送前：", i)
		// 如果缓冲区满了，会阻塞在这一步
		ch <- i
		fmt.Println("发送后：", i)
	}
	close(ch) // 关闭之后不回再阻塞
}

func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Println("接收到：", v)
	}
}

func learnChan() {
	fmt.Println("START")

	//go func() {
	//	fmt.Println("Hello Channel")
	//}()

	//ch := make(chan string, 3)
	//go producer(ch)
	//go readChan(ch) // 不及时读取就会堵塞

	ch := make(chan int, 3)
	go sendOnly(ch)
	//go receiveOnly(ch)

	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok { // 通道是什么类型，关闭后ok就是什么类型的零值
				fmt.Println("Channel已经关闭")
				return
			}
			fmt.Printf("主goroutine接收到：%d\n", v)
			time.Sleep(500 * time.Millisecond)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(100 * time.Millisecond)
		}
	}
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

func addDataInChan(ch chan int) {
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}
	close(ch) // 如果不关闭通道，程序进入asleep状态， 相当于死锁
}

func learnRange() {
	str1 := "abc一二三"
	for index := range str1 {
		fmt.Println(index, str1[index], string(str1[index]))
	}
	fmt.Println("length: ", len(str1))

	arr := [...]int{0, 1, 2, 3}
	slice := []int{0, 1}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	slice = append(slice, 2, 3, 4)
	fmt.Println("After append:", len(slice), cap(slice))
	for i, v := range slice {
		fmt.Println(i, v)
	}

	hashMap := map[string]int{
		"abandon": 1,
		"banana":  2,
		"cherry":  3,
	}
	for k, v := range hashMap {
		fmt.Println(k, v)
	}

	//遍历通道
	//ch := make(chan int, 10)
	//go addDataInChan(ch)
	//for v := range ch {
	//	fmt.Println(v)
	//}
}

type Supplier interface {
	Get() string
}

type DigitSupplier struct {
	value int
}

func (d *DigitSupplier) Get() string {
	return fmt.Sprintf("%d", d.value)
}

func learnTypeChanging() {
	var i int32 = 17
	var b byte = 5
	var f float32
	f = float32(i) / float32(b)
	fmt.Println(f)

	var i2 int32 = 256
	fmt.Println("i2:", i2)
	var b2 = byte(i2)
	fmt.Println(b2)

	str := "abc, 123, 世界，你好"
	bytes := []byte(str)
	runes := []rune(str)
	fmt.Println(bytes)
	fmt.Println(runes)

	s1 := string(bytes)
	s2 := string(runes)
	fmt.Println(s1, s2)

	// 字符串转数字
	numStr := "123" // "123a"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)

	numStr1 := strconv.Itoa(666)
	fmt.Println(numStr1)

	uin32, err := strconv.ParseUint(numStr, 10, 32)
	if err != nil {
		panic(err)
	}
	fmt.Println(uin32)
	numStr = "123a"
	i32, err := strconv.ParseInt(numStr, 16, 32)
	if err != nil {
		panic(err)
	}
	fmt.Println(i32)

	//数字转字符串
	numStr2 := strconv.FormatUint(0o127, 8)
	numStr3 := strconv.FormatInt(-1212, 2)
	fmt.Println(numStr2, numStr3)
	numStr4 := fmt.Sprintf("%d", int32(999))
	fmt.Println(numStr4)

	// 接口转换
	var interface1 interface{} = int32(3)
	a, ok := interface1.(int32) // int
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("转换失败")
	}

	switch t := interface1.(type) {
	case int:
		fmt.Println("The type is int", t)
	case string:
		fmt.Println("The type is string", t)
	case int32:
		fmt.Println("The type is int32", t)
	}

	// 结构体和接口互转
	var supplierA Supplier = &DigitSupplier{value: 9}
	fmt.Println(supplierA.Get())

	b111, ok := (supplierA).(*DigitSupplier)
	if ok {
		fmt.Println(b111.Get())
	}
}

type PayMethod interface {
	Account              // 接口允许嵌套
	Pay(amount int) bool // 接口方法不一定要公开（首字母大写），参数名也不一定要有，要有参数的数据类型
}

type Account interface {
	GetBalance() int
}

type emptyInterface interface{}

type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) GetBalance() int {
	return c.balance
}

func (c *CreditCard) Pay(amount int) bool {
	if amount > c.limit {
		fmt.Println("信用卡支付失败，超出额度")
		return false
	}

	c.balance += amount
	fmt.Println("信用卡支付成功：", amount)
	return true
}

type DebitCard struct {
	balance int
}

func (d *DebitCard) GetBalance() int {
	return d.balance
}

func (d *DebitCard) Pay(amount int) bool {
	if amount > d.balance {
		fmt.Println("借记卡余额不足，支付失败：", d.balance)
		return false
	}
	d.balance -= amount
	return true
}

func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Println("购买成功，剩余余额：", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

func learnInterface() {
	creditCard := &CreditCard{balance: 0, limit: 1000}
	debitCard := &DebitCard{balance: 1000}

	fmt.Println("使用信用卡购买")
	purchaseItem(creditCard, 1000)

	fmt.Println("使用借记卡购买")
	purchaseItem(debitCard, 800)

	fmt.Println("使用借记卡购买")
	purchaseItem(debitCard, 800)

	var account Account = creditCard
	fmt.Println("获取账户余额：", account.GetBalance())

	var a emptyInterface
	var b = 1
	a = debitCard
	a = 1
	a = 1.4
	a = true
	a = &b
	a = "Hello world"
	fmt.Println(a)
}

func repeatSaySomething(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}

type SaveCounter struct {
	mu    sync.Mutex
	count int
}
type UnsaveCounter struct {
	count int
}

func (s *SaveCounter) increment() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count++
}

func (s *SaveCounter) GetCount() int {
	return s.count
}

func (c *UnsaveCounter) increment() {
	c.count++
}
func (c *UnsaveCounter) GetCount() int {
	return c.count
}

func learnConcurrency() {
	go func() {
		fmt.Println("goroutine in closure")
	}()

	go func(s string) {
		fmt.Println(s)
	}("goroutine paragram")

	go repeatSaySomething("hello goroutine")

	//counter := SaveCounter{}
	unsaveCounter := UnsaveCounter{}
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				//counter.increment()
				unsaveCounter.increment()
			}
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("Final count: ", unsaveCounter.GetCount())
	fmt.Println("learn goroutine")
}

func learnSelect() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
	}()

	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	for i := 0; i < 5; i++ {
		select { // 可见，选择哪个分支是随机的
		case value := <-ch1:
			fmt.Printf("receive %d from ch1\n", value)
		case value := <-ch2:
			fmt.Printf("receive %d from ch2\n", value)
		case value := <-ch3:
			fmt.Printf("receive %d from ch3\n", value)
		case <-ctx.Done():
			fmt.Println("Context done")
			return
		}
		fmt.Println("select end")

	}
}

func main() {
	fmt.Println("Hello World!")
	//fmt.Println("----------------------learnBasicType---------------------")
	//learnBasicType()
	//fmt.Println("----------------------learnArray---------------------")
	//learnArray()
	//fmt.Println("----------------------learnVariable---------------------")
	//learnVariable()
	//fmt.Println("----------------------learnPointer---------------------")
	//learnPointer()
	fmt.Println("----------------------learnSlice---------------------")
	learnSlice()
	//fmt.Println("---------------------learnStruct----------------------")
	//learnStruct()
	fmt.Println("---------------------learnMap----------------------")
	learnMap()
	//fmt.Println("---------------------learnEnum----------------------")
	//learnEnum()
	//fmt.Println("---------------------learnLoop----------------------")
	//learnLoop()
	//fmt.Println("---------------------learnFun----------------------")
	//learnFun()
	fmt.Println("---------------------learnChan----------------------")
	learnChan()
	//fmt.Println("---------------------learnOperator----------------------")
	//learnOperator()
	fmt.Println("---------------------learnRange----------------------")
	learnRange()
	fmt.Println("---------------------learnTypeChanging----------------------")
	learnTypeChanging()
	fmt.Println("---------------------learnInterface----------------------")
	learnInterface()
	fmt.Println("---------------------learnConcurrency----------------------")
	learnConcurrency()
	fmt.Println("---------------------learnSelect----------------------")
	learnSelect()
}

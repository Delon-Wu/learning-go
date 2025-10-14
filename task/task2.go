package task

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"
)

// 指针
// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。

func Task2_1(p *int) {
	*p += 10
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。

func Task2_2(arr *[]int) {
	for i, v := range *arr {
		(*arr)[i] = v * 2
	}
}

// Goroutine
// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func procedure1() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
func procedure2() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
func Task2_3() {
	go procedure1()
	go procedure2()
	time.Sleep(1 * time.Second)
}

//题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
//考察点 ：协程原理、并发任务调度。

func Task2_4(funcArr []func()) {
	for i, f := range funcArr {
		go func() {
			start := time.Now()
			f()
			fmt.Printf("task %d runtime cost %d ms\n", i, time.Since(start).Milliseconds())
		}()
		time.Sleep(100 * time.Millisecond)
	}
}

//面向对象
//题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
//考察点 ：接口的定义与实现、面向对象编程风格。

type Shap interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Long  float64
	Width float64
}
type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Long
}

func (r *Rectangle) Perimeter() float64 {
	return (r.Long + r.Width) * 2
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c *Circle) Perimeter() float64 {
	return c.Radius * 2 * math.Pi
}

func Task2_5(i Shap) (p float64, a float64) {
	return i.Perimeter(), i.Area()
}

//题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//考察点 ：组合的使用、方法接收者。

type Employee struct {
	Person
	Employee   string
	EmployeeID string
}
type Person struct {
	Name string
	Age  int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("[ENPLOYEE INFO] \nname: %s, age: %d\n", e.Name, e.Age)
}

func Task2_6() {
	tom := Employee{
		Person: Person{
			Name: "tom",
			Age:  21,
		},
		EmployeeID: "11-2",
	}
	tom.PrintInfo()
}

// Channel
// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func sendOnly(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
func receiveOnly(ch <-chan int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Printf("receive %d\n", i)
		case <-ctx.Done():
			fmt.Println("timeout: ", ctx.Err())
			return
		}
	}
}

func Task2_7() {
	ch := make(chan int)
	go sendOnly(ch)
	go receiveOnly(ch)
	time.Sleep(time.Second * 2)
}

//题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
//考察点 ：通道的缓冲机制。

func sendOnly1(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func Task2_8() {
	ch := make(chan int, 10)

	go sendOnly1(ch)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Printf("receive %d\n", i)
		case <-ctx.Done():
			fmt.Println("timeout: ", ctx.Err())
			return
		}
	}
}

//锁机制
//题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//考察点 ： sync.Mutex 的使用、并发数据安全。

type SafeStruct struct {
	mu sync.Mutex
	A  int
}

func (s *SafeStruct) increment() {
	for i := 0; i < 1000; i++ {
		s.mu.Lock()
		s.A++
		s.mu.Unlock()
	}
}

func Task2_9() {
	s := &SafeStruct{A: 0}
	for i := 0; i < 10; i++ {
		go s.increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("The value of s.A after 1 second:", s.A)
}

// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
type UnsafeStruct struct {
	A int
}

func (s *UnsafeStruct) increment() {
	for i := 0; i < 1000; i++ {
		s.A++
	}
}

func Task2_10() {
	s := &UnsafeStruct{A: 0}
	for i := 0; i < 10; i++ {
		go s.increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("The value of s.A after 1 second:", s.A)
}

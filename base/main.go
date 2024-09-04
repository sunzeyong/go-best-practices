package main

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
)

func main() {
	twoDefer()
}

// defer 执行函数参数在定义的时候就确定了
// 但是 执行函数内部的变量运行时才确定
func twoDefer() {
	first()
	second()
}

// 直接调用fmt函数，相当于调用函数时需要传递一个参数，这个参数是在定义的时候就已经决定了
func first() {
	i := 1
	// first i: 1
	defer fmt.Printf("first i: %d\n", i)
	i++
}

// 其实这里的defer里面是个闭包函数 运行到这个函数时，会根据函数外的这个值确定
func second() {
	i := 1
	defer func() {
		// second i: 2
		fmt.Printf("second i: %d\n", i)
	}()
	i++
}

// defer的一个例子
type Test struct {
	Max int
}

func (t *Test) Println() {
	fmt.Println(t.Max)
}

func deferExec(f func()) {
	f()
}

func call() {
	var t *Test
	// panic 参数是先传入 执行等return
	defer deferExec(t.Println)

	t = new(Test)
}

func callV2() {
	t := new(Test)
	defer deferExec(t.Println)
}

// new一个slice 是否可以直接执行spend
func newSlice() {
	s := new([]int)
	*s = append(*s, 10)
	// &[10]
	fmt.Println(s)
}

// 空slice不能直接通过index赋值
func emptySlice() {
	s := make([]int, 0)
	fmt.Printf("len: %v, cap: %v", len(s), cap(s))
	s[0] = 1
}

// cancel函数可以传递，子goroutine可以控制所有的goroutine是否退出
func ctxUse() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func(c context.CancelFunc) {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		c()
	}(cancel)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine over by cancel")
				return

			default:
				fmt.Println("go routine is running")
				time.Sleep(time.Second / 2)
			}

		}
	}()

	wg.Wait()
	fmt.Println("main goroutine is over")
}

// sort slice方法 实现排序
func sortSlice() {
	a := []int{4, 6, 1, 3, 0, 10}
	fmt.Println(a)

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	fmt.Println(a)
}

// switch两种判断方式
// 1. switch后接变量 case中枚举值和前面变量对比
// 2. switch后不接任何东西 逐个判断case条件 看哪个第一个满足
// 在判断过程中，只要有一个满足条件就会跳出，不需要添加break，若需下一个也执行，使用fallthrough
func SwitchUse() {
	mark := 80

	var grade string
	switch mark {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 70, 60:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B":
		fmt.Println("良好")
	case grade == "C":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}

// 类型断言
// 结果中i是类型转换后的结果
func SwitchCheckType(x interface{}) {
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是 %T\n", i)
	case int:
		fmt.Printf("x类型是int, %d\n", i)
	case string:
		fmt.Printf("x的类型是string, %s\n", i)
	}
}

// slice和map如何初始化使用
// 1. slice定义后就可以使用，不需要初始化
// 2. map未初始化，只能读，可以读取到对应的零值，写会panic
// 3. 延伸一下，定义一个map[string][]string 是可以直接对某一个key直接进行append 不需要在判断是否需要初始化
func SliceAndMapUse() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// s1定义了切片但未初始化，s2和s3都初始化了
	// 但三者都能直接用
	var s1 []string
	s2 := []string{}
	s3 := make([]string, 0)
	// pointer, s1: 0x0, s2: 0x102daf4a0, s3: 0x102daf4a0
	fmt.Printf("pointer, s1: %p, s2: %p, s3: %p\n", s1, s2, s3)
	// s1 equal nil: true
	fmt.Printf("s1 equal nil: %v\n", s1 == nil)

	s1 = append(s1, "s1")
	s2 = append(s2, "s2")
	s3 = append(s3, "s3")
	// s1: [s1], s2: [s2], s3: [s3]
	fmt.Printf("s1: %v, s2: %v, s3: %v\n", s1, s2, s3)

	// 定义map value是slice
	ms := make(map[string][]string)
	ms["key"] = append(ms["key"], "value01")
	// uninit case: [value01]
	fmt.Printf("uninit case: %v\n", ms["key"])

	// m2和m3效果一致，但是m1未初始化，只能读，若写会panic
	var m1 map[string]string
	m2 := map[string]string{}
	m3 := make(map[string]string)

	m2["k"] = "v"
	m3["k"] = "v"
	// m1: , m2: v, m3: v
	fmt.Printf("m1: %v, m2: %v, m3: %v\n", m1["k"], m2["k"], m3["k"])

	// assignment to entry in nil map
	m1["k"] = "v"
}

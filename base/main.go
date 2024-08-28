package main

import (
	"fmt"
)

func main() {
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
	fmt.Printf("pointer, s1: %p, s2: %p, s3: %p\n", s1, s2, s3)
	fmt.Printf("s1 equal nil: %v\n", s1 == nil)

	s1 = append(s1, "s1")
	s2 = append(s2, "s2")
	s3 = append(s3, "s3")
	fmt.Printf("s1: %v, s2: %v, s3: %v\n", s1, s2, s3)

	// m2和m3效果一致，但是m1未初始化，只能读，若写会panic
	var m1 map[string]string
	m2 := map[string]string{}
	m3 := make(map[string]string)

	m2["k"] = "v"
	m3["k"] = "v"
	fmt.Printf("m1: %v, m2: %v, m3: %v\n", m1["k"], m2["k"], m3["k"])

	m1["k"] = "v"
}

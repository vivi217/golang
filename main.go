package main

import (
	"context"
	"fmt"
	"unsafe"
)

func EmptySlice() {
	var s1 []int //slice的初始化
	var s2 = []int{}
	var s3 = make([]int, 0)
	var s4 = *new([]int)
	var a1 = *(*[3]int)(unsafe.Pointer(&s1))
	var a2 = *(*[3]int)(unsafe.Pointer(&s2))
	var a3 = *(*[3]int)(unsafe.Pointer(&s3))
	var a4 = *(*[3]int)(unsafe.Pointer(&s4))
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	context.Background()
}

func CapAndLen() {
	//array := []int{1,2,3,4,5,6}
	s1 := make([]int, 5, 10)
	fmt.Println(s1)
	a1 := *(*[3]int)(unsafe.Pointer(&s1))
	fmt.Println(a1)
	for i := 1; i <= 3; i++ {
		s1 = append(s1, i) //往后面加
	}
	fmt.Println(s1)
	//s1[9] = 5  // 会panic吗，是的，越界了，cap不代表有内存空间
	a1 = *(*[3]int)(unsafe.Pointer(&s1))
	fmt.Println(a1)
	//10的含义是不会扩容，就是不会copy成新的
	for i := 1; i <= 5; i++ {
		s1 = append(s1, i+10) //往后面加
	}
	fmt.Println(s1)
	a1 = *(*[3]int)(unsafe.Pointer(&s1)) //明显看到内存地址发生变化，扩容造成内存替换新的，cap自动扩容
	fmt.Println(a1)

	//在不指定cap情况下，cap和len都是10
	s2 := make([]int, 10)
	fmt.Println(s2)
	a2 := *(*[3]int)(unsafe.Pointer(&s2))
	fmt.Println(a2)
}

func SliceCopy() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	var s3 []int
	copy(s3, s1) //坑来了，注意力集中
	s4 := make([]int, 5)
	copy(s4, s1)
	fmt.Println(s1)
	fmt.Println(s2) //内容一致
	fmt.Println(s3) //内容一致
	fmt.Println(s4) //内容一致
	var a1 = *(*[3]int)(unsafe.Pointer(&s1))
	var a2 = *(*[3]int)(unsafe.Pointer(&s2)) //跟s1一样，完全复制
	var a3 = *(*[3]int)(unsafe.Pointer(&s3))
	var a4 = *(*[3]int)(unsafe.Pointer(&s4)) //新的底册数组，跟s1不一样
	fmt.Println(a1)                          //结构体里面值也一模一样，完全复制
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	s2 = append(s2, 100)
	fmt.Println("=========after append s2===================")
	fmt.Println(s1)
	fmt.Println(s2) //内容不一致
	a1 = *(*[3]int)(unsafe.Pointer(&s1))
	a2 = *(*[3]int)(unsafe.Pointer(&s2))
	fmt.Println(a1) //结构体里面值不一样， s2成为一个全新slice
	fmt.Println(a2)
}

func SliceCut() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[:3]
	s3 := s1[1:3]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	var a1 = *(*[3]int)(unsafe.Pointer(&s1))
	var a2 = *(*[3]int)(unsafe.Pointer(&s2))
	var a3 = *(*[3]int)(unsafe.Pointer(&s3))
	fmt.Println(a1) //结构体里面值也一模一样，完全复制
	fmt.Println(a2)
	fmt.Println(a3)
}

func ArrayCut() {
	s1 := [5]int{1, 2, 3, 4, 5}
	s2 := s1[:3]
	s3 := s1[1:3]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	var a1 = *(*[5]int)(unsafe.Pointer(&s1)) //证明数组的空间在栈内，打印了1，2，3，4，5
	var a2 = *(*[3]int)(unsafe.Pointer(&s2))
	var a3 = *(*[3]int)(unsafe.Pointer(&s3)) //内存地址+8
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}

func QuoteData(s *[]int) {
	*s = append(*s, 100)
}

func NoQuoteData(s []int) {
	s = append(s, 100)
}

func SliceQuote() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	var a1 = *(*[3]int)(unsafe.Pointer(&s1))
	fmt.Println(a1)
	NoQuoteData(s1)
	fmt.Println("after no quote")
	fmt.Println(s1)
	a1 = *(*[3]int)(unsafe.Pointer(&s1))
	fmt.Println(a1)

	QuoteData(&s1)
	fmt.Println("after quote")
	fmt.Println(s1)
	a1 = *(*[3]int)(unsafe.Pointer(&s1))
	fmt.Println(a1)
}

func main() {
	CapAndLen()
	//EmptySlice()
	//SliceCopy()
	//SliceCut()
	//ArrayCut()
	//SliceQuote()
}

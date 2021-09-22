package main

import "fmt"

type bytes []byte

type bytesV []byte

func (b bytesV) Append(data []byte) {
	b = append(b, data...)
}

func (b *bytes) Append(data []byte) {
	if b == nil {
		*b = []byte{}
	}
	*b = append(*b, data...)
}

func main() {
	b1 := bytesV("123")
	b1.Append([]byte("abcdefg"))
	fmt.Printf("b1:%s\n", b1)
	// b1:123

	_ = append(b1[:len(b1)-2], []byte("abcde")...)
	fmt.Printf("b1:%s\n", b1)
	// b1:123

	overrideSliceExample()

	preventOverrideSliceExample()

	b2 := bytes("123")
	b2.Append([]byte("abcdefg"))
	fmt.Printf("b2:%s\n", b2)
	//b2:123abcdefg
}

func overrideSliceExample() {
	b1 := []byte("123")
	b1p := b1[:len(b1)-2]
	b1p = append(b1p, 'b')
	fmt.Printf("b1:%s\n", b1)
	// b1:1b3
}

func preventOverrideSliceExample() {
	b1 := []byte("123")
	b1p := b1[: len(b1)-2 : len(b1)-2]
	b1p = append(b1p, 'b')
	fmt.Printf("b1:%s\n", b1)
	// b1:123
	// b1p will get a whole new array
	// modify b1p b1 will not influenced
	b1p[0] = 'a'
	fmt.Printf("b1:%s\n", b1)
}

func mutableStringExample() {
	s := "abcdefg"
	// s[2] = '2'
	// compile error: cannot assign to s[2] (strings are immutable)
	fmt.Println("s:", s)
}

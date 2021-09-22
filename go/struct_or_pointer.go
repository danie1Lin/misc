package main

import "fmt"

type Struct struct {
	content string
}

func (s Struct) Clear() {
	s.content = ""
}

type Pointer struct {
	content string
}

func (p *Pointer) Clear() {
	p.content = ""
}

func main() {

	pointer := Pointer{"hi"}
	fmt.Print(pointer)
	pointer.Clear()
	fmt.Print(pointer)
}

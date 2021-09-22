package main

import (
	"fmt"
)

type Human struct {
	Hand
	Foot
	Head int
}

type Hand struct {
	Finger int
}

type Foot struct {
	Finger int
}

func main() {
	human := Human{}

	// get: ambiguous selector human.Finger
	fmt.Println(human.Finger)

	// no compiled error
	fmt.Println(human.Head)

	// anonymous struct dominant
	// get: 1
	human2 := struct {
		Finger int
		Hand
		Foot
	}{1, Hand{2}, Foot{3}}
	fmt.Println(human2.Finger)
e

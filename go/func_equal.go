package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.DeepEqual(func() {}, func() {}))
}

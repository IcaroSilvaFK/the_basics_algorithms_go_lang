package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var structureOne struct {
		/*
		*	Starting from the principle of always putting the larger fields always up in the struct
		 */
		name   string // 16
		age    int8   // 8
		active bool   // 1
	}
	var structTwo struct {
		//!Larger memory footprint!
		age    int8   // 8
		name   string // 16
		active bool   // 1
	}

	fmt.Println(unsafe.Sizeof(structureOne))
	fmt.Println(unsafe.Sizeof(structTwo))

}

package main

import (
	"fmt"
	"runtime"

	. "./mpz"
)

import "C"

// define clear
// https://gmplib.org/manual/Integer-Arithmetic

func runStuff() {
	str := ""
	for i, mpz := 0, MpzFromString("111", 8); i < 5; i++ {
		str = mpz.String()
		for i := 0; i < 1000; i++ {
			runtime.GC()
		}
		mpz = mpz.AddMpz(mpz)
		fmt.Println("value", str, "ptr", mpz.Ptr())
	}
}

func main() {
	for i := 0; i < 10; i++ {
		runStuff()
	}
	for i := 0; i < 1000; i++ {
		runtime.GC()
	}
}

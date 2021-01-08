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
		mpz = mpz.AddMpz(mpz)
		fmt.Println("created ptr", mpz.Ptr(), "value", str)
		runtime.GC()
	}
}

func main() {
	for i := 0; i < 2; i++ {
		runStuff()
	}
	for i := 0; i < 1000; i++ {
		runtime.GC()
	}
}

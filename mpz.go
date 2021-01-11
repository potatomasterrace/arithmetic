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
	collect := make([]Mpz, 0)
	for i, mpz := 0, MpzFromString("111", 8); i < 100; i++ {
		//str = mpz.String()
		//fmt.Println("created ptr", mpz_cpy.Ptr(), "value", str)
		fmt.Println(mpz.String())
		mpz_cpy := mpz

		go func() {
			if i%2 == 0 {
				collect = append(collect, mpz_cpy)
			} else if i%3 == 0 {
				mpz = mpz.Add((mpz_cpy))
			} else if i%5 == 0 {
				mpz_cpy = mpz.Add((mpz_cpy))
			}
		}()
		for j := 0; j < 10; j++ {
			runtime.GC()
		}
		mpz = mpz_cpy
	}
}

func main() {
	for i := 0; i < 2; i++ {
		runStuff()
	}
	for i := 0; i < 2; i++ {
		runtime.GC()
	}
	fmt.Println("done")
}

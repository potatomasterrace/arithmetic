package main

type Number interface {
	ToString() string
	Add(num Number) Number
	Sub(num Number) Number
	Div(num Number) Number
	Mult(num Number) Number
	Exp(num Number) Number
}

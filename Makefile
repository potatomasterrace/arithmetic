compiler = gcc


run:
	time go build -o build/mpz
	time ./build/mpz

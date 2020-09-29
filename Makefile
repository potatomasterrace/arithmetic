compiler = gcc


run:
	go build -o build/mpz
	./build/mpz

valgrind: 
	go build -o build/mpz
	valgrind --leak-check=full --show-leak-kinds=all ./build/mpz

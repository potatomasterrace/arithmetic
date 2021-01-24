
SHELL := /bin/bash
buildDir= ./build
coverPath= ${buildDir}/coverage.out
buildExec= $(buildDir)/mpz

clean:
	rm -rf $(buildDir)/*

compile: clean
	echo "compiling" &&	time go build -a -o $(buildExec)

run: clean compile
	$(buildExec)

valgrind: 
	go build -o build/mpz
	valgrind --leak-check=full --show-leak-kinds=all ./build/mpz

test:
	go test ./mpz -coverprofile=${coverPath}
	go tool cover -html=${coverPath}

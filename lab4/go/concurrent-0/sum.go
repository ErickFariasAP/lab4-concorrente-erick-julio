package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil
	}
	return data
}

// sum all bytes of a file
func sum(filEPATH string, out chan int) {
	data := readFile(filEPATH)

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	out <- _sum
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64
	sums := make(map[int][]string)
	saida := make(chan int)
	for _, path := range os.Args[1:] {
		go sum(path, saida)
	}

	for _, path := range os.Args[1:] {
		soma := <-saida
		totalSum += int64(soma)

		sums[soma] = append(sums[soma], path)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}
}

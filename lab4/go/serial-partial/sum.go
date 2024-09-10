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
func sum(filePath string) int {
	data := readFile(filePath)

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}

	return _sum
}

func similarity(base []string, target []string) (contador int) {
	var counter int
	counter = 0
	var targetCopy []string
	targetCopy = target
	for _, i := range base {
		for j := 0; j < len(targetCopy); j++ {
			if i == targetCopy[j] {
				counter += 1
				targetCopy = append(targetCopy[:j], targetCopy[j+1:]...)
			}
		}
	}

	return counter / len(base)
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64
	sums := make(map[int][]string)
	for _, path := range os.Args[1:] {
		_sum := sum(path)

		totalSum += int64(_sum)

		sums[_sum] = append(sums[_sum], path)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}

	fileFingerprint := make(map[string]int)
	for _, x := range os.Args[1:] {
		var finger int
		finger = sum(x)
		fileFingerprint[x] = finger
	}

	for i := 0; i < len(os.Args[1:]); i++ {
		for j := i + 1; j < len(os.Args[1:]); j++ {
			file1 := os.Args[i]
			file2 := os.Args[j]
			fp1 := fileFingerprint[file1]
			fp2 := fileFingerprint[file2]
			similaridade := similarity(fp1, fp2)
		}
	}
}

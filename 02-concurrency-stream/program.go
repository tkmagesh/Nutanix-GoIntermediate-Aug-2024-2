package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	total := SumFile("data1.dat")
	WriteFile("test.txt", fmt.Sprintf("total : %d", total))
}

func SumFile(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if no, err := strconv.Atoi(line); err == nil {
			total += no
		}
	}
	return total
}

func WriteFile(fileName string, data string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	fmt.Fprintln(file, data)
}

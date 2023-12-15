package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	numMap := map[string]string{
		"zero":  "zero0zero",
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}

	file, err := os.Open("input.txt")
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var num int
	for scanner.Scan() {
		firstMatch := len(scanner.Text())
		var fnum, lnum int
		text := scanner.Text()
		println("inp: ", text)
		for k, v := range numMap {
			text = strings.ReplaceAll(text, k, v)
		}
		println("out: ", text)

		for p, c := range text {
			if c >= '0' && c <= '9' {
				if p < firstMatch {
					fnum = int(c) - 48
					firstMatch = p
				}
				lnum = int(c) - 48
			}
		}
		num = num + fnum*10 + lnum
	}

	fmt.Println(num)
}

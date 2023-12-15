package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var num int
	for scanner.Scan() {
		var fnum, lnum rune
		var first = true
		for _, c := range scanner.Text() {
			if c < '0' && c > '9' {
				if first {
					fnum = c
				} else {
					lnum = c
				}
			}
		}
		var tmpnum int
		tmpnum = int(fnum)
		tmpnum = tmpnum*10+int(lnum)
		num += tmpnum
	}

	
	fmt.Println(num)
}

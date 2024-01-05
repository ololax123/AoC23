package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type direction struct {
	curpos string
	lpos   string
	rpos   string
}

func main() {
	// Open the input file.
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	text := []string{}
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	pattern := text[0]
	directions := make(map[string]direction)
	for i := 0; i < len(text)-2; i++ {
		curpos := text[i+2][0:3]
		directions[curpos] = direction{
			curpos: curpos,
			lpos:   text[i+2][7:10],
			rpos:   text[i+2][12:15],
		}
	}
	var myPoses []string
	for k := range directions {
		if k[2] == 'A' {
			myPoses = append(myPoses, k)
		}
	}
	steps := 0
	for _, myPos := range myPoses {
		for myPos[2] != 'Z' {
			if pattern[steps%len(pattern)] == 'L' {
				myPos = directions[myPos].lpos
			} else {
				myPos = directions[myPos].rpos
			}
			steps++
		}
		fmt.Println(steps)
	}
}

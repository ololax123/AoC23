package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type seedMap struct {
	output int
	input  int
	size   int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	seeds := strings.Split(text[0][7:], " ")
	facit := 0
	var newSeeds []int
	for v, seed := range seeds {
		if v%2 == 0 && v < len(seeds)-2 {
			intSeed, _ := strconv.Atoi(seed)
			mulSeed, _ := strconv.Atoi(seeds[v+1])
			facit += mulSeed
			for i := 0; i < mulSeed; i++ {
				newSeeds = append(newSeeds, intSeed)
				intSeed++
			}
		}
	}

	text = text[3:]
	i := 0
	seedmaps := make([][]seedMap, 7)
	for _, line := range text {
		if strings.Contains(line, ":") {
			i++
		}
		if strings.Contains(line, "-") || line == "" {
			continue
		}
		seedinfo := strings.Split(line, " ")
		seedoutput, _ := strconv.Atoi(seedinfo[0])
		seedinput, _ := strconv.Atoi(seedinfo[1])
		seedsize, _ := strconv.Atoi(seedinfo[2])
		seedmaps[i] = append(seedmaps[i], seedMap{
			output: seedoutput,
			input:  seedinput,
			size:   seedsize,
		})
	}
	fmt.Println("Facit: ", facit)
	fmt.Println("NewSeeds: ", len(newSeeds))
	retnum2 := math.MaxInt64
	retnum := math.MaxInt64
	for _, seed := range seeds {
		seednumber, _ := strconv.Atoi(seed)
		for _, seedmap := range seedmaps {
			seednumber = modifyNumber(seedmap, seednumber)
		}
		if seednumber < retnum {
			retnum = seednumber
		}
	}
	for _, seed := range newSeeds {
		seednumber := seed
		for _, seedmap := range seedmaps {
			seednumber = modifyNumber(seedmap, seednumber)
		}
		if seednumber < retnum2 {
			retnum2 = seednumber
		}
	}

	fmt.Println("Lowest seed: ", retnum)
	fmt.Println("Lowest seed: ", retnum2-1)
}

func modifyNumber(sa []seedMap, number int) int {
	for _, s := range sa {
		if number >= s.input && number <= s.input+s.size {
			return s.output + (number - s.input)
		}
	}
	return number
}

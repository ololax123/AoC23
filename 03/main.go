package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type star struct {
	Nbr    int
	Amount int
	Bool   bool
}

var gear [][]star
var latestXY []int

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
	gear = make([][]star, len(text))
	for i := range gear {
		gear[i] = make([]star, len(text[i]))
	}
	num := 0
	foundNum := false
	firstX := 0
	numLen := 0
	for y := 0; y < len(text); y++ {
		for x := 0; x < len(text[y]); x++ {
			if isNum(string(text[y][x])) {
				if !foundNum {
					firstX = x
					foundNum = true
				}
				numLen += 1
			} else if foundNum {
				for nx := firstX; nx < x; nx++ {
					if checkaround(nx, y, text) {
						tmpnum, _ := strconv.Atoi(text[y][firstX:x])
						cy, cx := latestXY[0], latestXY[1]
						if gear[cy][cx].Bool {
							gear[cy][cx].Amount = gear[cy][cx].Amount * tmpnum
							latestXY = []int{0, 0}
						}
						num += tmpnum
						foundNum = false
						firstX = 0
						numLen = 0
						break
					}
				}
				foundNum = false
				firstX = 0
				numLen = 0

			}
			if x == len(text[y])-1 && foundNum {
				for nx := firstX; nx < x; nx++ {
					if checkaround(nx, y, text) {
						tmpnum, _ := strconv.Atoi(text[y][firstX : x+1])
						num += tmpnum
						cy, cx := latestXY[0], latestXY[1]
						if gear[cy][cx].Bool {

							gear[cy][cx].Amount = gear[cy][cx].Amount * tmpnum
							latestXY = []int{0, 0}
						}
						foundNum = false
						firstX = 0
						numLen = 0
						break
					}
				}
				foundNum = false
				firstX = 0
				numLen = 0
			}
		}

	}
	gearNum := 0
	for a, _ := range gear {
		for b, _ := range gear[a] {
			if gear[a][b].Nbr == 2 {
				gearNum += gear[a][b].Amount
			}
		}
	}
	fmt.Println(num)
	fmt.Println(gearNum)
}

func isNum(char string) bool {
	_, err := strconv.Atoi(char)
	if err != nil {
		return false
	}
	return true
}

func checkaround(x int, y int, text []string) bool {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			ny := y + dy
			nx := x + dx
			// Check if the coordinates are inside the grid
			if ny >= 0 && ny < len(text) && nx >= 0 && nx < len(text[ny]) {
				// Check if the cell contains something other than numbers and .
				if !isNum(string(text[ny][nx])) && text[ny][nx] != '.' {
					if text[ny][nx] == '*' {
						gear[ny][nx].Bool = true
						gear[ny][nx].Nbr += 1
						if gear[ny][nx].Amount < 2 {
							gear[ny][nx].Amount = 1
						}
						latestXY = []int{ny, nx}
					}
					return true
				}
			}
		}
	}
	return false
}

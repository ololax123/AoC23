package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Red   int
	Green int
	Blue  int
}

func NewGame() Game {
	return Game{
		Red:   0,
		Green: 0,
		Blue:  0,
	}
}
func main() {
	file, err := os.Open("input.txt")
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var num int
	var power int
	for scanner.Scan() {
		text := scanner.Text()
		input := strings.Split(text, ":")[1]
		games := strings.Split(input, ";")
		gameId, _ := strconv.Atoi(strings.Split(strings.Split(text, ":")[0], " ")[1])
		gameIsViable := true
		gameBoxes := NewGame()
		for _, game := range games {
			boxes := strings.Split(game, ",")
			for _, box := range boxes {
				numBoxes, _ := strconv.Atoi(strings.Split(box, " ")[1])
				color := strings.Split(box, " ")[2]
				switch color {
				case "red":
					if numBoxes > gameBoxes.Red {
						gameBoxes.Red = numBoxes
					}
					if numBoxes > 12 {
						gameIsViable = false
					}
				case "green":
					if numBoxes > gameBoxes.Green {
						gameBoxes.Green = numBoxes
					}
					if numBoxes > 13 {
						gameIsViable = false
					}
				case "blue":
					if numBoxes > gameBoxes.Blue {
						gameBoxes.Blue = numBoxes
					}
					if numBoxes > 14 {
						gameIsViable = false
					}
				}
			}
		}
		gamePower := gameBoxes.Red * gameBoxes.Green * gameBoxes.Blue
		power += gamePower
		if gameIsViable {
			num += gameId
		}
	}
	fmt.Println(num)
	fmt.Println(power)
}

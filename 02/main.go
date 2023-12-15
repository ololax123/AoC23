package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var num int
	for scanner.Scan() {
		text := scanner.Text()
		input := strings.Split(text, ":")[1]
		games := strings.Split(input, ";")
		gameId,_ := strconv.Atoi(strings.Split(strings.Split(text, ":")[0]," ")[1])
		gameIsViable := true
		for _, game := range games {
			boxes := strings.Split(game, ",")
			for _, box := range boxes {
				numBoxes,_ := strconv.Atoi(strings.Split(box, " ")[0])
				color := strings.Split(box, " ")[1]
				switch color {
				case "red":
					if numBoxes > 12 {
						gameIsViable = false
					}
				case "green":
					if numBoxes > 13 {
						gameIsViable = false
					}
				case "blue":
					if numBoxes > 14 {
						gameIsViable = false
					}
				}
			}
		}
		if gameIsViable {
			num += gameId
		}
	}	
	fmt.Println(num)
}

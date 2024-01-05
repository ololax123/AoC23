package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// card represents a poker hand with its value, cards, and bet.
type card struct {
	handvalue int
	cards     string
	bet       int
}

func main() {
	// Open the input file.
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cards []card
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cards = append(cards, parseCards(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort the cards by hand strength.
	sort.Slice(cards, func(i, j int) bool {
		return compareCards(cards[i], cards[j]) == 1
	})

	// Calculate total winnings.
	totalWin := 0
	for i, c := range cards {
		totalWin += c.bet * (len(cards) - i)
	}
	fmt.Println("Total winnings:", totalWin)
}

func parseCards(input string) card {
	bet, _ := strconv.Atoi(input[6:])
	return card{
		handvalue: handValue(input[:5]),
		cards:     input[:5],
		bet:       bet,
	}
}

func handValue(cards string) int {
	handStrength := make([]int, 15)
	for i := 0; i < 5; i++ {
		if cards[i] == 'J' {
			for j := 0; j < 15; j++ {
				handStrength[j]++
			}
		} else {
			handStrength[cardValueToInt(cards[i])]++
		}
	}
	for i := 0; i < 15; i++ {
		if handStrength[i] == 5 {
			return 6
		}
		if handStrength[i] == 4 {
			return 5
		}
		if handStrength[i] == 3 {
			for j := 0; j < 15; j++ {
				if handStrength[j] == 2 {
					return 4
				}
			}
			return 3
		}
		if handStrength[i] == 2 {
			for j := 0; j < 15; j++ {
				if handStrength[j] == 2 && j != i {
					return 2
				}
			}
			return 1
		}

	}
	return 0
}

func compareCards(card1, card2 card) int {
	if card1.handvalue > card2.handvalue {
		return 1
	} else if card1.handvalue < card2.handvalue {
		return 2
	} else {
		for i := 0; i < 5; i++ {
			if cardValueToInt(card1.cards[i]) > cardValueToInt(card2.cards[i]) {
				return 1
			} else if cardValueToInt(card1.cards[i]) < cardValueToInt(card2.cards[i]) {
				return 2
			}
		}
	}
	return 0
}

func cardValueToInt(card byte) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

func bubbleSort(cards []card) {
	for i := 0; i < len(cards); i++ {
		for j := 0; j < len(cards)-i-1; j++ {
			if compareCards(cards[j], cards[j+1]) == 1 {
				cards[j], cards[j+1] = cards[j+1], cards[j]
			}
		}
	}
}

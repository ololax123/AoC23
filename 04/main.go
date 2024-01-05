package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ticket struct {
	winning_numbers map[string]bool
	my_numbers      map[string]bool
	run             int
}

func create_ticket() ticket {
	return ticket{make(map[string]bool), make(map[string]bool), 1}
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
	tickets := make([]ticket, len(text))
	var returncount1 int
	var returncount2 int
	for n := range text {
		tickets[n] = create_ticket()
	}
	for n, line := range text {
		t := tickets[n]
		returncount2 += t.run
		text[n] = strings.Split(line, ":")[1]
		text[n] = strings.Replace(text[n], "  ", " ", -1)
		for t.run != 0 {
			for _, nbr := range strings.Split(strings.Split(text[n], "|")[0], " ") {
				t.winning_numbers[nbr] = true
			}
			for _, nbr := range strings.Split(strings.Split(text[n], "|")[1], " ") {
				t.my_numbers[nbr] = true
			}
			count1 := 0
			count2 := 0
			for nbr := range t.winning_numbers {
				if _, err := strconv.Atoi(nbr); err != nil {
					continue
				}
				if t.my_numbers[nbr] {
					if count2 == 0 {
						count2 = 1
						if t.run == 1 {
							count1 = 1
						}
					} else {
						count2++
						if t.run == 1 {
							count1 *= 2
						}
					}
				}
			}
			if t.run == 1 {
				returncount1 += count1
			}
			nadder := 1
			for count2 > 0 {

				tickets[n+nadder].run++
				nadder++
				count2--
			}
			t.run--
		}

	}
	fmt.Println("You won a total of:", returncount1)
	fmt.Println("You won a total of:", returncount2)
}

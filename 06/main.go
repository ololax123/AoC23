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

type race struct {
	time     int
	distance int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {

		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var races []race
	input := make([]string, 3)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	raceTimes := strings.Split(input[3][11:], " ")
	raceDistances := strings.Split(input[4][11:], " ")
	raceTimes = filter(raceTimes)
	raceDistances = filter(raceDistances)
	races = make([]race, len(raceTimes))
	p2Time := input[3][11:]
	p2Distance := input[4][11:]
	for i := 0; i < len(raceTimes); i++ {
		races[i].time, _ = strconv.Atoi(raceTimes[i])
		races[i].distance, _ = strconv.Atoi(raceDistances[i])
	}
	var solutions []int
	for _, race := range races {
		minPressButton := math.MaxInt64
		maxPressButton := 0
		for i := 1; i < race.time; i++ {
			distance := (race.time - i) * i
			if distance > race.distance {
				if i < minPressButton {
					minPressButton = i
				}
				if i > maxPressButton {
					maxPressButton = i
				}
			}
		}
		solutions = append(solutions, maxPressButton-minPressButton+1)
	}
	mulWin := 1
	for _, solution := range solutions {
		mulWin *= solution
	}
	p2Time = strings.Replace(p2Time, " ", "", -1)
	p2Distance = strings.Replace(p2Distance, " ", "", -1)
	p2Timeint, _ := strconv.Atoi(p2Time)
	p2Distanceint, _ := strconv.Atoi(p2Distance)
	minPressButton := math.MaxInt64
	maxPressButton := 0
	left, right := 1, p2Timeint
	for left < right {
		mid := (left + right) / 2
		distance := (p2Timeint - mid) * mid
		if distance > p2Distanceint {
			right = mid
		} else {
			left = mid + 1
		}
	}
	minPressButton = right
	left, right = 1, p2Timeint
	for left < right {
		mid := (left + right + 1) / 2 // Add 1 to avoid infinite loop
		distance := (p2Timeint - mid) * mid
		if distance > p2Distanceint {
			left = mid
		} else {
			right = mid - 1
		}
	}
	maxPressButton = left
	solutions = append(solutions, maxPressButton-minPressButton+1)
	fmt.Println(mulWin)
	fmt.Println(maxPressButton - minPressButton + 1)

}

func filter(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" && str != " " {
			r = append(r, str)
		}
	}
	return r

}

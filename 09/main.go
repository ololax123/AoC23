package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file.
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	text := []string{}
	input := [][]int{}
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	for line := range text {
		input = append(input, []int{})
		inputs := strings.Split(text[line], " ")
		for n := range inputs {
			numb, _ := strconv.Atoi(inputs[n])
			input[line] = append(input[line], numb)
		}
	}
	retNum := 0

	for _, inp := range input {
		fmt.Println("Next run:")
		//We have an array of all inputs.
		//We need to find the difference between all numbers and move them "down" a layer of the array.
		//This will need to be repeated untill we have all zeroes in all layers.
		//Then we need to add one digit to every layer. It should be "Last in current layer + last in layer below"
		layers := make([][]int, len(input))
		for i := range layers {
			layers[i] = make([]int, len(inp)-i)
			for j := 0; j < len(layers[i]); j++ {
				if i == 0 {
					layers[i][j] = inp[j]
				} else {
					layers[i][j] = layers[i-1][j+1] - layers[i-1][j]
				}
			}
		}

		for x, layer := range layers {
			for !allZero(layer) {
				newLayer := make([]int, len(layer)-1)
				for i := range newLayer {
					newLayer[i] = layer[i+1] - layer[i]
				}
				fmt.Println("NewLayer is: ", newLayer)
				layers[x][len(layers[x])-1] = newLayer
				layer = newLayer
			}
		}
		fmt.Println(layers)
		for i := len(layers) - 2; i >= 0; i-- {
			layers[i] = append(layers[i], layers[i][len(layers[i])-1]+layers[i+1][len(layers[i+1])-1])
		}
		retNum += layers[0][len(layers[0])-1]
		fmt.Println("Adding ", layers[0][len(layers[0])-1])
		fmt.Println(retNum)
	}
}
func allZero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

const base = 96

func main() {
	//define weighted average of each character
	//identify substring of slice
	//sum each slice weight
	//check input is in slice
	var input string
	var queries []int

	input, queries = readInput()
	//input, queries = readInputFromFile()

	var weights = alphabetWeights()

	values := stringsWeight(input, weights)

	for _, query := range queries {
		if _, ok := values[query]; ok {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func readInputFromFile() (string, []int) {
	var input string
	var queries []int

	file, err := os.Open("input02.in")
	if err != nil {
		panic(err)
	}

	fmt.Fscanf(file, "%s", &input)

	file, err = os.Open("input02.query.in")
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		_ = sc.Text() // GET the line string
		value, _ := strconv.Atoi(sc.Text())
		queries = append(queries, value)
		if err := sc.Err(); err != nil {
			log.Fatalf("scan file error: %v", err)
			return "", nil
		}
	}
	return input, queries
}

func stringsWeight(input string, weights map[string]int) map[int]int {
	subStrWeights := make(map[int]int)
	var temp bytes.Buffer

	for i := 0; i < len(input); i++ {
		if i == 0 || input[i] != input[i-1] {
			subStrWeights[stringWeight(string(input[i]), weights)]++
			temp.Reset()
			temp.WriteByte(input[i])
			continue
		}
		if input[i] == input[i-1] {
			temp.WriteByte(input[i])
			subStrWeights[stringWeight(temp.String(), weights)]++
		}
	}
	return subStrWeights
}

func stringWeight(s string, weights map[string]int) int {
	return charWeight(s[0], weights) * len(s)
}

func charWeight(r byte, weights map[string]int) int {
	return weights[string(r)]
}

func alphabetWeights() map[string]int {
	var weight = make(map[string]int)
	for i := 97; i <= 122; i++ {
		weight[string(i)] = i - base
	}
	return weight
}

func readInput() (string, []int) {
	var input string
	var count int

	fmt.Scanf("%s", &input)
	fmt.Scanf("%d", &count)

	var queries = []int{}
	for i := 0; i < count; i++ {
		var item int
		fmt.Scanf("%d", &item)
		queries = append(queries, item)
	}

	return input, queries
}

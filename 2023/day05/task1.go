package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSmallestNumber(arr []int) int {
	var smallest = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}

	return smallest
}


func transformSingle(seed int, transformations []transformationMap) int {
	for _, t := range transformations {
		if t.sourceStart <= seed && t.sourceEnd > seed {
			seed = seed + t.transformation
			break
		}
	}

	return seed
}

func task1(r *bufio.Reader) {
	var seeds []int
	var transformations []transformationMap

	s, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(1)
	}

	s = strings.TrimLeft(s, "seeds: ")
	s = strings.TrimRight(s, "\n")

	for _, seed := range strings.Split(s, " ") {
		// fmt.Println(seed)

		seedNum, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println("There was an error while collecting seeds")
			os.Exit(1)
		}

		seeds = append(seeds, seedNum)
	}

	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if s == "\n" {
			continue
		}

		if strings.Contains(s, "map") {
			for i, seed := range seeds {
				seeds[i] = transformSingle(seed, transformations)
			}

			transformations = []transformationMap{}
			continue
		}

        t := getTransformation(s)
		transformations = append(transformations, t)
	}

	for i, seed := range seeds {
		seeds[i] = transformSingle(seed, transformations)
	}

	fmt.Println("Task 1:")
	fmt.Println(getSmallestNumber(seeds))
}


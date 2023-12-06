package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TODO - use ranges instead of brute-forcing
/*
type seedRange struct {
	start int
	end   int
}

func getSmallestRangeNumber(arr []seedRange) int {
    if len(arr) == 0 {
        return -1
    }

    smallest := arr[0].start

	for i := 1; i < len(arr); i++ {
		if arr[i].start < smallest {
			smallest = arr[i].start
		}
	}

	return smallest
}

func transformRange(seedR seedRange, transformations []transformationMap) []seedRange {
	if len(transformations) == 0 {
		return []seedRange{seedR}
	}

	ranges := []seedRange{}

	for _, t := range transformations {
        if seedR.end < t.sourceEnd && seedR.start >= t.sourceStart {
            seedR.start += t.transformation
            seedR.end += t.transformation
            ranges = append(ranges, seedR)

            return ranges
        }

        if seedR.end > t.sourceStart && seedR.end < t.sourceEnd {
            ranges = append(ranges, seedRange{
                start: t.sourceStart + t.transformation,
                end: seedR.end + t.transformation,
            })

            seedR.end = t.sourceStart - 1
            continue
        }

        if seedR.start >= t.sourceStart && seedR.start < t.sourceEnd {
            ranges = append(ranges, seedRange{
                start: seedR.start + t.transformation,
                end: t.sourceEnd + t.transformation,
            })

            seedR.start = t.sourceEnd + 1
            continue
        }
	}

    ranges = append(ranges, seedR)

	return ranges
}

func task2(r *bufio.Reader) {
	var seedRanges []seedRange
	var transformations []transformationMap

	s, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(1)
	}

	s = strings.TrimLeft(s, "seeds: ")
	s = strings.TrimRight(s, "\n")

	ss := strings.Split(s, " ")

	for i := 0; i < len(ss); i += 2 {
		start, err := strconv.Atoi(ss[i])
		if err != nil {
			fmt.Println("There was an error while collecting seeds")
			os.Exit(1)
		}
		length, err := strconv.Atoi(ss[i+1])
		if err != nil {
			fmt.Println("There was an error while collecting seeds")
			os.Exit(1)
		}

		seedRanges = append(seedRanges, seedRange{start: start, end: start + length})
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
			newSeedRanges := []seedRange{}
			// fmt.Println("r", seedRanges)
			// fmt.Println("t", transformations)

			for _, seedR := range seedRanges {
				newSeedRanges = append(newSeedRanges, transformRange(seedR, transformations)...)
			}
			seedRanges = newSeedRanges

			transformations = []transformationMap{}
			continue
		}

		t := getTransformation(s)
		transformations = append(transformations, t)
	}

	newSeedRanges := []seedRange{}
	for _, seedR := range seedRanges {
		newSeedRanges = append(newSeedRanges, transformRange(seedR, transformations)...)
	}
	seedRanges = newSeedRanges

	// fmt.Println("r", seedRanges)

	fmt.Println("Task 2:")
	fmt.Println(getSmallestRangeNumber(seedRanges))
} */

func task2bf(r *bufio.Reader) {
	var seeds []int
	var transformations []transformationMap

	s, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(1)
	}

	s = strings.TrimLeft(s, "seeds: ")
	s = strings.TrimRight(s, "\n")

	ss := strings.Split(s, " ")

	for i := 0; i < len(ss); i += 2 {
		start, err := strconv.Atoi(ss[i])
		if err != nil {
			fmt.Println("There was an error while collecting seeds")
			os.Exit(1)
		}
		length, err := strconv.Atoi(ss[i+1])
		if err != nil {
			fmt.Println("There was an error while collecting seeds")
			os.Exit(1)
		}

        for j := start; j < length + start; j++ {
            seeds = append(seeds, j)
        }
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

	fmt.Println("Task 2:")
	fmt.Println(getSmallestNumber(seeds))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type transformationMap struct {
	sourceStart    int
	sourceEnd      int
	transformation int
}

func main() {
	var r *bufio.Reader

	r = getData()
    task1(r)

	/* r = getData()
	task2(r) */

	r = getData()
	task2bf(r)
}

func getData() *bufio.Reader {
	if len(os.Args) == 2 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Could not read file")
			os.Exit(1)
		}

		f.Seek(0, 0)
		return bufio.NewReader(f)
	}

	os.Stdin.Seek(0, 0)
	return bufio.NewReader(os.Stdin)
}

func getTransformation(s string) transformationMap {
	var t = transformationMap{}
    var err error

	s = strings.TrimRight(s, "\n")
	ns := strings.Split(s, " ")

    t.transformation, err = strconv.Atoi(ns[0])
	if err != nil {
		fmt.Println("There was an error while drawing maps 0")
		os.Exit(1)
	}
	t.sourceStart, err = strconv.Atoi(ns[1])
	if err != nil {
		fmt.Println("There was an error while drawing maps 1")
		os.Exit(1)
	}
	t.sourceEnd, err = strconv.Atoi(ns[2])
	if err != nil {
		fmt.Println("There was an error while drawing maps 2")
		os.Exit(1)
	}

	t.sourceEnd += t.sourceStart
	t.transformation -= t.sourceStart

    return t
}


package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/AbuzzBuzzard664/AOC-2017/Day_01"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_02"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_03"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_04"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_05"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_06"
)

func main() {
	var Ans string
	var i int16
	packages := []string{}
	files, err := os.ReadDir(".")
	check(err)
	fmt.Println("What Package would you like to run?")
	for _, val := range files {
		if !strings.ContainsRune(val.Name(), '.') {
			i++
			fmt.Printf("%v: %v\n", i, val.Name())
			packages = append(packages, val.Name())
		}
	}
	fmt.Scan(&Ans)
	val, err := strconv.ParseInt(Ans, 10, 64)
	check(err)
	switch val {
	case int64(slices.Index(packages, "Day_01")) + 1:
		Day_01.Run()
	case int64(slices.Index(packages, "Day_02")) + 1:
		Day_02.Run()
	case int64(slices.Index(packages, "Day_03")) + 1:
		Day_03.Run()
	case int64(slices.Index(packages, "Day_04")) + 1:
		Day_04.Run()
	case int64(slices.Index(packages, "Day_05")) + 1:
		Day_05.Run()
	case int64(slices.Index(packages, "Day_06")) + 1:
		Day_06.Run()
	default:
		fmt.Println("That option isn't available")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
	"github.com/AbuzzBuzzard664/AOC-2017/Day_07"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_08"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_09"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_10"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_11"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_12"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_13"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_14"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_15"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_16"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_17"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_18"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_19"
	"github.com/AbuzzBuzzard664/AOC-2017/Day_20"
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
	case int64(slices.Index(packages, "Day_07")) + 1:
		Day_07.Run()
	case int64(slices.Index(packages, "Day_08")) + 1:
		Day_08.Run()
	case int64(slices.Index(packages, "Day_09")) + 1:
		Day_09.Run()
	case int64(slices.Index(packages, "Day_10")) + 1:
		Day_10.Run()
	case int64(slices.Index(packages, "Day_11")) + 1:
		Day_11.Run()
	case int64(slices.Index(packages, "Day_12")) + 1:
		Day_12.Run()
	case int64(slices.Index(packages, "Day_13")) + 1:
		Day_13.Run()
	case int64(slices.Index(packages, "Day_14")) + 1:
		Day_14.Run()
	case int64(slices.Index(packages, "Day_15")) + 1:
		Day_15.Run()
	case int64(slices.Index(packages, "Day_16")) + 1:
		Day_16.Run()
	case int64(slices.Index(packages, "Day_17")) + 1:
		Day_17.Run()
	case int64(slices.Index(packages, "Day_18")) + 1:
		Day_18.Run()
	case int64(slices.Index(packages, "Day_19")) + 1:
		Day_19.Run()
	case int64(slices.Index(packages, "Day_20")) + 1:
		Day_20.Run()
	default:
		fmt.Println("That option isn't available")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

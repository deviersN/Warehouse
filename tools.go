package main

import (
	"fmt"
	"math"
)

func myStrCapitalize(src string) (string) {
	var dest string

	for _, v := range src {
		if v >= 97 && v <= 122 {
			dest = dest + string(v - 32)
			v = v - 33
		} else {
			dest = dest + string(v)
		}
	}
	return dest
}

func compareCoords(a Point, b Point) (ret bool) {
	ret = true
	if a.x != b.x || a.y != b.y {
		ret = false
	}
	return
}

func printMap(view [][]int8) {
	for _, line := range view {
		fmt.Println(line)
	}
}

func pythagoras(x float64, y float64) (ret int) {
	z1 := math.Pow(x, 2)
	z2 := math.Pow(y, 2)
	ret = int(math.Sqrt(z1 + z2))
	return
}

func absolute(src int) (int) {
	if src < 0 {
		src *= -1
	}
	return src
}
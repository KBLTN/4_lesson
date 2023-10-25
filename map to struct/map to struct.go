package main

import (
	"fmt"
)

type Point struct {
	X int `mapstructure:"xx"`
	Y int `mapstructure:"yy"`
}

func (p Point) method() {
	fmt.Println("call point method")
}

func main() {
	pointsMap := map[string]int{
		"xx": 123,
		"yy": 245,
	}

	for k, v := range pointsMap {
		fmt.Println(k)
		fmt.Println(v)
	}

}

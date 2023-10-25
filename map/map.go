package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	pointsMap := map[string]Point{
		"b": {
			X: 13,
			Y: 15,
		},
	}
	otherPointsMap := make(map[int]Point)
	pointsMap["a"] = Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(pointsMap)
	fmt.Println(pointsMap["a"])
	otherPointsMap[1] = Point{1, 2}
	fmt.Println(otherPointsMap)
	fmt.Println(otherPointsMap[1])

	var oneMorePointsMap map[string]Point
	if oneMorePointsMap == nil {
		oneMorePointsMap = map[string]Point{
			"a": {1, 2},
		}
	}
	fmt.Println(oneMorePointsMap["a"].X)
	fmt.Println(oneMorePointsMap["a"].Y)
	//oneMorePointsMap["a"].method()

	key := "a"
	value, ok := oneMorePointsMap[key]
	if ok {
		fmt.Printf("key=%s exist in map\n", key)
		fmt.Println(value)
	} else {
		fmt.Printf("key=%s does exist in map\n", key)
		fmt.Println(value)
	}

}

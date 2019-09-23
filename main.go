package main

import (
	"fmt"
	"math"
)

func main() {
	e := calculateEccentricity(9, 2)
	fmt.Println(e)
	//this will give answer as 0.9749960430435691
}

func calculateEccentricity(a, b float64) float64 {
	return (math.Sqrt(math.Pow(a, 2) - math.Pow(b, 2))) / a
}

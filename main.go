package main

import (
	"fmt"

	ellipse "./ellipse"
)

func main() {
	e := ellipse.Init{9, 2}

	fmt.Println(e.GetEccentricity())
	//this will give answer as 0.9749960430435691
}

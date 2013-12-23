package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Dice struct {
	min, max int
}

func (d Dice) Roll() int {
	rand.Seed( time.Now().UTC().UnixNano())
	return rand.Intn((d.max - d.min) + 1) + d.min
}

func main() {
	d1 := Dice{min: 1, max: 6}
	
	for i := 0 ; i < 1000 ; i++ {
		fmt.Println(d1.Roll())
	}
}
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var catArts = [][]string{
	{
		` /\_/\
( ^.^ )
c(")(")
`,
		` /\_/\
( ^.- )
 c(")(")
`,
	},
	{
		` /\_/\
( -.- )
 z z z
`,
		` /\_/\
( -.- )
 Z Z Z
`,
	},
	{
		` /\_/\
( o.o )
 > ^ <
`,
		` /\_/\
( O.O )
 > O <
`,
	},
}

func main() {
	rand.NewSource(time.Now().UnixNano())
	printRandomCatArt()
}

func printRandomCatArt() {
	randomIndex := rand.Intn(len(catArts))
	catPair := catArts[randomIndex]

	for i := 0; i < 10; i++ {
		for _, catArt := range catPair {
			fmt.Print("\033[H\033[2J")
			fmt.Println(catArt)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

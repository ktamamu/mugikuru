package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	AnimationLoops = 10
	FrameDelay     = 100 * time.Millisecond
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
	{
		` /\_/\
( =.= )
 ♥ ♥ ♥
`,
		` /\_/\
( =~= )
 ♥ ♥ ♥
`,
	},
	{
		` /\_/\
( -.o )
 ~(")
`,
		` /\_/\
( o.- )
 (")'
`,
	},
	{
		` /\_/\
( >.< )
 \_/\_
`,
		` /\_/\
( <.> )
 \_/\_
`,
	},
}

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := printRandomCatArt(rnd); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func printRandomCatArt(rnd *rand.Rand) error {
	randomIndex := rnd.Intn(len(catArts))
	catPair := catArts[randomIndex]

	for i := 0; i < AnimationLoops; i++ {
		for _, catArt := range catPair {
			if _, err := fmt.Print("\033[H\033[2J"); err != nil {
				return fmt.Errorf("failed to clear screen: %w", err)
			}
			if _, err := fmt.Println(catArt); err != nil {
				return fmt.Errorf("failed to print cat art: %w", err)
			}
			time.Sleep(FrameDelay)
		}
	}
	return nil
}

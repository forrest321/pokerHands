# pokerHands

This library contains a poker hand evaluator written in pure go, with no external dependencies.  There were several 
packages already available that could do this, but they used large lookup tables or confusing bitwise operations.  
I wanted something more readable, so I created this package.

Instructions for installation and use are below. If you'd rather just download something and run it, 
check out [pokerBench](https://github.com/forrest321/pokerBench) which was created for exactly that reason.

Currently, pokerHands only handles 5 card hands, but it is possible to add more game types in the future.

## Installing

First, add the package to your imports:
```go
import ph "github.com/forrest321/pokerHands"
```
Get the package
```shell
 go get -u github.com/forrest321/pokerHands
```
Make sure it is in your mod file:
```shell
go mod tidy
```

## Use
You can use it like so:
```go
package main

import (
	"fmt"
	ph "github.com/forrest321/pokerHands"
)

func main() {
	var hands []ph.Hand
	deck := ph.NewDeck()
	deck.Shuffle()
	for i := 0; i < 5; i++ {
		hand := ph.Hand{}
		hand.Cards = deck.Deal(0, 5)
		hand.Evaluate()
		fmt.Println(hand.String())
		hands = append(hands, hand)
	}
	winners := ph.FindWinners(hands)
	fmt.Println("winners")
	for _, w := range winners {
		fmt.Println(w.String())
	}
}
```

## TODO:
- [X] Fix broken tests
- [ ] Cleanup
- [ ] Fix exports, only export needed funcs
- [ ] Config?
- [ ] Add ID to hand to indicate its owner
- [ ] Simplify Dealing / New hands
- [ ] Fix String()s
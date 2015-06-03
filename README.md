# gondorcet

[![GoDoc](https://godoc.org/github.com/vsayer/gondorcet?status.svg)](https://godoc.org/github.com/vsayer/gondorcet) [![Build Status](https://travis-ci.org/vsayer/gondorcet.svg?branch=master)](https://travis-ci.org/vsayer/gondorcet)

Gondorcet is a library of condorcet methods for gophers.

## Install
Install gondorcet as you would any other Go program:
```shell
go get github.com/vsayer/gondorcet
```

## Example Usage
```go
package main

import (
	"fmt"
	"github.com/vsayer/gondorcet"
)

func main() {
	// create election with candidates "A", "B", and "C"
	election := gondorcet.NewElection("A", "B", "C")

	// or create election and add candidates ad-hoc
	//election := gondorcet.NewElection()
	//election.AddCandidate("A")
	//election.AddCandidate("B")
	//election.AddCandidate("C")

	// add votes
	election.AddVote("A>B>C")
	election.AddVote("B>C>A")
	election.AddVote("A>B>C")

	// tally votes
	election.TallyCopeland()

	// show election results
	fmt.Println("results:", e.GetResults())
	fmt.Println("winner:", e.GetWinner())
}
```

## Features
* supports copeland's method
* add candidates to an election adhoc

## Roadmap for v1.0
- [ ] kemeny-young method
- [ ] minimax
- [ ] nanson's method
- [ ] dodgson's method
- [ ] ranked pairs
- [ ] schulze method
- [ ] unit tests courtesy of [GoConvey](http://goconvey.co/)
- [ ] readthedocs integration
- [ ] lint for more idiomatic code
- [ ] refactor for more idiomatic code
- [ ] logo

## License
[BSD 3-Clause](LICENSE)
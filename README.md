# Queue-it Proof-of-Work

[![Go Reference](https://pkg.go.dev/badge/github.com/etaaa/Queue-it-Proof-of-Work.svg)](https://pkg.go.dev/github.com/etaaa/Queue-it-Proof-of-Work)
[![Go Report Card](https://goreportcard.com/badge/github.com/etaaa/Queue-it-Proof-of-Work)](https://goreportcard.com/report/github.com/etaaa/Queue-it-Proof-of-Work)

A Golang solution for Queue-it's Proof-of-Work challenge (<https://queue-it.com/blog/proof-of-work-block-bad-bots/>).

## Usage

Install:
```bash
go get github.com/etaaa/Queue-it-Proof-of-Work
```
Usage:
```go
package main

import (
	"fmt"
	pow "github.com/etaaa/Queue-it-Proof-of-Work"
)

func main() {
	/* Get these values from the response when fetching
	the challenge at .../challengeapi/pow/challenge/... */
	input := "f02b931c-52f0-4507-9406-f1221678dc16"
	zeroCount := 2
	// Proof of Work challenge solution
	solution := SolveChallenge(input, zeroCount)
	fmt.Println(solution)
}
```

## Questions
For any questions feel free to add and DM me on Discord (eta#0001).

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
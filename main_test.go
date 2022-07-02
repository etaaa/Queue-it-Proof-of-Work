package lib

import (
	"fmt"
	"testing"
)

func TestSolveChallenge(t *testing.T) {
	/* Get these values from the response when fetching
	the challenge at .../challengeapi/pow/challenge/... */
	input := "f02b931c-52f0-4507-9406-f1221678dc16"
	zeroCount := 2
	// Proof of Work challenge solution
	solution := SolveChallenge(input, zeroCount)
	fmt.Println(solution)
}

package lib

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	/* GET THESE VALUES FROM THE RESPONSE WHEN FETCHING
	THE CHALLENGE AT .../challengeapi/pow/challenge/... */
	input := "f02b931c-52f0-4507-9406-f1221678dc16"
	zeroCount := 2
	// RETURNS THE CHALLENGE SOLUTION
	solution := SolveChallenge(input, zeroCount)
	fmt.Println(solution)
}

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/* GET THESE VALUES FROM THE RESPONSE WHEN FETCHING
	THE CHALLENGE AT .../serviceapi/pow/challenge/... */
	input := "f02b931c-52f0-4507-9406-f1221678dc16"
	zeroCount := 4
	// RETURNS THE POSTFIX AND HASH SOLUTION
	postfix, hash := getHash(input, zeroCount)
	fmt.Println(postfix, hash)
}

func getHash(input string, zeroCount int) (int, string) {
	zeros := strings.Repeat("0", zeroCount)
	for postfix := 0; ; postfix++ {
		str := input + strconv.Itoa(postfix)
		hash := sha256.New()
		hash.Write([]byte(str))
		encodedHash := hex.EncodeToString(hash.Sum(nil))
		if strings.HasPrefix(encodedHash, zeros) {
			return postfix, encodedHash
		}
	}
}

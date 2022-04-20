package lib

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
)

func SolveChallenge(input string, zeroCount int) string {
	zeros := strings.Repeat("0", zeroCount)
	for postfix := 0; ; postfix++ {
		str := input + strconv.Itoa(postfix)
		hash := sha256.New()
		hash.Write([]byte(str))
		encodedHash := hex.EncodeToString(hash.Sum(nil))
		if strings.HasPrefix(encodedHash, zeros) {
			solutionList := []solutionType{}
			solution := solutionType{
				Hash:    encodedHash,
				Postfix: postfix,
			}
			for i := 0; i < 10; i++ {
				solutionList = append(solutionList, solution)
			}
			solutionData, _ := json.Marshal(solutionList)
			return base64.StdEncoding.EncodeToString(solutionData)
		}
	}
}

type solutionType struct {
	Hash    string `json:"hash"`
	Postfix int    `json:"postfix"`
}

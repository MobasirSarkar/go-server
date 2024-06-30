package utils

import (
	"strconv"
	"strings"
)

func Valid(luhnstring string) bool {
	checksMod := calculateCheckSum(luhnstring, false) % 10
	return checksMod == 0
}

func calculateCheckSum(luhnstring string, double bool) int {

	source := strings.Split(luhnstring, "")
	checkSum := 0

	for i := len(source) - 1; i > -1; i-- {

		t, _ := strconv.ParseInt(source[i], 10, 8)
		n := int(t)

		if double {
			n = n * 2
		}
		double = !double

		if n >= 10 {
			n = n - 9
		}

		checkSum += n

	}

	return checkSum

}

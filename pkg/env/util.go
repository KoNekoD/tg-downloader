package env

import (
	"strconv"
	"strings"
)

func asInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return result
}

func asInt64(s string) int64 {
	result, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return result
}

func asSliceOfInt64(s string) []int64 {
	items := strings.Split(s, ",")
	results := make([]int64, 0)
	for _, item := range items {
		results = append(results, asInt64(strings.TrimSpace(item)))
	}

	return results
}

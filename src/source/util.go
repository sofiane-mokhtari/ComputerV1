package source

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

func Strstr(haystack string, needle string) string {
	if needle == "" {
		return ""
	}
	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return ""
	}
	idx = idx +1
	return haystack[idx+len([]byte(needle))-1:]
}

func Conv_value(input string) float64 {
	neg := false
	if (input[0] == '-') {
		neg = true
		input = Strstr(input, " ")
	}
	i, err := strconv.ParseFloat(input, 64)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
	}
	if (neg) {
		i = i * -1
	}
	return i
}

func Conv_degre(input string) int {
	i, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
	}
	return i
}


package main

import (
	"os"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Cast struct {
	value		float64
	degre		int
}

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

func conv_value(input string) float64 {
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

func conv_degre(input string) int {
	i, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
	}
	return i
}

func casting_one(input string) Cast{
	r := regexp.MustCompile(`((\-\s):?)?(\d+(\.\d+)?)`)
	value :=  conv_value(r.FindString(input))
	degre := conv_degre(Strstr(input, "^"))
	cast := Cast{value, degre}
	return cast
}

func reduce(input []Cast, to_add Cast) []Cast {
	new_tab := make([]Cast, 0 ,1)
	for _, i := range input {
		if (i.degre == to_add.degre) {
			i.value = i.value + to_add.value
		}
		new_tab = append(new_tab, i)
	}
	return new_tab
}

func casting_reduce(input []string) []Cast {
	tab := make([]Cast, 0 ,1)
	do_reduce := false
	for _, i := range input {
		if (i == "=") {
			do_reduce = true
		} else if do_reduce {
			tab = reduce(tab, casting_one(i))
		} else {
			tab = append(tab, casting_one(i))
		}
	}
	return tab
}

func print_equation(to_print []Cast) {
	for index, i := range to_print {
		if (index > 0 && i.value > 0) {
			fmt.Print("+ ")
		}
		fmt.Print(i.value)
		fmt.Print(" * X^")
		fmt.Print(i.degre)
		fmt.Print(" ")
	}
	fmt.Println(" = 0")
}

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("Missing arg")
		return
	}
	r := regexp.MustCompile(`((((\+|\-)\s):?)?(\d+(\.\d+)?)\s[*]\s[X][\^]\d)|[=]`)
	matches := r.FindAllString(os.Args[1], -1)
	print_equation(casting_reduce(matches))
}

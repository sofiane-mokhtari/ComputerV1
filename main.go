package main

import (
	"os"
	"fmt"
	"regexp"
	"strconv"
	"math"
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

func casting_one(input string) (Cast, bool) {
	err := false
	degre := 0
	r := regexp.MustCompile(`((\-\s):?)?(\d+(\.\d+)?)`)
	value :=  conv_value(r.FindString(input))
	tmp := Strstr(input, "^")
	if (strings.Contains(tmp, "-")) {
		fmt.Println("Degre cannot be negative")
		err = true
	} else if (strings.Contains(tmp, ".")) {
		fmt.Println("Degre must be an integer")
		err = true
	} else {
		degre = conv_degre(Strstr(input, "^"))
	}
	cast := Cast{value, degre}
	return cast, err
}

func reduce(input []Cast, to_add Cast) []Cast {
	have_not_been_add := true
	new_tab := make([]Cast, 0 ,1)
	for _, i := range input {
		if (i.degre == to_add.degre) {
			i.value = i.value - to_add.value
			have_not_been_add = false
		}
		new_tab = append(new_tab, i)
	}
	if (have_not_been_add) {
		new_tab = append(new_tab, to_add)
	}
	return new_tab
}

func casting_reduce(input []string) ([]Cast, bool) {
	tab := make([]Cast, 0 ,1)
	do_reduce := false
	for _, i := range input {
		if (i == "=") {
			do_reduce = true
		} else if do_reduce {
			tmp, err := casting_one(i)
			if err {
				return tab, true
			}
			tab = reduce(tab, tmp)
		} else {
			tmp, err := casting_one(i)
			if err {
				return tab, true
			}
			tab = append(tab, tmp)
		}
	}
	return tab, false
}

func print_polynomial_degre(equa []Cast) int {
	higher_degre := 0
	for _, i := range equa {
		if (i.degre > higher_degre) {
			higher_degre = i.degre
		}
	}
	fmt.Println("Polynomial degree:", higher_degre)
	return higher_degre
}

func print_equation(to_print []Cast) {
	fmt.Print("Reduced form: ")
	for index, i := range to_print {
		if (index > 0 && i.value > 0) {
			fmt.Print("+ ")
			fmt.Print(i.value)
		} else if (i.value < 0) {
			fmt.Print("- ")
			fmt.Print(i.value * -1)
		} else {
			fmt.Print(i.value)
		}
		fmt.Print(" * X^")
		fmt.Print(i.degre)
		fmt.Print(" ")
	}
	fmt.Println(" = 0")
}

func get_for_degre(equa []Cast, degre int) float64 {
	for _, i := range equa {
		if (i.degre == degre) {
			return i.value
		}
	}
	return 0
}

func discriminate(equa []Cast) {
	a := get_for_degre(equa, 2)
	b := get_for_degre(equa, 1)
	c := get_for_degre(equa, 0)
	delta := b * b - 4 * a * c
	if (delta > 0) {
		delta_sqrt := math.Sqrt(delta)
		x_1 := (b * -1 + delta_sqrt) / (2 * a)
		x_2 := (b * -1 - delta_sqrt) / (2 * a)
		fmt.Println("Discriminant is strictly positive, the two solutions are :")
		fmt.Println(x_1)
		fmt.Println(x_2)
	} else if (delta == 0) {
		x := (b * -1) / (2 * a)
		fmt.Println("Discriminant is equal to zero, the only solution is :")
		fmt.Println(x)
	} else {
		fmt.Println("Their is no solution")
	}
}

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("Missing arg")
		return
	}
	r := regexp.MustCompile(`((((\+|\-)\s):?)?(\d+(\.\d+)?)\s[*]\s[X][\^](-)?(\d+(\.\d+)?))|[=]`)
	matches := r.FindAllString(os.Args[1], -1)
	tmp, err := casting_reduce(matches)
	if err == false { 
		print_equation(tmp)
		higher_degre := print_polynomial_degre(tmp)
		if (higher_degre > 2) {
			fmt.Println("The polynomial degree is stricly greater than 2, I can't solve.")
		} else if higher_degre == 2 {
			discriminate(tmp)
		}
	}
}

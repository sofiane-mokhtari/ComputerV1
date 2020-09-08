package main

import (
	"os"
	"fmt"
	"regexp"
	"source"
)

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("Missing arg")
		return
	}
	r := regexp.MustCompile(`((((\+|\-)\s):?)?(\d+(\.\d+)?)\s[*]\s[X][\^](-)?(\d+(\.\d+)?))|[=]`)
	matches := r.FindAllString(os.Args[1], -1)
	tmp, err := source.Casting_reduce(matches)
	if err == false { 
		source.Print_equation(tmp)
		higher_degre := source.Print_polynomial_degre(tmp)
		if (higher_degre > 2) {
			fmt.Println("The polynomial degree is stricly greater than 2, I can't solve.")
		} else if higher_degre == 2 {
			source.Discriminate(tmp)
		} else {
			source.Get_resultat(tmp)
		}
	}
}

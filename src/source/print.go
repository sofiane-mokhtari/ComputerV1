package source

import (
	"fmt"
)

func Print_polynomial_degre(equa []Cast) int {
	higher_degre := 0
	for _, i := range equa {
		if (i.degre > higher_degre) {
			higher_degre = i.degre
		}
	}
	fmt.Println("Polynomial degree:", higher_degre)
	return higher_degre
}

func Print_equation(to_print []Cast) {
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

func Get_resultat(equa []Cast) {
	b := get_value_for_degre(equa, 1)
	c := get_value_for_degre(equa, 0)
	fmt.Print("Their is only one solution x = ")
	fmt.Print(c * -1)
	fmt.Print(" / ")
	fmt.Println(b)
}
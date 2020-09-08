package source

import (
	"math"
	"fmt"
)

func Discriminate(equa []Cast) {
	a := get_value_for_degre(equa, 2)
	b := get_value_for_degre(equa, 1)
	c := get_value_for_degre(equa, 0)
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



func is_got_already_degre(equa []Cast, degre int) bool {
	for _, i := range equa {
		if (i.degre == degre) {
			return true
		}
	}
	return false
}


func get_value_for_degre(equa []Cast, degre int) float64 {
	for _, i := range equa {
		if (i.degre == degre) {
			return i.value
		}
	}
	return 0
}

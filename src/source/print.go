/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   print.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: smokhtar <smokhtar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/09/23 12:48:20 by smokhtar          #+#    #+#             */
/*   Updated: 2020/10/21 13:47:24 by smokhtar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package source

import (
	"fmt"
)

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

func Print_one_resultat(equa []Cast) {
	b := get_value_for_degre(equa, 1)
	c := get_value_for_degre(equa, 0)
	fmt.Print("Their is only one solution x = ")
	fmt.Print(c * -1)
	fmt.Print(" / ")
	fmt.Println(b)
}

func Print_two_resultat(ret []float64) {
	if (ret[0] > 0) {
		fmt.Println("Discriminant is strictly positive, the two solutions are :")
		fmt.Println(ret[1])
		fmt.Println(ret[2])
	} else if (ret[0] == 0) {
		fmt.Println("Discriminant is equal to zero, the only solution is :")
		fmt.Println(ret[3])
	} else {
		fmt.Println("Their is no solution")
	}
}
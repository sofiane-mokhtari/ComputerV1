/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   logic.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: smokhtar <smokhtar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/09/23 12:48:15 by smokhtar          #+#    #+#             */
/*   Updated: 2020/10/21 13:44:54 by smokhtar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package source

import (
	"math"
)

func Discriminate(equa []Cast) ([]float64) {
	ret := make([]float64, 3)
	a := get_value_for_degre(equa, 2)
	b := get_value_for_degre(equa, 1)
	c := get_value_for_degre(equa, 0)
	ret[0] = b * b - 4 * a * c
	if (ret[0] > 0) {
		delta_sqrt := math.Sqrt(ret[0])
		ret[1] = (b * -1 + delta_sqrt) / (2 * a)
		ret[2] = (b * -1 - delta_sqrt) / (2 * a)
		return ret
	} else if (ret[0] == 0) {
		ret[1] = (b * -1) / (2 * a)
		return ret
	} else {
		return ret
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

func Get_polynomial_degre(equa []Cast) int {
	higher_degre := 0
	for _, i := range equa {
		if (i.degre > higher_degre) {
			higher_degre = i.degre
		}
	}
	return higher_degre
}
/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   casting.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: smokhtar <smokhtar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/09/23 12:48:11 by smokhtar          #+#    #+#             */
/*   Updated: 2020/11/04 17:30:36 by smokhtar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package source

import (
	"fmt"
	"regexp"
	"strings"
)

func casting_one(input string) (Cast, bool) {
	err := false
	degre := 0
	r := regexp.MustCompile(`((\-\s):?)?(\d+(\.\d+)?)`)
	value := Conv_value(r.FindString(input))
	tmp := Strstr(input, "^")
	if (strings.Contains(tmp, "-")) {
		fmt.Println("Degre cannot be negative")
		err = true
	} else if (strings.Contains(tmp, ".")) {
		fmt.Println("Degre must be an integer")
		err = true
	} else if tmp != "" {
		degre = Conv_degre(Strstr(input, "^"))
	}
	cast := Cast{value, degre}
	return cast, err
}

func reduce(input []Cast, to_add Cast, before bool) []Cast {
	have_not_been_add := true
	new_tab := make([]Cast, 0 ,1)
	for _, i := range input {
		if (i.degre == to_add.degre) {
			if before {
				i.value = i.value + to_add.value
			} else {
				i.value = i.value - to_add.value
			}
			have_not_been_add = false
		}
		if (i.value != 0 || i.degre != 0) {
			new_tab = append(new_tab, i)
		}
	}
	if (have_not_been_add) {
		new_tab = append(new_tab, to_add)
	}
	return new_tab
}

func Casting_reduce(input []string) ([]Cast, bool) {
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
			tab = reduce(tab, tmp, false)
		} else {
			tmp, err := casting_one(i)
			if err {
				return tab, true
			}
			is_already_get_it := is_got_already_degre(tab, tmp.degre)
			if is_already_get_it {
				tab = reduce(tab, tmp, true)
			} else {
				tab = append(tab, tmp)
			}
		}
	}
	return tab, false
}
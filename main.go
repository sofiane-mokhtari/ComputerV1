/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: smokhtar <smokhtar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/09/23 12:47:53 by smokhtar          #+#    #+#             */
/*   Updated: 2020/11/02 11:00:44 by smokhtar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"os"
	"fmt"
	"./src/source"
)

func computer(input string) {
	tmp, err := source.Parsing(input)
	if err == false { 
		source.Print_equation(tmp)
		higher_degre := source.Get_polynomial_degre(tmp)
		fmt.Println("Polynomial degree:", higher_degre)
		if (higher_degre > 2) {
			fmt.Println("The polynomial degree is stricly greater than 2, I can't solve.")
		} else if higher_degre == 2 {
			tab_discriminate := source.Discriminate(tmp)
			source.Print_two_resultat(tab_discriminate)
		} else {
			source.Print_one_resultat(tmp)
		}
	}
}

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("Missing arg")
		return
	}
	computer(os.Args[1])
}

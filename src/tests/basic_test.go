/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   basic_test.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: smokhtar <smokhtar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/09/23 12:48:29 by smokhtar          #+#    #+#             */
/*   Updated: 2020/10/21 13:51:13 by smokhtar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package tests

import (
    "fmt"
    "source"
    "testing"
)

func TestParsing(t *testing.T) {

    inputs := [...]string{
        "",
        "2X^2 + 1 = 1",
        "5 * X^0 + 4 * X^1 = 4 * X^0",
        "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0",
    }

	for index, i := range inputs {
        tmp, err := source.Parsing(i)
        fmt.Println("\t test", index , tmp, err, "Degre : " , source.Get_polynomial_degre(tmp))
    }

}

// func TestDegre(t *testing.T) {

//     inputs := [...]string{
//         "",
//         "2X^2 + 1 = 1",
//         "5 * X^0 + 4 * X^1 = 4 * X^0",
//         "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0",
//     }

// 	for index, i := range inputs {
//         tmp, err := source.Parsing(i)
//         fmt.Println("\t test", index , tmp, err)
//     }

// }

// func TestReduce(t *testing.T) {

//     inputs := [...]string{
//         "",
//         "2X^2 + 1 = 1",
//         "5 * X^0 + 4 * X^1 = 4 * X^0",
//         "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0",
//     }

// 	for index, i := range inputs {
//         tmp, err := source.Parsing(i)
//         fmt.Println("\t test", index , tmp, err)
//     }

// }

// func TestResult(t *testing.T) {

//     inputs := [...]string{
//         "",
//         "2X^2 + 1 = 1",
//         "5 * X^0 + 4 * X^1 = 4 * X^0",
//         "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0",
//     }

// 	for index, i := range inputs {
//         tmp, err := source.Parsing(i)
//         fmt.Println("\t test", index , tmp, err)
//     }

// }
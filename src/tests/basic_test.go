/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   basic_test.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: smokhtar <smokhtar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/09/23 12:48:29 by smokhtar          #+#    #+#             */
/*   Updated: 2020/11/09 13:07:59 by smokhtar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package tests

import (
    "fmt"
    "../source"
    "strconv"
    "testing"
)

var (
    Black   = Color("\033[1;30m%s\033[0m")
    Red     = Color("\033[1;31m%s\033[0m")
    Green   = Color("\033[1;32m%s\033[0m")
    Yellow  = Color("\033[1;33m%s\033[0m")
    Purple  = Color("\033[1;34m%s\033[0m")
    Magenta = Color("\033[1;35m%s\033[0m")
    Teal    = Color("\033[1;36m%s\033[0m")
    White   = Color("\033[1;37m%s\033[0m")
  )

func Color(colorString string) func(...interface{}) string {
    sprint := func(args ...interface{}) string {
        return fmt.Sprintf(colorString,
        fmt.Sprint(args...))
    }
    return sprint
}

func castToString(input []source.Cast) string {
    if (len(input) == 0) {
        return "[]"
    }
    str := "["
    for i, e := range (input) {
        if (i >= 1) {
            str = str + " "
        }
        str = str + fmt.Sprintf("%v", e)
    }
    str = str + "]"
    return str
}

func TestComputerV1(t *testing.T) {
    doTestParsing()
}

func doTestParsing() {
    inputs := [...]string{
        "",
        "2X^2 + 1 = 1",
        "5 * X^0 + 4 * X^1 = 4 * X^0",
        "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0",
        "5 * X^0 + 13 * X^1 + 3 * X^2 = 1 * X^0 + 1 * X^1",
        "6 * X^0 + 11 * X^1 + 5 * X^2 = 1 * X^0 + 1 * X^1",
        "5 * X^0 + 3 * X^1 + 3 * X^2 = 1 * X^0 + 0 * X^1",
        "5 + 3X^1 + 3 * X^2 = 1 + 0 * X^1",
    }

    expected_cast := [...]string{
        "[]",
        "[{2 2}]",
        "[{1 0} {4 1}]",
        "[{4 0} {4 1} {-9.3 2}]",
        "[{4 0} {12 1} {3 2}]",
        "[{5 0} {10 1} {5 2}]",
        "[{4 0} {3 1} {3 2}]",
        "[{4 0} {3 1} {3 2}]",
    }

    expected_degre := [...]int{
        0,
        2,
        1,
        2,
        2,
        2,
        2,
        2,
    }

    expected_discrimant := [...]string{
        "[]",
        "[0 -0 0]",
        "[]",
        "[164.8 -0.47513146390886934 0.9052389907905898]",
        "[96 -0.3670068381445481 -3.632993161855452]",
        "[ 0 -1 0]",
        "[-39 0 0]",
        "[-39 0 0]",
    }
    var higher_degre int
    var tab_discriminate []float64
    for i, e := range inputs {
        tmp, err := source.Parsing(e)
        if expected_cast[i] != castToString(tmp) {
            fmt.Println(Red("--- FAIL:\t\tTestParsing Number " + strconv.Itoa(i) + " ", err))
        } else {
            higher_degre = source.Get_polynomial_degre(tmp)
            if expected_degre[i] == higher_degre {
                if (higher_degre == 2) {
                    tab_discriminate = source.Discriminate(tmp)
                    if fmt.Sprintf(expected_discrimant[i]) == expected_discrimant[i] {
                        fmt.Println(Green("--- SUCCES:\t\tTest Number " + strconv.Itoa(i) + " "))
                    } else {
                        fmt.Println(Red("--- FAIL:\t\tTestDiscriminate Number " + strconv.Itoa(i) ))
                    }
                } else {
                    fmt.Println(Green("--- SUCCES:\t\tTest Number " + strconv.Itoa(i) + " "))
                }
            } else {
                fmt.Println(Red("--- FAIL:\t\tTestNbDiscriminate Number " + strconv.Itoa(i) ))
            }
        }
    }
}

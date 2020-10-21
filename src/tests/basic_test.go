/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   basic_test.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: smokhtar <smokhtar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/09/23 12:48:29 by smokhtar          #+#    #+#             */
/*   Updated: 2020/10/21 15:22:16 by smokhtar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package tests

import (
    "fmt"
    "source"
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
    }

    expected := [...]string{
        "[]",
        "[]d",
        "[{1 0} {4 1}]",
        "[{4 0} {4 1} {-9.3 2}]",
    }

    for i, e := range inputs {
        tmp, err := source.Parsing(e)
        if (expected[i] != castToString(tmp)) {
            fmt.Println(Red("--- FAIL:\t\tTestParsing Number " + strconv.Itoa(i) + " ", err))
        } else {
            fmt.Println(Green("--- SUCCES:\t\tTestParsing Number " + strconv.Itoa(i) + " "))
        }
    }
}

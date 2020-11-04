/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   parsing.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: smokhtar <smokhtar@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/09/23 12:47:46 by smokhtar          #+#    #+#             */
/*   Updated: 2020/11/04 17:32:40 by smokhtar         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package source

import (
	"regexp"
)

func Parsing(input string) ([]Cast, bool) {
	// r := regexp.MustCompile(`((((\+|\-)\s):?)?(\d+(\.\d+)?)\s[*]\s[X][\^](-)?(\d+(\.\d+)?))|[=]`)
	r := regexp.MustCompile(`((((\+|\-)\s):?)?(\d+(\.\d+)?)(\s[*]\s)?[X][\^](-)?(\d+(\.\d+)?))|[=]|(((((\+|\-)\s):?)?)(\s)?(\d+(\.\d+)?))`)
	matches := r.FindAllString(input, -1)

	return Casting_reduce(matches)
}
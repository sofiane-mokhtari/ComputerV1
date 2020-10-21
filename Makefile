# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: smokhtar <smokhtar@student.42.fr>          +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2020/09/21 09:13:31 by smokhtar          #+#    #+#              #
#    Updated: 2020/10/18 15:50:10 by smokhtar         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

name = computor_v1

all: $(name)

$(name) :
	go build -o $(name)

test :
	go test ./src/tests/. -v

clean :
	rm -rf $(name)

re : clean all
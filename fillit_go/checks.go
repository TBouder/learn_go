/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   checks.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/18 23:47:53 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 15:19:12 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main

import	"os"
import	"fmt"
import	"strings"

func	ft_check_file(filename string, err error) {
	if (err != nil) {
		fmt.Println("Err : cannot read file ->", filename)
		os.Exit(-1)
	}
}

func	ft_check_char(char byte) {
	if (char != '#' && char != '.') {
		fmt.Println("Err: bad char ->", char)
		os.Exit(-1)
	}
}

func	ft_check_cr(char byte) {
	if (char != '\n') {
		fmt.Println("Err: bad char at the end of the line ->", string(char))
		os.Exit(-1)
	}
}

func	ft_check_nb_sharp(arr []string) {
	nb_sharp	:= 0
	
	nb_sharp += strings.Count(arr[0], "#")
	nb_sharp += strings.Count(arr[1], "#")
	nb_sharp += strings.Count(arr[2], "#")
	nb_sharp += strings.Count(arr[3], "#")
	if (nb_sharp != 4) {
		fmt.Println("Err: bad number of sharps ->", nb_sharp, "/ 4")
		os.Exit(-1)
	}
}

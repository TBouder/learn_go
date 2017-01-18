/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   checks.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/18 23:47:53 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 00:01:11 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main

import	"os"
import	"fmt"

func	ft_check_file(filename string, err error) {
	if err != nil {
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
		fmt.Println("Err: bad char at the end of the line -> ", string(char))
		os.Exit(-1)
	}
}

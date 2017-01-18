/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/18 18:53:06 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 00:01:51 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main

import	"os"
import	"fmt"

func	ft_error(str string) {
	fmt.Println(str)
	os.Exit(-1)
}

func	main() {
	if len(os.Args) != 2 {
		os.Exit(2)
	}

	var		prog_name = os.Args[0]
	var		file_name = os.Args[1]

	file_arr := ft_get_file(file_name)

	if (len(file_arr) == 0) {
		ft_error("Err : file is empty")
	}

	fmt.Println(prog_name)
	fmt.Println(file_name)
}

/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/18 18:53:06 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 00:20:04 by tbouder          ###   ########.fr       */
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
	prog_name 	:= os.Args[0]

	if len(os.Args) != 2 {
		fmt.Println("usage :", prog_name, "file")
		os.Exit(2)
	}

	var		file_name = os.Args[1]

	file_arr := ft_get_file(file_name)

	if (len(file_arr) == 0) {
		ft_error("Err : file is empty")
	}
	if (len(file_arr) > 26) {
		ft_error("Err : more than 26 tetriminos")
	}
}

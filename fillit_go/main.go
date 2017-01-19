/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/18 18:53:06 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 12:52:17 by tbouder          ###   ########.fr       */
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

	file_name	:= os.Args[1]
	file_arr	:= ft_get_file(file_name)
	fmt.Println("len : ", len (file_arr), " ->" ,file_arr)
	fmt.Println("len : ", len (file_arr[0]), " ->" ,file_arr[0])
	fmt.Println("len : ", len (file_arr[0][0]), " ->" ,file_arr[0][0])
}

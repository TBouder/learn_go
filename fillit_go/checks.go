/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   checks.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/18 23:47:53 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 13:50:27 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main

import	"os"
import	"fmt"
import	"strings"

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
/******************************************************************************/

var valid_pieces = [][]string {
	{"#", "#", "#", "#"},
	{"####"},

	{"##", "##"},

	{"###", ".#."},
	{"#.", "##", "#."},
	{"###", ".#."},
	{"###", ".#."},

	{".##", "##."}}
// const s string = "constant"


func	ft_check_adj(piece []string) int {
	nb_lines	:= len(piece)
	index		:= 0

	for (index < nb_lines) {
		nb_char		:= len(piece[index])
		sub_index	:= 0

		for (sub_index < nb_char) {

			// if (piece[index][sub_index] == '#') {
			// 	fmt.Println("First # found");
			// 	if (sub_index + 1 < nb_char && piece[index][sub_index + 1] == '#') {
			// 		fmt.Println("Second # found");
			// 	}
			// 	if (index + 1 < nb_lines && piece[index + 1][sub_index] == '#') {
			// 		fmt.Println("Third # found");
			// 	}
			// }
			// ft_test := []string{"..#", "#.."}
			// fmt.Println(ft_test)
			// fmt.Println(valid_pieces)




			sub_index++
		}
		index++;
	}

	// fmt.Println(piece)

	return (1)
}


func	ft_check_valid_tetriminos(file_arr [][]string) {
	nb_pieces	:= len(file_arr)
	index		:= 0

	for (index < nb_pieces) {
		ft_check_adj(file_arr[index])
		index++
		fmt.Println("")
	}

	// fmt.Println("len : ", len (file_arr), " ->" , file_arr)
	// fmt.Println("len : ", len (file_arr[0]), " ->" , file_arr[0])
	// fmt.Println("len : ", len (file_arr[0][0]), " ->" , file_arr[0][0])
}

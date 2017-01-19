/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_valid_tetriminos.go                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/19 15:12:39 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 15:44:33 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main
import	"strings"

var valid_pieces = [19][]string {
	{"#", "#", "#", "#"},
	{"####"},
	{"##", "##"},
	{"###", ".#."},
	{"#.", "##", "#."},
	{".#", "##", ".#"},
	{".#.", "###"},
	{"###", "#.."},
	{"###", "..#"},
	{"..#", "###"},
	{"#..", "###"},
	{"#.", "#.", "##"},
	{"##", "#.", "#."},
	{".#", ".#", "##"},
	{"##", ".#", ".#"},
	{"##.", ".##"},
	{".##", "##."},
	{".#", "##", "#."},
	{"#.", "##", ".#"}}

func	ft_verif_tetri(piece []string) int {
	valid_len	:= 19
	index		:= 0
	piece_len	:= len(piece)
	match		:= 0

	for (index < valid_len) {
		sub_valid_len	:=  len(valid_pieces[index])
		sub_index		:= 0

		if (sub_valid_len != piece_len) {
			index++
			continue
		}
		for (sub_index < sub_valid_len) {
			if (valid_pieces[index][sub_index] == piece[sub_index]) {
				sub_index++
				if (sub_index == sub_valid_len) {
					match = 1
				}
			} else {
				break
			}
		}
		index++
	}
	return (match)
}


func	ft_valid_tetriminos(file_arr [][]string) {
	nb_pieces	:= len(file_arr)
	index		:= 0

	for (index < nb_pieces) {
		if (ft_verif_tetri(file_arr[index]) == 0) {
			ft_error("Err : bad tetriminos")
		}

		i := 0
		for (i < len(file_arr[index])) {
			file_arr[index][i] = strings.Replace(file_arr[index][i], "#", string(index % 20 + 65), -1)
			i++
		}
		index++
	}
}

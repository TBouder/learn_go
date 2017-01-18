/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_get_pieces.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/19 00:01:30 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 00:10:44 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main

import	"io/ioutil"
import	"strings"

func	ft_get_line(index *int, sub_index *int, total_file_len int, file []byte) []string {
	sub_sub_index	:= 0
	total_line_len	:= 4
	line_arr		:= make([]string, 0)

	for sub_sub_index < total_line_len {
		if ((*index) < total_file_len) {
			ft_check_char(file[*index])
			line_arr = append(line_arr, string(file[*index]))
		}
		sub_sub_index++
		(*index) += 1
	}
	ft_check_cr(file[*index])
	(*index) += 1
	return (line_arr)
}

func	ft_get_piece(index *int, total_file_len int, file []byte) []string {
	sub_index			:= 0
	total_piece_len		:= 4
	piece_arr			:= make([]string, 0)

	for sub_index < total_piece_len {
		piece_arr = append(piece_arr, strings.Join(ft_get_line(index, &sub_index, total_file_len, file), ""))
		sub_index++
	}
	ft_check_nb_sharp(piece_arr)

	if ((*index) < total_file_len) {
		ft_check_cr(file[*index])
		(*index) += 1
	}

	return (piece_arr)
}

func	ft_get_file(filename string) []string {

	dat, err		:= ioutil.ReadFile(filename)
	file_arr		:= make([]string, 0)
	total_file_len	:= len(dat)
	index			:= 0

	ft_check_file(filename, err)
	for index < total_file_len {
		file_arr = append(file_arr, strings.Join(ft_get_piece(&index, total_file_len, dat), ""))
	}

	return (file_arr)
}

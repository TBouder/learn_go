/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_trim_pieces.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/19 12:17:13 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 12:22:31 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main

/*Will go through the piece to delete the empty raws*/
func	ft_trim_piece_rows(piece []string) []string {
	trim_tmp		:= make([]string, 0)
	i				:= 0

	for (i < 4) {
		if (string(piece[i])[0] != '.' || string(piece[i])[1] != '.' ||
			string(piece[i])[2] != '.' || string(piece[i])[3] != '.') {
			trim_tmp = append(trim_tmp, string(piece[i]))
		}
		i++
	}
	return (trim_tmp)
}

/*Will assign the the trimmed piece to trim_fnl*/
func	ft_trim_piece_col(piece []string, nb_columns int, to_trim int) []string {
	trim_fnl		:= make([]string, 0)
	i				:= 0

	for (i < nb_columns) {
		trim_fnl = append(trim_fnl, piece[i][to_trim:])
		i++
	}
	return (trim_fnl)
}

/*Will go through the piece to delete the empty columns*/
func	ft_trim_count_col(piece []string, nb_columns int) int {
	i				:= 0
	to_trim			:= 5

	for (i < nb_columns) {
		y	:= 0
		for (y < 4 && string(piece[i][y]) == ".") {
			y++
		}
		if (y < to_trim) {
			to_trim = y
		}
		i++
	}
	return (to_trim)
}

/*Launcher*/
func	ft_trim_piece(piece []string) []string {
	trim_tmp		:= ft_trim_piece_rows(piece)
	nb_columns		:= len(trim_tmp)
	to_trim			:= ft_trim_count_col(trim_tmp, nb_columns)
	trim_fnl		:= ft_trim_piece_col(trim_tmp, nb_columns, to_trim)

	// fmt.Println(trim_fnl)
	return (trim_fnl)
}

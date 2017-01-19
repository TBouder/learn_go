/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_algo.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/19 15:48:46 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 22:49:13 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main
// import	"strings"
import	"fmt"

func	ft_find_map_min_size(nb_sharps int) int {
	base	:= 4

	for ((base * base) < nb_sharps) {
		base++
	}
	return (base);
}

func	ft_init_board(map_size int) [][]string {
	board	:= make([][]string, map_size)
	i		:= 0

	for (i < map_size) {
		board[i] = make([]string, map_size)
		for x := range board {
			board[i][x] = " "
		}
		i++
	}
	return (board)
}

func	ft_copy_board(board [][]string) [][]string {

	new_board	:= ft_init_board(len(board))
	for i := range new_board {
	    copy(new_board[i], board[i])
	}
	return (new_board)
}

func	ft_print_arr(board [][]string) {
	i	:= 0

	for (i < len(board)) {
		fmt.Println(board[i])
		i++
	}
}

func	ft_print_arr_color_green(board [][]string) {
	i	:= 0

	for (i < len(board)) {
		fmt.Print("\033[32m")
		fmt.Println(board[i], "\033[0m")
		i++
	}
}

func	ft_print_arr_color_red(board [][]string) {
	i	:= 0

	for (i < len(board)) {
		fmt.Print("\033[31m")
		fmt.Println(board[i], "\033[0m")
		i++
	}
}

func	ft_populate_map(board [][]string, map_y int, map_x int, pieces []string) ([][]string, bool) {

	new_board	:= ft_copy_board(board)
	match		:= 0
	total_len	:= 0
	piece_y		:= 0

	for (piece_y < len(pieces)) {
		total_len += len(pieces[piece_y])
		piece_x	:= 0
		for (piece_x < len(pieces[piece_y])) {

			if (map_y + piece_y > len(board) - 1) {return board, false}
			if (map_x + piece_x > len(board[0]) - 1) {return board, false}

			if (pieces[piece_y][piece_x] == '.') {
				match++
			} else if (pieces[piece_y][piece_x] != '.' && new_board[map_y + piece_y][map_x + piece_x] == " ") {
				match++
				new_board[map_y + piece_y][map_x + piece_x] = string(pieces[piece_y][piece_x])
			} else {return board, false}

			piece_x++
		}
		piece_y++
	}
	if (match != total_len) {return board, false} else {return new_board, true}
}

func	ft_recursiv(pieces [][]string, index int, board [][]string, verif bool) (bool, [][]string) {

	nb_pieces	:= len(pieces)

	if (verif == true) {return true, board}
	if (index >= nb_pieces) {return true, board}

	map_y := 0
	for (map_y < len(board)) {

		map_x := 0
		for (map_x < len(board[map_y])) {

			if (index < nb_pieces) {
				new_board	:= board
				board, verif = ft_populate_map(board, map_y, map_x, pieces[index])
				if (verif == true) {
					verif, board = ft_recursiv(pieces, index + 1, board, false);
					if (verif == true) {return true, board}
					board = new_board
				}
			}
			map_x++
		}
		map_y++
	}
	return verif, board
}

func	ft_algo_launcher(pieces [][]string) {
	nb_sharps	:= len(pieces) * 4
	board		:= ft_init_board(ft_find_map_min_size(nb_sharps))
	verif		:= false

	index := 0
	for (verif == false) {
		verif, board = ft_recursiv(pieces, index, board, false)

		if (verif == false) {
			ft_print_arr_color_red(board)
			nb_sharps++
			board = ft_init_board(ft_find_map_min_size(nb_sharps))
		} else {
			ft_print_arr_color_green(board)
		}
	}

}

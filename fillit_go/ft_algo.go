/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_algo.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/19 15:48:46 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 23:32:33 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main
import	"fmt"

type s_len struct {
	board_len	int
	nb_pieces	int
}

/*	ft_populate_map ************************************************************
**	The ft_recursiv() function takes the array with the board, the y current pos
**	in the map, the x current pos, an array with the current piece and the s_len
**	struct as arguments.
**	The functions goes through the piece, on the map, and try to put it. On
**	success, it will return the new board and true.
*******************************************************************************/
func	ft_populate_map(board [][]string, map_y int, map_x int, piece []string, lens s_len) ([][]string, bool) {
	new_board		:= ft_copy_board(board)
	match			:= 0
	total_len		:= 0
	piece_y			:= 0
	piece_row		:= len(piece)
	piece_column	:= len(piece[0])

	for (piece_y < piece_row) {

		total_len += piece_column
		piece_x	:= 0
		for (piece_x < piece_column) {

			if (map_y + piece_y > lens.board_len - 1) {return board, false}
			if (map_x + piece_x > lens.board_len - 1) {return board, false}
			if (piece[piece_y][piece_x] == '.') {
				match++
			} else if (piece[piece_y][piece_x] != '.' && new_board[map_y + piece_y][map_x + piece_x] == " ") {
				match++
				new_board[map_y + piece_y][map_x + piece_x] = string(piece[piece_y][piece_x])
			} else {return board, false}
			piece_x++

		}
		piece_y++

	}
	if (match != total_len) {return board, false} else {return new_board, true}
}

/*	ft_recursiv ****************************************************************
**	The ft_recursiv() function takes the array with all the pieces, the index
**	of the current piece, the board and the lens struct as arguments.
**	Firstly, it returns [true] and the board if the last piece (index) was put
**	in the map.
**	If not, it goes through the map, perfom a copy of the old board, and launch
**	the ft_populate_map() function. If this one fails, the old board is reloaded
**	and we continue to go through the map. However, if the functions returns
**	true, we re-call the ft_recursiv() function, with the next piece.
*******************************************************************************/
func	ft_recursiv(pieces [][]string, index int, board [][]string, lens s_len) (bool, [][]string) {

	verif		:= false
	map_y		:= 0

	if (index >= lens.nb_pieces) {return true, board}
	for (map_y < lens.board_len) {

		map_x := 0
		for (map_x < lens.board_len) {

			new_board	:= board
			board, verif = ft_populate_map(board, map_y, map_x, pieces[index], lens)
			if (verif == true) {
				verif, board = ft_recursiv(pieces, index + 1, board, lens);
				if (verif == true) {return true, board}
			}
			board = new_board
			map_x++

		}
		map_y++

	}

	return verif, board
}

/*	ft_algo_launcher ***********************************************************
**	The ft_algo_launcher() function inits a board :
**	([total_number of sharps][total_number of sharps]string)
**	which is extended if the ft_recursiv() function returns false (no way to
**	make a square with this board).
*******************************************************************************/
func	ft_algo_launcher(pieces [][]string) {
	nb_sharps	:= len(pieces) * 4
	board		:= ft_init_board(ft_find_map_min_size(nb_sharps))
	verif		:= false
	index		:= 0

	fmt.Println("Starting map size :", len(board))
	for (verif == false) {

		lens := s_len {board_len: len(board), nb_pieces: len(pieces)}
		verif, board = ft_recursiv(pieces, index, board, lens)
		if (verif == false) {
			nb_sharps++
			board = ft_init_board(ft_find_map_min_size(nb_sharps))
			fmt.Println("Map too small. New map size :", len(board))
		} else {ft_print_arr(board, "\033[32m")}

	}
}

/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_tools.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/19 23:31:44 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 23:36:10 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main
import	"fmt"
import	"os"

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

func	ft_print_arr(board [][]string, color string) {
	i	:= 0

	for (i < len(board)) {
		if (color != "") {fmt.Print(color)}
		fmt.Println(board[i])
		if (color != "") {fmt.Print("\033[0m")}
		i++
	}
}

func	ft_error(str string) {
	fmt.Println(str)
	os.Exit(-1)
}

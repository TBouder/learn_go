/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/18 18:53:06 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/19 23:36:46 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package	main
import	"fmt"
import	"os"

func	main() {
	prog_name 	:= os.Args[0]

	if (len(os.Args) != 2) {
		fmt.Println("usage :", prog_name, "file")
		os.Exit(2)
	}

	file_name	:= os.Args[1]
	file_arr	:= ft_get_file(file_name)
	ft_algo_launcher(file_arr)
}

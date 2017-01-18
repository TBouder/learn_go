/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   start_serv.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbouder <tbouder@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2017/01/18 17:16:30 by tbouder           #+#    #+#             */
/*   Updated: 2017/01/18 18:05:37 by tbouder          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

/*******************************************************************************
**	Here, the data are formated this way (openweathermap API) :
**	{
**	    "name": "Tokyo",
**	    "coord": {
**	        "lon": 139.69,
**	        "lat": 35.69
**	    },
**	    "weather": [
**	        {
**	            "id": 803,
**	            "main": "Clouds",
**	            "description": "broken clouds",
**	            "icon": "04n"
**	        }
**	    ],
**	    "main": {
**	        "temp": 296.69,
**	        "pressure": 1014,
**	        "humidity": 83,
**	        "temp_min": 295.37,
**	        "temp_max": 298.15
**	    }
**	}
**	It's a json file
*******************************************************************************/

package	main
/*******************************************************************************
**	Import the standard lib to start a web server + strings + json
*******************************************************************************/
import (
    "encoding/json"
    "net/http"
	"strings"
)

/*******************************************************************************
**	We are creating a new type (t_weather), wich is a struct that contain :
**	- Name (as a string) with a tag [`json:"name"`]
**	- Main (as a struct : Kelvin as float64) with a tag [`json:"main"`]
*******************************************************************************/
type	t_weather struct {
	Name string `json:"name"`
    Main struct {Kelvin float64 `json:"temp"`} `json:"main"`
}

/*******************************************************************************
**	This function will actually do the request and populate t_weather struct.
**	It will take a string as argument (city) and return a struct (t_weather)
**	and an error (error) -> Error handling
*******************************************************************************/
func	ft_get_weather(city string) (t_weather, error) {
	var		data	t_weather

    resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=74c42413e5b08c8a3443e574762037ea&q=" + city)
	if err != nil	{return t_weather{}, err}

    defer resp.Body.Close()

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return t_weather{}, err
    }

    return data, nil
}

/*******************************************************************************
**	This function will parse the url and take the part after the second [/] :
**	http://localhost:8080/weather/paris -> [paris]
**	and will call the ft_get_weather function with this city, and print the
**	result as json
*******************************************************************************/
func	ft_weather(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]

	data, err := ft_get_weather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}
func	ft_hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello!"))
}

/*******************************************************************************
**	http.HandleFunc(FIRST, SECOND) ->	When your path is FIRST arg, will call
**										SECOND arg (function)
*******************************************************************************/

func	main() {
    http.HandleFunc("/hello", ft_hello)
	http.HandleFunc("/weather/", ft_weather)
	http.HandleFunc("/test/", func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("It's working!"))
	})
    http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"

	"google.golang.org/appengine"
)

func returnJSON(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, response)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")

	// url[1] == method
	// url[2] == param
	if len(url) > 0 {
		switch url[1] {
		case "departures":
			if utf8.RuneCountInString(url[2]) != 3 {
				fmt.Fprintln(w, "bad input")
				w.WriteHeader(http.StatusBadRequest)
			} else {
				response := getDepartures(strings.ToUpper(url[2]), r)
				returnJSON(w, response)
			}
			return
		default:
			fmt.Fprintln(w, url[1])
		}
	}

	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/", requestHandler)
	appengine.Main() // Starts the server to receive requests
}

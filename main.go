package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func returnJSON(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response[:]))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	ctx := appengine.NewContext(r)

	// url[1] == method
	// url[2] == param
	if len(url) > 0 {
		switch url[1] {
		case "departures":
			if utf8.RuneCountInString(url[2]) != 3 {
				fmt.Fprintln(w, "bad input")
				w.WriteHeader(http.StatusBadRequest)
			} else {
				if cacheRes, err := memcache.Get(ctx, r.URL.Path); err == memcache.ErrCacheMiss {
					response := getDepartures(strings.ToUpper(url[2]), r)
					memItem := &memcache.Item{
						Key:        r.URL.Path,
						Value:      response,
						Expiration: time.Duration(30) * time.Second,
					}
					memcache.Set(ctx, memItem)

					returnJSON(w, response)
				} else {
					returnJSON(w, cacheRes.Value)
				}
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

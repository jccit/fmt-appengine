package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

const updateTime int = 20

func returnJSON(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response[:]))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	reqPath := strings.Split(r.URL.Path, "/")
	ctx := appengine.NewContext(r)

	origin := r.Header.Get("Origin")

	switch origin {
	case "http://localhost:3000":
		fallthrough
	case "http://localhost:5000":
		w.Header().Set("Access-Control-Allow-Origin", origin)

	default:
		w.Header().Set("Access-Control-Allow-Origin", "https://findmytrain.app")
	}

	// method == method
	// param == param
	if len(reqPath) > 2 {
		method := reqPath[1]
		param, _ := url.PathUnescape(reqPath[2])

		switch method {
		case "departures":
			if utf8.RuneCountInString(param) != 3 {
				fmt.Fprintln(w, "bad input")
				w.WriteHeader(http.StatusBadRequest)
			} else {
				if cacheRes, err := memcache.Get(ctx, r.URL.Path); err == memcache.ErrCacheMiss {
					response := getDepartures(strings.ToUpper(param), r)
					memItem := &memcache.Item{
						Key:        r.URL.Path,
						Value:      response,
						Expiration: time.Duration(updateTime) * time.Second,
					}
					memcache.Set(ctx, memItem)

					returnJSON(w, response)
				} else {
					returnJSON(w, cacheRes.Value)
				}
			}
			return
		case "service":
			if cacheRes, err := memcache.Get(ctx, r.URL.Path); err == memcache.ErrCacheMiss {
				response := getService(param, r)

				memItem := &memcache.Item{
					Key:        r.URL.Path,
					Value:      response,
					Expiration: time.Duration(updateTime) * time.Second,
				}
				memcache.Set(ctx, memItem)

				returnJSON(w, response)
			} else {
				returnJSON(w, cacheRes.Value)
			}
			return
		}
	}

	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/", requestHandler)
	appengine.Main() // Starts the server to receive requests
}

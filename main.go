package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	//body := getURL("https://httpbin.org/headers", r)

	filter := crsSelector("SOP")
	soapReq := getDeparturesXML(filter)
	response := sendPOST(endpoint, soapReq, r)

	fmt.Fprintf(w, "%s", response)
}

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
}

package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

const apiKey string = "***REMOVED***"

const headerTemplate string = `<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:typ="http://thalesgroup.com/RTTI/2013-11-28/Token/types" xmlns:ldb="http://thalesgroup.com/RTTI/2017-10-01/ldb/">
<soap:Header>
<typ:AccessToken>
<typ:TokenValue>{{.AccessKey}}</typ:TokenValue>
</typ:AccessToken>
</soap:Header>`

const endpoint = "https://lite.realtime.nationalrail.co.uk/OpenLDBWS/ldb11.asmx"

type Header struct {
	AccessKey string
}

func getHeader() string {
	t := template.New("req header")
	t, err := t.Parse(headerTemplate)

	if err != nil {
		return err.Error()
	}

	authData := Header{AccessKey: apiKey}

	var out bytes.Buffer
	t.Execute(&out, authData)

	return out.String()
}

func sendPOST(url string, soapStr string, r *http.Request) []byte {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	soapBody := []byte(soapStr)

	resp, err := client.Post(url, "text/xml", bytes.NewBuffer(soapBody))

	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

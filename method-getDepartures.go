package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
)

const departuresXML = `<ldb:GetDepartureBoardRequest>||FILTER||</ldb:GetDepartureBoardRequest>`

type DepartureBoardResponse struct {
	Board DepartureBoard `xml:"Body>GetDepartureBoardResponse>GetStationBoardResult"`
}

type DepartureBoard struct {
	GeneratedAt string         `xml:"generatedAt"`
	Location    string         `xml:"locationName"`
	CRS         string         `xml:"crs"`
	Platform    bool           `xml:"platformAvailable"`
	Services    []BoardService `xml:"trainServices>service"`
}

func getDeparturesRequestXML(filter string) string {
	body := strings.Replace(departuresXML, "||FILTER||", filter, 1)
	parts := []string{getHeader(), "<soap:Body>", body, "</soap:Body>", "</soap:Envelope>"}
	combined := strings.Join(parts, "\n")
	return strings.Replace(combined, "\n", "", -1)
}

func getDepartures(crs string, r *http.Request) string {
	filter := crsSelector(crs)
	soapReq := getDeparturesRequestXML(filter)
	response := sendPOST(endpoint, soapReq, r)

	var parsedResponse DepartureBoardResponse
	xml.Unmarshal(response, &parsedResponse)
	converted, _ := json.Marshal(parsedResponse)

	return string(converted[:])
}

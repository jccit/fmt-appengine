package main

import (
	"encoding/xml"
	"fmt-appengine/trainproto"
	"github.com/golang/protobuf/proto"
	"net/http"
	"strings"
)

const departuresXML = `<ldb:GetDepartureBoardRequest>||FILTER||</ldb:GetDepartureBoardRequest>`

type DepartureBoardResponse struct {
	Board DepartureBoard `xml:"Body>GetDepartureBoardResponse>GetStationBoardResult"`
}

type DepartureBoard struct {
	GeneratedAt string                     `xml:"generatedAt"`
	Location    string                     `xml:"locationName"`
	CRS         string                     `xml:"crs"`
	Platform    bool                       `xml:"platformAvailable"`
	Services    []*trainproto.BoardService `xml:"trainServices>service"`
}

func getDeparturesRequestXML(filter string) string {
	body := strings.Replace(departuresXML, "||FILTER||", filter, 1)
	parts := []string{getHeader(), "<soap:Body>", body, "</soap:Body>", "</soap:Envelope>"}
	combined := strings.Join(parts, "\n")
	return strings.Replace(combined, "\n", "", -1)
}

func getDepartures(crs string, r *http.Request) []byte {
	filter := crsSelector(crs)
	soapReq := getDeparturesRequestXML(filter)
	response := sendPOST(endpoint, soapReq, r)

	var parsedResponse DepartureBoardResponse
	xml.Unmarshal(response, &parsedResponse)
	// converted, _ := json.Marshal(parsedResponse)

	protoResponse := &trainproto.DepartureBoard{
		GeneratedAt: parsedResponse.Board.GeneratedAt,
		Location:    parsedResponse.Board.Location,
		CRS:         parsedResponse.Board.CRS,
		HasPlatform: parsedResponse.Board.Platform,
		Services:    parsedResponse.Board.Services,
	}

	out, _ := proto.Marshal(protoResponse)

	return out
}

package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
)

type ServiceResponse struct {
	Detail ServiceDetail `xml:"Body>GetServiceDetailsResponse>GetServiceDetailsResult"`
}

type ServiceDetail struct {
	GeneratedAt        string         `xml:"generatedAt"`
	Location           string         `xml:"locationName"`
	CRS                string         `xml:"crs"`
	Platform           int            `xml:"platform"`
	Operator           string         `xml:"operator"`
	OperatorCode       string         `xml:"operatorCode"`
	ServiceType        string         `xml:"serviceType"`
	ScheduledDeparture string         `xml:"std"`
	EstimatedDeparture string         `xml:"etd"`
	ActualDeparture    string         `xml:"atd"`
	ScheduledArrival   string         `xml:"sta"`
	EstimatedArrival   string         `xml:"eta"`
	ActualArrival      string         `xml:"ata"`
	Previous           []CallingPoint `xml:"previousCallingPoints>callingPointList>callingPoint"`
	Next               []CallingPoint `xml:"subsequentCallingPoints>callingPointList>callingPoint"`
}

const serviceXML = `<ldb:GetServiceDetailsRequest><ldb:serviceID>||ID||</ldb:serviceID></ldb:GetServiceDetailsRequest>`

func getServiceRequestXML(id string) string {
	body := strings.Replace(serviceXML, "||ID||", id, 1)
	parts := []string{getHeader(), "<soap:Body>", body, "</soap:Body>", "</soap:Envelope>"}
	combined := strings.Join(parts, "\n")
	return strings.Replace(combined, "\n", "", -1)
}

func getService(id string, r *http.Request) []byte {
	soapReq := getServiceRequestXML(id)
	response := sendPOST(endpoint, soapReq, r)

	var parsedResponse ServiceResponse
	xml.Unmarshal(response, &parsedResponse)

	if strings.Contains(string(response[:]), "Invalid Service ID") {
		return []byte("invalid id")
	} else {
		converted, _ := json.Marshal(parsedResponse)
		return converted
	}
}

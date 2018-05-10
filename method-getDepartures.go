package main

import "strings"

const departuresXML = `<ldb:GetDepartureBoardRequest>||FILTER||</ldb:GetDepartureBoardRequest>`

func getDeparturesXML(filter string) string {
	body := strings.Replace(departuresXML, "||FILTER||", filter, 1)
	parts := []string{getHeader(), "<soap:Body>", body, "</soap:Body>", "</soap:Envelope>"}
	combined := strings.Join(parts, "\n")
	return strings.Replace(combined, "\n", "", -1)
}

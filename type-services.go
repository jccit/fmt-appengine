package main

type BoardService struct {
	ScheduledDeparture string   `xml:"std"`
	EstimatedDeparture string   `xml:"etd"`
	ScheduledArrival   string   `xml:"sta"`
	EstimatedArrival   string   `xml:"eta"`
	Platform           string   `xml:"platform"`
	Operator           string   `xml:"operator"`
	OperatorCode       string   `xml:"operatorCode"`
	ServiceType        string   `xml:"serviceType"`
	ServiceID          string   `xml:"serviceID"`
	Origin             Location `xml:"origin>location"`
	Destination        Location `xml:"destination>location"`
}

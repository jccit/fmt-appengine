package main

type Location struct {
	CRS  string `xml:"crs"`
	Name string `xml:"locationName"`
}

type CallingPoint struct {
	Location
	ScheduledTime string `xml:"st"`
	EstimatedTime string `xml:"et"`
	ActualTime    string `xml:"at"`
}

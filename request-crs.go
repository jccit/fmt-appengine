package main

import "strings"

const crsTemplate = `<ldb:crs>||CRS||</ldb:crs>`

func crsSelector(crs string) string {
	return strings.Replace(crsTemplate, "||CRS||", crs, 1)
}

package helper

import "strings"

func GenerateReffNumber(date, nik string) string {
	nik = nik[6:]
	dateTime := strings.Split(date, " ")
	dateTime[0] = strings.ReplaceAll(dateTime[0], "-", "")
	dateTime[1] = strings.ReplaceAll(dateTime[1], ":", "")

	return dateTime[0] + dateTime[1] + nik
}

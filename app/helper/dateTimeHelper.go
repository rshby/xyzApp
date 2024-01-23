package helper

import (
	"time"
)

// convert from string to time.Time
func StringToDateTime(date string) (time.Time, error) {
	// format yyyy-MM-dd HH:mm:ss
	dateTime, err := time.Parse("2006-04-02 15:04:05", date)
	if err != nil {
		return time.Time{}, err
	}

	return dateTime, nil
}

// convert from time.Time to string
func DateTimeToString(date time.Time) string {
	return date.Format("2006-04-02 15-04:05")
}

// convert bulan (angka) ke text
func MonthToText(bulan int) string {
	switch bulan {
	case 1:
		return "Januari"
	case 2:
		return "Februari"
	case 3:
		return "Maret"
	case 4:
		return "April"
	case 5:
		return "Mei"
	case 6:
		return "Juni"
	case 7:
		return "Juli"
	case 8:
		return "Agustus"
	case 9:
		return "September"
	case 10:
		return "Oktober"
	case 11:
		return "November"
	case 12:
		return "Desember"
	}

	return ""
}

// get first date and last date by month & year
func GetFirstAndLastDate(month, year int) (time.Time, time.Time) {
	firstDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	lastDate := firstDate.AddDate(0, 1, -1)

	return firstDate, lastDate
}

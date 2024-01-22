package helper

import "time"

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

package test

import (
	"fmt"
	"testing"
	"time"
)

func TestStringToDateTime(t *testing.T) {

	strDate := "2020-10-10 00:00:00"
	parse, err := time.Parse("2006-01-02 15:04:05", strDate)
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(parse)
}

func TestGetFirstDate(t *testing.T) {
	getFirstDate := func(month, year int) (time.Time, error) {
		date := fmt.Sprintf("%v-%v-01 00:00:00", year, month)
		if month < 10 {
			date = fmt.Sprintf("%v-0%v-01 00:00:00", year, month)
		}
		dateTime, err := time.Parse("2006-01-02 15:04:05", date)
		if err != nil {
			return time.Now(), err
		}

		return dateTime, nil
	}

	date, _ := getFirstDate(2, 2023)
	fmt.Println(date)
}

func TestGetFirstAndLastDate(t *testing.T) {
	firstLast := func(month, year int) (time.Time, time.Time) {
		firstDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		lastDate := firstDate.AddDate(0, 1, -1)

		return firstDate, lastDate
	}

	first, last := firstLast(12, 2024)

	fmt.Println("fist :", first)
	fmt.Println("last :", last)
}

package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
	"xyzApp/app/config"
	"xyzApp/app/helper"
	mck "xyzApp/test/mock"
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

func TestReffNumber(t *testing.T) {
	refNumber := func(date, nik string) string {
		nik = nik[6:]
		dateTime := strings.Split(date, " ")
		dateTime[0] = strings.ReplaceAll(dateTime[0], "-", "")
		dateTime[1] = strings.ReplaceAll(dateTime[1], ":", "")

		return dateTime[0] + dateTime[1] + nik
	}

	fmt.Println(refNumber("2024-01-01 10:25:30", "3310250502990002"))
}

func TestJWT(t *testing.T) {
	cfgMock := mck.NewConfigMock()

	// mock
	cfgMock.Mock.On("GetConfig").Return(&config.AppConfig{
		App:      nil,
		Database: nil,
		Jaeger:   nil,
		Jwt:      &config.Jwt{SecretKey: "sangatRahasia123"},
	})

	// test
	token, err := helper.GenerateToken(cfgMock, "reoshby@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

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

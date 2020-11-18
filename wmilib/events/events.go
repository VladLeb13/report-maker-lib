package events

import (
	"log"
	"strconv"
	"time"

	"github.com/VladLeb13/report-maker-lib/wmilib/tools"
)

type (
	Events struct {
		List []win32_NTLogEvent
	}

	win32_NTLogEvent struct {
		User           string
		LogFile        string
		Message        string
		CategoryString string
		ComputerName   string
		SourceName     string
		TimeWritten    string
	}
)

func (events *Events) Get() {
	if err := tools.ExecuteQuery(&events.List, queryСondition()); err != nil {
		log.Fatal(err)
	}
}

func queryСondition() string {
	now := time.Now().UTC()
	monthAgo, err := getMonthAgo(now)
	if err != nil {
		log.Println(err)
	}

	layout := "20060102150405.898655-000"
	stringNow := now.Format(layout)
	stringMounthAgo := monthAgo.Format(layout)

	reqСondition := "WHERE TimeWritten >= '" + stringMounthAgo + "' AND TimeWritten <= '" + stringNow + "'"

	return reqСondition
}

func getMonthAgo(now time.Time) (time.Time, error) {
	var stringMounth string
	month := int(now.Month()) - 1
	if month < 10 {
		stringMounth = "0" + strconv.Itoa(month)
	} else {
		stringMounth = strconv.Itoa(month)
	}

	var stringDay string
	day := int(now.Day())

	if (day == 31) || (day == 29) {
		day--
	}

	if day < 10 {
		stringDay = "0" + strconv.Itoa(day)
	} else {
		stringDay = strconv.Itoa(day)
	}

	stringYear := strconv.Itoa(now.Year())
	restOfDate := "T15:04:05Z"

	layout := "2006-01-02T15:04:05Z"
	ago, err := time.Parse(layout, stringYear+"-"+stringMounth+"-"+stringDay+restOfDate)
	if err != nil {
		log.Fatal(err)
		out, _ := time.Parse(layout, "0001-01-01T00:00:00Z")
		return out, err
	}

	return ago, nil

}

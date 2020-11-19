package events

import (
	"log"
	"time"

	"github.com/VladLeb13/report-maker-lib/wmilib/tools"
)

type (
	Events struct {
		List []win32_NTLogEvent
	}

	win32_NTLogEvent struct {
		User        string
		LogFile     string
		Message     string
		TimeWritten string
	}
)

func (events *Events) Get() {
	if err := tools.ExecuteQuery(&events.List, queryСondition()); err != nil {
		log.Fatal(err)
	}
}

func queryСondition() string {
	now := time.Now().UTC()
	twoWeekAgo := getTwoWeekAgo(now)

	layout := "20060102150405.898655-000"
	stringNow := now.Format(layout)
	stringTwoWeekAgo := twoWeekAgo.Format(layout)

	reqСondition := "WHERE TimeWritten >= '" + stringTwoWeekAgo + "' AND TimeWritten <= '" + stringNow + "'"

	return reqСondition
}

func getTwoWeekAgo(now time.Time) time.Time {

	timestamp := now.Unix() - 604800*2
	ago := time.Unix(timestamp, 0)

	return ago

}

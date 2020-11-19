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
	fourDayAgo := getFourDayAgo(now)

	layout := "20060102150405.898655-000"
	stringNow := now.Format(layout)
	stringTwoWeekAgo := fourDayAgo.Format(layout)

	reqСondition := "WHERE TimeWritten >= '" + stringTwoWeekAgo + "' AND TimeWritten <= '" + stringNow + "'"

	return reqСondition
}

func getFourDayAgo(now time.Time) time.Time {

	timestamp := now.Unix() - 86400*4
	ago := time.Unix(timestamp, 0)

	return ago

}

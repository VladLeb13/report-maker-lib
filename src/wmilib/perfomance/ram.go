package perfomance

import (
	"log"

	"github.com/StackExchange/wmi"
)

type (
	RAM struct {
		Data []win32_PerfFormattedData_PerfOS_Memory
	}
	win32_PerfFormattedData_PerfOS_Memory struct {
		PercentCommittedBytesInUse uint64
		AvailableMBytes            uint64
	}
)

func (ram *RAM) Get() {
	queryString := wmi.CreateQuery(&ram.Data, "")
	if err := wmi.Query(queryString, &ram.Data); err != nil {
		log.Fatal(err)
	}
}

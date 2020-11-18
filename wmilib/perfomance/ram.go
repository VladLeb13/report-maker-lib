package perfomance

import (
	"log"

	"report-maker-lib/wmilib/tools"
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
	if err := tools.ExecuteQuery(&ram.Data, ""); err != nil {
		log.Fatal(err)
	}
}

package perfomance

import (
	"log"

	"github.com/StackExchange/wmi"
)

type (
	Process struct {
		Data []win32_PerfFormattedData_PerfProc_Process
	}

	win32_PerfFormattedData_PerfProc_Process struct {
		IOReadOperationsPersec  uint64
		IOWriteOperationsPersec uint64
		Name                    string
		PercentProcessorTime    uint64
		ThreadCount             uint32
		VirtualBytesPeak        uint64
		IDProcess               uint32
	}
)

func (proc *Process) Get() {
	queryString := wmi.CreateQuery(&proc.Data, "")
	if err := wmi.Query(queryString, &proc.Data); err != nil {
		log.Fatal(err)
	}
}

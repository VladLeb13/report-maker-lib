package perfomance

import (
	"log"

	"github.com/StackExchange/wmi"
)

type (
	CPU struct {
		Data []win32_PerfFormattedData_Counters_ProcessorInformation
	}
	win32_PerfFormattedData_Counters_ProcessorInformation struct {
		PercentProcessorPerformance uint64
		PercentProcessorUtility     uint64
		ProcessorFrequency          uint64
	}
)

func (cpu *CPU) Get() {
	queryString := wmi.CreateQuery(&cpu.Data, "")
	if err := wmi.Query(queryString, &cpu.Data); err != nil {
		log.Fatal(err)
	}
}

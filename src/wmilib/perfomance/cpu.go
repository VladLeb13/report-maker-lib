package perfomance

import (
	"log"
	"wmilib/tools"
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
	if err := tools.ExecuteQuery(&cpu.Data, ""); err != nil {
		log.Fatal(err)
	}
}

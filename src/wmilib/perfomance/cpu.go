package perfomance

import (
	"log"

	"github.com/VladLeb13/report-maker-lib/wmilib/tools"
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

	averages := win32_PerfFormattedData_Counters_ProcessorInformation{}

	for _, v := range cpu.Data {
		averages.PercentProcessorPerformance = averages.PercentProcessorPerformance + v.PercentProcessorPerformance
		averages.PercentProcessorUtility = averages.PercentProcessorUtility + v.PercentProcessorUtility
		averages.ProcessorFrequency = averages.ProcessorFrequency + v.ProcessorFrequency
	}

	count := len(cpu.Data)
	averages.PercentProcessorPerformance = averages.PercentProcessorPerformance / uint64(count)
	averages.PercentProcessorUtility = averages.PercentProcessorUtility / uint64(count)
	averages.ProcessorFrequency = averages.PercentProcessorPerformance / uint64(count)

	cpu.Data = []win32_PerfFormattedData_Counters_ProcessorInformation{averages}
}

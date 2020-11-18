package hardware

import (
	"log"

	"github.com/VladLeb13/report-maker-lib/wmilib/tools"
)

type (
	CPU struct {
		Units []win32_Processor
	}

	win32_Processor struct {
		Manufacturer  string
		Name          string
		MaxClockSpeed uint32
		NumberOfCores uint32
		ThreadCount   uint32
		Caption       string
		Description   string
	}
)

func (cpu *CPU) Get() {
	if err := tools.ExecuteQuery(&cpu.Units, ""); err != nil {
		log.Fatal(err)
	}
}

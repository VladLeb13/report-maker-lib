package hardware

import (
	"log"

	"github.com/StackExchange/wmi"
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
	queryString := wmi.CreateQuery(&cpu.Units, "")
	if err := wmi.Query(queryString, &cpu.Units); err != nil {
		log.Fatal(err)
	}
}

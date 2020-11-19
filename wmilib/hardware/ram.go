package hardware

import (
	"log"

	"github.com/VladLeb13/report-maker-lib/wmilib/tools"
)

type (
	RAM struct {
		Units []win32_PhysicalMemory
	}

	win32_PhysicalMemory struct {
		Capacity      uint64
		Speed         uint32
		DeviceLocator string
		PartNumber    string
		Manufacturer  string
		Name          string
		FormFactor    uint16
	}
)

func (ram *RAM) Get() {
	if err := tools.ExecuteQuery(&ram.Units, ""); err != nil {
		log.Fatal(err)
	}
}

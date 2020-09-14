package hardware

import (
	"log"

	"github.com/VladLeb13/report-maker-lib/src/wmilib/tools"
)

type (
	RAM struct {
		Units []win32_PhysicalMemory
	}

	win32_PhysicalMemory struct {
		Capacity             uint64
		Speed                uint32
		DeviceLocator        string
		PartNumber           string
		Manufacturer         string
		Model                string
		Name                 string
		FormFactor           uint16
		OtherIdentifyingInfo string
	}
)

func (ram *RAM) Get() {
	if err := tools.ExecuteQuery(&ram.Units, ""); err != nil {
		log.Fatal(err)
	}
}

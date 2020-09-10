package hardware

import (
	"log"

	"github.com/StackExchange/wmi"
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
	}
)

func (ram *RAM) Get() {
	queryString := wmi.CreateQuery(&ram.Units, "")
	if err := wmi.Query(queryString, &ram.Units); err != nil {
		log.Fatal(err)
	}
}

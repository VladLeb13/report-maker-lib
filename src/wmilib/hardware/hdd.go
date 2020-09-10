package hardware

import (
	"log"

	"github.com/StackExchange/wmi"
)

type (
	HDD struct {
		Units []win32_DiskDrive
	}

	win32_DiskDrive struct {
		SerialNumber  string
		InterfaceType string
		Size          uint64
		Partitions    uint32
	}
)

func (hdd *HDD) Get() {
	queryString := wmi.CreateQuery(&hdd.Units, "")
	if err := wmi.Query(queryString, &hdd.Units); err != nil {
		log.Fatal(err)
	}
}

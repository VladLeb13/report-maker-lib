package hardware

import (
	"log"
	"wmilib/tools"
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
	if err := tools.ExecuteQuery(&hdd.Units, ""); err != nil {
		log.Fatal(err)
	}
}

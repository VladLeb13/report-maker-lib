package perfomance

import (
	"log"
	"wmilib/tools"
)

type (
	HDD struct {
		Data []win32_PerfFormattedData_PerfDisk_PhysicalDisk
	}
	win32_PerfFormattedData_PerfDisk_PhysicalDisk struct {
		PercentDiskTime       uint64
		DiskWriteBytesPersec  uint64
		DiskReadBytesPersec   uint64
		AvgDisksecPerTransfer uint32
		AvgDiskQueueLength    uint64
	}
)

func (hdd *HDD) Get() {
	if err := tools.ExecuteQuery(&hdd.Data, ""); err != nil {
		log.Fatal(err)
	}
}

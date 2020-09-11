package perfomance

import (
	"log"

	"github.com/StackExchange/wmi"
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
	queryString := wmi.CreateQuery(&hdd.Data, "")
	if err := wmi.Query(queryString, &hdd.Data); err != nil {
		log.Fatal(err)
	}
}

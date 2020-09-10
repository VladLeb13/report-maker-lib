package hardware

import (
	"log"

	"github.com/StackExchange/wmi"
)

type Volume struct {
	Partitions []win32_Volume
}

type win32_Volume struct {
	Name       string
	Capacity   uint32
	FreeSpace  uint32
	Label      string
	FileSystem string
	DeviceID   string
	Automount  bool
}

func (volume *Volume) Get() {
	queryString := wmi.CreateQuery(&volume.Partitions, "")
	if err := wmi.Query(queryString, &volume.Partitions); err != nil {
		log.Fatal(err)
	}
}

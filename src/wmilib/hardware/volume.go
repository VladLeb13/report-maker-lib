package hardware

import (
	"log"
	"wmilib/tools"
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
	if err := tools.ExecuteQuery(&volume.Partitions, ""); err != nil {
		log.Fatal(err)
	}
}

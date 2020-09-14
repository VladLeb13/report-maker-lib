package hardware

import (
	"log"
	"wmilib/tools"
)

type Volume struct {
	Disks []win32_LogicalDisk
}

type win32_LogicalDisk struct {
	Caption     string
	Description string
	DeviceID    string
	FileSystem  string
	FreeSpace   uint64
	Name        string
	Size        string
	VolumeName  string
}

func (volume *Volume) Get() {
	if err := tools.ExecuteQuery(&volume.Disks, ""); err != nil {
		log.Fatal(err)
	}
}

package wmilib

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/VladLeb13/report-maker-lib/src/wmilib/events"
	"github.com/VladLeb13/report-maker-lib/src/wmilib/hardware"
	"github.com/VladLeb13/report-maker-lib/src/wmilib/perfomance"
	"github.com/VladLeb13/report-maker-lib/src/wmilib/software"
)

func TestGeneral(t *testing.T) {

	TestGetCPU(t)
	TestGetHDD(t)
	TestGetMatherboard(t)
	TestGetRAM(t)
	TestGetNIC(t)
	TestGetVolume(t)

	TestGetOS(t)
	TestGetPrograms(t)
	TestGetShare(t)
	TestGetStartup(t)
	TestGetUpdate(t)

	TestGetEvent(t)
	TestGetPerfomance(t)

}

func TestGeneralParalell(t *testing.T) {
	ago := time.Now()
	go func() {
		TestGetCPU(t)
		TestGetHDD(t)
		TestGetMatherboard(t)
		TestGetRAM(t)
		TestGetNIC(t)
		TestGetVolume(t)
		fmt.Println(ago, time.Now())
	}()

	go func() {
		TestGetOS(t)
		TestGetPrograms(t)
		TestGetShare(t)
		TestGetStartup(t)
		TestGetUpdate(t)
		fmt.Println(ago, time.Now())
	}()

	go func() {
		TestGetEvent(t)
		fmt.Println(ago, time.Now())
	}()

	go func() {
		TestGetPerfomance(t)
		fmt.Println(ago, time.Now())
	}()

	fmt.Println(ago, time.Now())
	time.Sleep(40 * time.Second)

}

func TestGetCPU(t *testing.T) {
	cpu := new(hardware.CPU)
	cpu.Get()

	for _, v := range cpu.Units {
		fmt.Println(v.Caption)
	}

}

func TestGetMatherboard(t *testing.T) {
	board := new(hardware.Board)
	board.Get()

	log.Println("===== baseboard =====")
	for _, v := range board.Сomponents {
		log.Println("Manufacturer: ", v.Manufacturer)
		log.Println("Model: ", v.Model)
		log.Println("PartNumber: ", v.PartNumber)
		log.Println("Product: ", v.Product)
		log.Println("SerialNumber: ", v.SerialNumber)
		log.Println("Version: ", v.Version)
	}
	log.Println("===== bios =====")
	for _, v := range board.Bios {
		log.Println("Manufacturer: ", v.Manufacturer)
		log.Println("Name: ", v.Name)
		log.Println("SMBIOSBIOSVersion: ", v.SMBIOSBIOSVersion)

	}

}

func TestGetRAM(t *testing.T) {
	ram := new(hardware.RAM)
	ram.Get()

	log.Println("===== ram =====")
	for _, v := range ram.Units {
		log.Println("Capacity: ", v.Capacity)
		log.Println("DeviceLocator: ", v.DeviceLocator)
		log.Println("Manufacturer: ", v.Manufacturer)
		log.Println("PartNumber: ", v.PartNumber)
		log.Println("Speed: ", v.Speed)
		log.Println("Model: ", v.Model)
		log.Println("Name: ", v.Name)
		log.Println("FormFactor: ", v.FormFactor)
		log.Println("OtherIdentifyingInfo: ", v.OtherIdentifyingInfo)
	}
}

func TestGetHDD(t *testing.T) {
	hdd := new(hardware.HDD)
	hdd.Get()

	log.Println("===== hdd =====")
	for _, v := range hdd.Units {
		log.Println("InterfaceType: ", v.InterfaceType)
		log.Println("Partitions: ", v.Partitions)
		log.Println("SerialNumber: ", v.SerialNumber)
		log.Println("Size: ", v.Size)
		log.Println("Model: ", v.Model)
		log.Println("Name: ", v.Name)

	}
}

func TestGetVolume(t *testing.T) {
	volume := new(hardware.Volume)
	volume.Get()

	log.Println("===== volume  =====")
	for _, v := range volume.Disks {
		log.Println("===== volume  =====")
		log.Println("Caption: ", v.Caption)
		log.Println("Description: ", v.Description)
		log.Println("DeviceID: ", v.DeviceID)
		log.Println("FileSystem: ", v.FileSystem)
		log.Println("FreeSpace: ", v.FreeSpace)
		log.Println("Name: ", v.Name)
		log.Println("Size: ", v.Size)
		log.Println("VolumeName: ", v.VolumeName)

	}

}

func TestGetNIC(t *testing.T) {
	nic := new(hardware.NIC)
	nic.Get()

	log.Println("===== nic  =====")
	for _, v := range nic.Units {
		log.Println("===== nic  =====")
		log.Println("Description: ", v.Description)
		log.Println("IPAddress: ", v.IPAddress)
		log.Println("DefaultIPGateway: ", v.DefaultIPGateway)
		log.Println("DNSServerSearchOrder: ", v.DNSServerSearchOrder)
		log.Println("DHCPServer: ", v.DHCPServer)
		log.Println("DHCPEnabled: ", v.DHCPEnabled)

	}

}

/*
 *
 *			Программное обеспечение
 *
 */

func TestGetOS(t *testing.T) {
	os := new(software.OS)
	os.Get()
	log.Println("==== os ====")
	for _, v := range os.SystemInfo {
		fmt.Println("Domain: ", v.Domain)
		fmt.Println("UserName: ", v.UserName)
	}
	for _, v := range os.WindowsInfo {
		fmt.Println("BuildNumber: ", v.BuildNumber)
		fmt.Println("Caption :", v.Caption)
		fmt.Println("InstallDate: ", v.InstallDate)
		fmt.Println("OSArchitecture: ", v.OSArchitecture)
		fmt.Println("SerialNumber: ", v.SerialNumber)
		fmt.Println("Version: ", v.Version)
		fmt.Println("Name: ", v.Name)
	}
}

func TestGetShare(t *testing.T) {
	shared := new(software.Shared)
	shared.Get()
	log.Println("==== share ====")
	for i, v := range shared.Resource {
		fmt.Println("==== share position - ", i)
		fmt.Println("Caption: ", v.Caption)
		fmt.Println("Description: ", v.Description)
		fmt.Println("Name: ", v.Name)
		fmt.Println("Path: ", v.Path)
		fmt.Println("Status: ", v.Status)
	}
}

func TestGetStartup(t *testing.T) {
	startup := new(software.Startup)
	startup.Get()
	log.Println("==== startup ====")
	for i, v := range startup.Commands {
		fmt.Println("==== startup position - ", i)
		fmt.Println("Caption: ", v.Caption)
		fmt.Println("Command: ", v.Command)
		fmt.Println("User: ", v.User)
		fmt.Println("Name: ", v.Name)
		fmt.Println("Location: ", v.Location)
	}
}

func TestGetUpdate(t *testing.T) {
	updates := new(software.Updates)
	updates.Get()

	log.Println("==== update ====")
	for i, v := range updates.List {
		fmt.Println("==== update position - ", i)
		fmt.Println("Description: ", v.Description)
		fmt.Println("FixComments: ", v.FixComments)
		fmt.Println("HotFixID: ", v.HotFixID)
		fmt.Println("InstalledBy: ", v.InstalledBy)
		fmt.Println("InstalledOn: ", v.InstalledOn)
		fmt.Println("Name: ", v.Name)
		fmt.Println("Status: ", v.Status)
	}
}

func TestGetPrograms(t *testing.T) {
	prog := new(software.Programs)
	prog.Get()
	log.Println("==== programs ====")
	for i, v := range prog.List {
		fmt.Println("==== programs position - ", i)
		fmt.Println("Caption: ", v.Caption)
		fmt.Println("Description: ", v.Description)
		fmt.Println("InstallDate: ", v.InstallDate)
		fmt.Println("InstallLocation: ", v.InstallLocation)
		fmt.Println("Name: ", v.Name)
		fmt.Println("RegCompany: ", v.RegCompany)
		fmt.Println("RegOwner: ", v.RegOwner)
		fmt.Println("Vendor: ", v.Vendor)
		fmt.Println("Version: ", v.Version)
	}
}

/*
 *
 *			События системы
 *
 */
func TestGetEvent(t *testing.T) {
	events := new(events.Events)
	events.Get()

	log.Println("==== event ====")
	for _, v := range events.List {
		fmt.Println("==== event ====")
		fmt.Println("User: ", v.User)
		fmt.Println("LogFile: ", v.LogFile)
		fmt.Println("Message: ", v.Message)
		fmt.Println("CategoryString: ", v.CategoryString)
		fmt.Println("ComputerName: ", v.ComputerName)
		fmt.Println("Data: ", v.Data)
		fmt.Println("SourceName: ", v.SourceName)
		fmt.Println("TimeWritten: ", v.TimeWritten)
	}

}

/*
 *
 *			Производительность
 *
 */

func TestGetPerfomance(t *testing.T) {

	perfCPU := new(perfomance.CPU)
	perfCPU.Get()
	log.Println("==== perf_processor ====")
	for _, v := range perfCPU.Data {
		fmt.Println(v.PercentProcessorUtility) //исп проца
		fmt.Println(v.ProcessorFrequency)
		fmt.Println(v.PercentProcessorPerformance)
	}

	perfRAM := new(perfomance.RAM)
	perfRAM.Get()

	log.Println("==== perf_mem ====")
	for _, v := range perfRAM.Data {
		fmt.Println(v.AvailableMBytes)            //доступно
		fmt.Println(v.PercentCommittedBytesInUse) //процентов использованно
	}

	perfHDD := new(perfomance.HDD)
	perfHDD.Get()
	log.Println("==== pref_disk ====")
	for _, v := range perfHDD.Data {
		fmt.Println(v.PercentDiskTime)
		fmt.Println(v.AvgDiskQueueLength)
		fmt.Println(v.DiskReadBytesPersec)
		fmt.Println(v.DiskWriteBytesPersec)
	}

	perfProc := new(perfomance.Process)
	perfProc.Get()

	log.Println("==== perf_process ====")
	for _, v := range perfProc.Data {
		fmt.Println(v.Name)
		fmt.Println(v.IDProcess)
		fmt.Println(v.IOReadOperationsPersec)
		fmt.Println(v.IOWriteOperationsPersec)
		fmt.Println(v.PercentProcessorTime)
	}

}

package wmilib

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestGeneral(t *testing.T) {

	ago := time.Now()
	tGetCPU()
	tGetHDD()
	tGetMatherboard()
	tGetRAM()
	tGetNIC()
	tGetVolume()

	tGetOS()
	tGetPrograms()
	tGetShare()
	tGetStartup()
	tGetUpdate()

	tGetEvent()
	tGetPerfomance()

	fmt.Println(ago, time.Now())
}

func TestGeneralParalell(t *testing.T) {
	ago := time.Now()
	go func() {
		tGetCPU()
		tGetHDD()
		tGetMatherboard()
		tGetRAM()
		tGetNIC()
		tGetVolume()
		fmt.Println(ago, time.Now())
	}()

	go func() {
		tGetOS()
		tGetPrograms()
		tGetShare()
		tGetStartup()
		tGetUpdate()
		fmt.Println(ago, time.Now())
	}()

	go func() {
		tGetEvent()
		fmt.Println(ago, time.Now())
	}()

	go func() {
		tGetPerfomance()
		fmt.Println(ago, time.Now())
	}()

	fmt.Println(ago, time.Now())
	time.Sleep(40 * time.Second)

}

func tGetCPU() {
	GetCPU()
}

func tGetMatherboard() {
	baseboard := GetMotherboard()
	log.Println("===== baseboard =====")
	for _, v := range baseboard.baseboard {
		log.Println("Manufacturer: ", v.Manufacturer)
		log.Println("Model: ", v.Model)
		log.Println("PartNumber: ", v.PartNumber)
		log.Println("Product: ", v.Product)
		log.Println("SerialNumber: ", v.SerialNumber)
		log.Println("Version: ", v.Version)
	}
	log.Println("===== bios =====")
	for _, v := range baseboard.bios {
		log.Println("Manufacturer: ", v.Manufacturer)
		log.Println("Name: ", v.Name)
		log.Println("SMBIOSBIOSVersion: ", v.SMBIOSBIOSVersion)

	}

}

func tGetRAM() {
	ram := GetRAM()
	log.Println("===== ram =====")

	for _, v := range ram.ram {
		log.Println("Capacity: ", v.Capacity)
		log.Println("DeviceLocator: ", v.DeviceLocator)
		log.Println("Manufacturer: ", v.Manufacturer)
		log.Println("PartNumber: ", v.PartNumber)
	}
}

func tGetHDD() {
	hdd := GetHDD()

	log.Println("===== hdd =====")

	for _, v := range hdd.hdd {
		log.Println("InterfaceType: ", v.InterfaceType)
		log.Println("Partitions: ", v.Partitions)
		log.Println("SerialNumber: ", v.SerialNumber)
		log.Println("Size: ", v.Size)

	}
}

func tGetVolume() {
	volume := GetVolume()

	log.Println("===== volume  =====")

	for _, v := range volume.volume {
		log.Println("===== volume  =====")
		log.Println("Capacity: ", v.Capacity)
		log.Println("FileSystem: ", v.FileSystem)
		log.Println("FreeSpace: ", v.FreeSpace)
		log.Println("Label: ", v.Label)
		log.Println("Name: ", v.Name)
		log.Println("DeviceID: ", v.DeviceID)
		log.Println("Automount: ", v.Automount)

	}

}

func tGetNIC() {
	nic := GetNIC()

	log.Println("===== nic  =====")

	for _, v := range nic.nic {
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

func tGetOS() {
	info := GetOS()
	log.Println("==== os ====")
	for _, v := range info.cs {
		fmt.Println("Domain: ", v.Domain)
		fmt.Println("UserName: ", v.UserName)
	}
	for _, v := range info.os {
		fmt.Println("BuildNumber: ", v.BuildNumber)
		fmt.Println("Caption :", v.Caption)
		fmt.Println("InstallDate: ", v.InstallDate)
		fmt.Println("OSArchitecture: ", v.OSArchitecture)
		fmt.Println("SerialNumber: ", v.SerialNumber)
		fmt.Println("Version: ", v.Version)
	}

}

func tGetShare() {
	info := GetShare()
	log.Println("==== share ====")
	for i, v := range info.share {
		fmt.Println("==== share position - ", i)
		fmt.Println("Caption: ", v.Caption)
		fmt.Println("Description: ", v.Description)
		fmt.Println("Name: ", v.Name)
		fmt.Println("Path: ", v.Path)
		fmt.Println("Status: ", v.Status)
	}
}

func tGetStartup() {
	info := GetStartup()
	log.Println("==== startup ====")
	for i, v := range info.startup {
		fmt.Println("==== startup position - ", i)
		fmt.Println("Caption: ", v.Caption)
		fmt.Println("Command: ", v.Command)
		fmt.Println("User: ", v.User)
	}
}

func tGetUpdate() {
	info := GetUpdate()
	log.Println("==== update ====")
	for i, v := range info.update {
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

func tGetPrograms() {
	info := GetPrograms()
	log.Println("==== programs ====")
	for i, v := range info.programs {
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
func tGetEvent() {
	evnt := GetEvent()
	log.Println("==== event ====")
	for _, v := range evnt.events {
		log.Println("==== event ====")
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

func tGetPerfomance() {

	perf := GetPerfomance()
	log.Println("==== perf_mem ====")
	for _, v := range perf.perf_mem {
		fmt.Println(v.AvailableMBytes)            //доступно
		fmt.Println(v.PercentCommittedBytesInUse) //процентов использованно
	}
	log.Println("==== perf_process ====")
	for _, v := range perf.perf_process {
		fmt.Println(v.Name)
		fmt.Println(v.IDProcess)
		fmt.Println(v.IOReadOperationsPersec)
		fmt.Println(v.IOWriteOperationsPersec)
		fmt.Println(v.PercentProcessorTime)
	}
	log.Println("==== perf_processor ====")
	for _, v := range perf.perf_processor {
		fmt.Println(v.PercentProcessorUtility) //исп проца
		fmt.Println(v.ProcessorFrequency)
		fmt.Println(v.PercentProcessorPerformance)
	}
	log.Println("==== pref_disk ====")
	for _, v := range perf.pref_disk {
		fmt.Println(v.PercentDiskTime)
		fmt.Println(v.AvgDiskQueueLength)
		fmt.Println(v.DiskReadBytesPersec)
		fmt.Println(v.DiskWriteBytesPersec)
	}

}

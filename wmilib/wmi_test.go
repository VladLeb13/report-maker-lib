package wmilib

import (
	"log"
	"testing"
)

func TestGetCPU(t *testing.T) {
	GetCPU()
}

func TestGetMatherboard(t *testing.T) {
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

func TestGetRAM(t *testing.T) {
	ram := GetRAM()
	log.Println("===== ram =====")

	for _, v := range ram.ram {
		log.Println("Capacity: ", v.Capacity)
		log.Println("DeviceLocator: ", v.DeviceLocator)
		log.Println("Manufacturer: ", v.Manufacturer)
		log.Println("PartNumber: ", v.PartNumber)
	}
}

func TestGetHDD(t *testing.T) {
	hdd := GetHDD()

	log.Println("===== hdd =====")

	for _, v := range hdd.hdd {
		log.Println("InterfaceType: ", v.InterfaceType)
		log.Println("Partitions: ", v.Partitions)
		log.Println("SerialNumber: ", v.SerialNumber)
		log.Println("Size: ", v.Size)

	}
}

func TestGetVolume(t *testing.T) {
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

func TestGetNIC(t *testing.T) {
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

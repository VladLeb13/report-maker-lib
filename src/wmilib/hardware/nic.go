package hardware

import (
	"log"

	"github.com/StackExchange/wmi"
)

type NIC struct {
	Units []win32_NetworkAdapterConfiguration
}

type win32_NetworkAdapterConfiguration struct {
	Description          string
	DHCPEnabled          bool
	DHCPServer           string
	IPAddress            []string
	DefaultIPGateway     []string
	DNSServerSearchOrder []string
}

func (nic *NIC) Get() {
	queryString := wmi.CreateQuery(&nic.Units, "")
	if err := wmi.Query(queryString, &nic.Units); err != nil {
		log.Fatal(err)
	}
}

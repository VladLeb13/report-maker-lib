package hardware

import (
	"log"

	"report-maker-lib/wmilib/tools"
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
	if err := tools.ExecuteQuery(&nic.Units, ""); err != nil {
		log.Fatal(err)
	}
}

package software

import (
	"log"

	"github.com/StackExchange/wmi"
)

type (
	OS struct {
		WindowsInfo []win32_OperatingSystem
		SystemInfo  []win32_ComputerSystem
	}

	win32_OperatingSystem struct {
		Caption        string
		BuildNumber    string
		SerialNumber   string
		InstallDate    string
		OSArchitecture string
		Version        string
	}

	win32_ComputerSystem struct {
		Domain   string
		UserName string
	}
)

func (os *OS) Get() {
	queryStringSystemInfo := wmi.CreateQuery(&os.SystemInfo, "")
	if err := wmi.Query(queryStringSystemInfo, &os.SystemInfo); err != nil {
		log.Fatal(err)
	}

	queryStringWindowsInfo := wmi.CreateQuery(&os.WindowsInfo, "")
	if err := wmi.Query(queryStringWindowsInfo, &os.WindowsInfo); err != nil {
		log.Fatal(err)
	}
}

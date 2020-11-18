package software

import (
	"log"

	"report-maker-lib/wmilib/tools"
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
		Name           string
	}

	win32_ComputerSystem struct {
		Domain   string
		UserName string
	}
)

func (os *OS) Get() {
	if err := tools.ExecuteQuery(&os.SystemInfo, ""); err != nil {
		log.Fatal(err)
	}

	if err := tools.ExecuteQuery(&os.WindowsInfo, ""); err != nil {
		log.Fatal(err)
	}
}

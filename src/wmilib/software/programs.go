package software

import (
	"log"

	"github.com/VladLeb13/report-maker-lib/wmilib/tools"
)

type Programs struct {
	List []win32_Product
}

type win32_Product struct {
	Name            string
	Caption         string
	Description     string
	InstallDate     string
	InstallLocation string
	RegOwner        string
	RegCompany      string
	Vendor          string
	Version         string
}

func (prog *Programs) Get() {
	if err := tools.ExecuteQuery(&prog.List, ""); err != nil {
		log.Fatal(err)
	}
}

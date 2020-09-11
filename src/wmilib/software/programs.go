package software

import (
	"log"

	"github.com/StackExchange/wmi"
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
	queryString := wmi.CreateQuery(&prog.List, "")
	if err := wmi.Query(queryString, &prog.List); err != nil {
		log.Fatal(err)
	}
}

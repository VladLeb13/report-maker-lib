package software

import (
	"log"

	"github.com/StackExchange/wmi"
)

type (
	Updates struct {
		List []win32_QuickFixEngineering
	}

	win32_QuickFixEngineering struct {
		HotFixID    string
		Description string
		InstalledBy string
		InstalledOn string
		FixComments string
		Status      string
		Name        string
	}
)

func (updates *Updates) Get() {
	queryString := wmi.CreateQuery(&updates.List, "")
	if err := wmi.Query(queryString, &updates.List); err != nil {
		log.Fatal(err)
	}
}

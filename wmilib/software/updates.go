package software

import (
	"log"

	"report-maker-lib/wmilib/tools"
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
	if err := tools.ExecuteQuery(&updates.List, ""); err != nil {
		log.Fatal(err)
	}
}

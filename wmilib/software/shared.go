package software

import (
	"log"

	"github.com/VladLeb13/report-maker-lib/wmilib/tools"
)

type (
	Shared struct {
		Resource []win32_Share
	}

	win32_Share struct {
		Name        string
		Path        string
		Description string
		Caption     string
		Status      string
	}
)

func (shared *Shared) Get() {
	if err := tools.ExecuteQuery(&shared.Resource, ""); err != nil {
		log.Fatal(err)
	}
}

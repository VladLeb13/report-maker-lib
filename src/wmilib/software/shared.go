package software

import (
	"log"

	"github.com/StackExchange/wmi"
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
	queryString := wmi.CreateQuery(&shared.Resource, "")
	if err := wmi.Query(queryString, &shared.Resource); err != nil {
		log.Fatal(err)
	}
}

package software

import (
	"log"

	"github.com/StackExchange/wmi"
)

type Startup struct {
	Commands []win32_StartupCommand
}

type win32_StartupCommand struct {
	Caption string
	Command string
	User    string
	Name    string
}

func (startup *Startup) Get() {
	queryString := wmi.CreateQuery(&startup.Commands, "")
	if err := wmi.Query(queryString, &startup.Commands); err != nil {
		log.Fatal(err)
	}
}

package software

import (
	"log"
	"wmilib/tools"
)

type Startup struct {
	Commands []win32_StartupCommand
}

type win32_StartupCommand struct {
	Caption  string
	Command  string
	User     string
	Name     string
	Location string
}

func (startup *Startup) Get() {
	if err := tools.ExecuteQuery(&startup.Commands, ""); err != nil {
		log.Fatal(err)
	}
}

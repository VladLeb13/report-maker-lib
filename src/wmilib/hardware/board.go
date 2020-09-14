package hardware

import (
	"log"

	"github.com/VladLeb13/report-maker-lib/src/wmilib/tools"
)

type (
	Board struct {
		Bios       []win32_BIOS
		Сomponents []win32_BaseBoard
	}

	win32_BaseBoard struct {
		Manufacturer string
		Model        string
		Product      string
		PartNumber   string
		SerialNumber string
		Version      string
	}

	win32_BIOS struct {
		Manufacturer      string
		Name              string
		SMBIOSBIOSVersion string
	}
)

func (board *Board) Get() {
	if err := tools.ExecuteQuery(&board.Сomponents, ""); err != nil {
		log.Fatal(err)
	}

	if err := tools.ExecuteQuery(&board.Bios, ""); err != nil {
		log.Fatal(err)
	}
}

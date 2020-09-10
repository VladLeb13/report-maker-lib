package hardware

import (
	"log"

	"github.com/StackExchange/wmi"
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
	queryStringСomponents := wmi.CreateQuery(&board.Сomponents, "")
	if err := wmi.Query(queryStringСomponents, &board.Сomponents); err != nil {
		log.Fatal(err)
	}

	queryStringBios := wmi.CreateQuery(&board.Bios, "")
	if err := wmi.Query(queryStringBios, &board.Bios); err != nil {
		log.Fatal(err)
	}

}

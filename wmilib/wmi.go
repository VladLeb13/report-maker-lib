package wmilib

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/StackExchange/wmi"
)

/*
*	Получение
*				аппаратных
*							ресурсов
 */

//GetCPU - Получает информацию о процессоре
func GetCPU() *WMICPU {
	dst := new(WMICPU)

	qcpu := wmi.CreateQuery(&dst.processor, "")
	err := wmi.Query(qcpu, &dst.processor)
	if err != nil {
		log.Fatal(err)
	}

	return dst
}

//GetMotherboard -  Получает информацию о мат.плате
func GetMotherboard() *WMIMatherboard {
	dst := new(WMIMatherboard)

	qboard := wmi.CreateQuery(&dst.baseboard, "")
	qbios := wmi.CreateQuery(&dst.bios, "")

	if err := wmi.Query(qboard, &dst.baseboard); err != nil {
		log.Fatal(err)
	}

	if err := wmi.Query(qbios, &dst.bios); err != nil {
		log.Fatal(err)
	}

	return dst
}

//GetRAM - Получает инфу о памяти
func GetRAM() *WMIRAM {
	dst := new(WMIRAM)
	qram := wmi.CreateQuery(&dst.ram, "")
	if err := wmi.Query(qram, &dst.ram); err != nil {
		log.Fatal(err)
	}
	return dst
}

//GetHDD - Получает инфу о HDD
func GetHDD() *WMIHDD {
	dst := new(WMIHDD)
	qhdd := wmi.CreateQuery(&dst.hdd, "")
	if err := wmi.Query(qhdd, &dst.hdd); err != nil {
		log.Fatal(err)
	}
	return dst
}

//GetVolume - Информация о разделах системы
func GetVolume() *WMIVolume {
	dst := new(WMIVolume)
	qvolume := wmi.CreateQuery(&dst.volume, "")
	if err := wmi.Query(qvolume, &dst.volume); err != nil {
		log.Fatal(err)
	}
	return dst

}

//GetNIC - Информация о сетевых адаптерах
func GetNIC() *WMINIC {
	dst := new(WMINIC)
	qnic := wmi.CreateQuery(&dst.nic, "")
	if err := wmi.Query(qnic, &dst.nic); err != nil {
		log.Fatal(err)
	}
	return dst
}

/*
*	Получение
*				программных
*							ресурсов
 */

func GetOS() *WMIOS {
	dst := new(WMIOS)
	qcs := wmi.CreateQuery(&dst.cs, "")
	qos := wmi.CreateQuery(&dst.os, "")
	if err := wmi.Query(qcs, &dst.cs); err != nil {
		log.Fatal(err)
	}
	if err := wmi.Query(qos, &dst.os); err != nil {
		log.Fatal(err)
	}
	return dst
}

func GetShare() *WMIShare {
	dst := new(WMIShare)
	qshare := wmi.CreateQuery(&dst.share, "")
	if err := wmi.Query(qshare, &dst.share); err != nil {
		log.Fatal(err)
	}
	return dst
}

func GetStartup() *WMIStartup {
	dst := new(WMIStartup)
	qstartup := wmi.CreateQuery(&dst.startup, "")
	if err := wmi.Query(qstartup, &dst.startup); err != nil {
		log.Fatal(err)
	}
	return dst
}

func GetUpdate() *WMIUpdate {
	dst := new(WMIUpdate)
	qupdate := wmi.CreateQuery(&dst.update, "")
	if err := wmi.Query(qupdate, &dst.update); err != nil {
		log.Fatal(err)
	}
	return dst
}

func GetPrograms() *WMIPrograms {
	dst := new(WMIPrograms)
	qprograms := wmi.CreateQuery(&dst.programs, "")
	if err := wmi.Query(qprograms, &dst.programs); err != nil {
		log.Fatal(err)
	}
	return dst
}

/*
*	Получение
*				информации
*							о
*								событиях
*											ОС
 */

func GetEvent() *WMIEvent {
	evnt := new(WMIEvent)
	n := time.Now().UTC()

	m := int(n.Month() - 1)
	var mounth string
	if m < 10 {
		mounth = "0" + strconv.Itoa(m)
	} else {
		mounth = strconv.Itoa(m)
	}

	d := int(n.Day())
	var day string
	if d < 10 {
		day = "0" + strconv.Itoa(d)
	} else {
		day = strconv.Itoa(d)
	}

	year := strconv.Itoa(n.Year())

	ago, err := time.Parse("2006-01-02T15:04:05Z", year+"-"+mounth+"-"+day+"T15:04:05Z")
	if err != nil {
		log.Fatal(err)
	}

	layout := "20060102150405.898655-000"
	now := n.Format(layout)
	qevent := wmi.CreateQuery(&evnt.events, "WHERE TimeWritten >= '"+ago.Format(layout)+"' AND TimeWritten <= '"+now+"'")
	fmt.Println(qevent)
	if err := wmi.Query(qevent, &evnt.events); err != nil {
		log.Fatal(err)
	}
	return evnt
}

/*
* 	Получение
*				информации
*             				о
*                				производительности
 */

//GetPerfomance - информация о производительности
func GetPerfomance() *WMIPerfomance {

	perfom := new(WMIPerfomance)
	qpmem := wmi.CreateQuery(&perfom.perf_mem, "")
	qpcpu := wmi.CreateQuery(&perfom.perf_processor, "")
	qpdisk := wmi.CreateQuery(&perfom.pref_disk, "")
	qpproccess := wmi.CreateQuery(&perfom.perf_process, "")

	if err := wmi.Query(qpmem, &perfom.perf_mem); err != nil {
		log.Fatal(err)
	}
	if err := wmi.Query(qpcpu, &perfom.perf_processor); err != nil {
		log.Fatal(err)
	}
	if err := wmi.Query(qpdisk, &perfom.pref_disk); err != nil {
		log.Fatal(err)
	}
	if err := wmi.Query(qpproccess, &perfom.perf_process); err != nil {
		log.Fatal(err)
	}
	return perfom
}

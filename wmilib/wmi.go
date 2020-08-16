package wmilib

import (
	"log"

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

func GetOS() {

}

func GetShare() {

}

func GetStartup() {

}

func GetUpdate() {

}

func GetPrograms() {

}

/*
*	Получение
*				информации
*							о
*								событиях
*											ОС
 */

func GetEvent() {

}

/*
* 	Получение
*				информации
*             				о
*                				производительности
 */

func GetPerfomance() {

}

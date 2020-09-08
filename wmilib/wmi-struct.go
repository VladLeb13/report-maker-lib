package wmilib

/*
*  Информация
*             об
*                аппаратных
*                           ресурсах
 */

//WMICPU - Данные о процессоре
type WMICPU struct {
	processor []Win32_Processor
}

type Win32_Processor struct {
	Manufacturer  string
	Name          string
	MaxClockSpeed uint32
	NumberOfCores uint32
	ThreadCount   uint32
}

//WMIMatherboard - Данные о материнской плате
type WMIMatherboard struct {
	bios      []Win32_BIOS
	baseboard []Win32_BaseBoard
}

type Win32_BaseBoard struct {
	Manufacturer string
	Model        string
	Product      string
	PartNumber   string
	SerialNumber string
	Version      string
}

type Win32_BIOS struct {
	Manufacturer      string
	Name              string
	SMBIOSBIOSVersion string
}

//WMIRAM - Данные о ОЗУ
type WMIRAM struct {
	ram []Win32_PhysicalMemory
}

type Win32_PhysicalMemory struct {
	Capacity      uint64
	Speed         uint32
	DeviceLocator string
	PartNumber    string
	Manufacturer  string
}

//WMIHDD - Данные о HDD
type WMIHDD struct {
	hdd []Win32_DiskDrive
}

type Win32_DiskDrive struct {
	SerialNumber  string
	InterfaceType string
	Size          uint64
	Partitions    uint32
}

//WMIVolume - данные о разделах
type WMIVolume struct {
	volume []Win32_Volume
}

type Win32_Volume struct {
	Name       string
	Capacity   uint32
	FreeSpace  uint32
	Label      string
	FileSystem string
	DeviceID   string
	Automount  bool
}

//WMINIC - Данные о сетевых адаптерах
type WMINIC struct {
	nic []Win32_NetworkAdapterConfiguration
}

type Win32_NetworkAdapterConfiguration struct {
	Description          string
	DHCPEnabled          bool
	DHCPServer           string
	IPAddress            []string
	DefaultIPGateway     []string
	DNSServerSearchOrder []string
}

/*
*  Информация
*             об
*                программных
*                           ресурсах
 */

//Информация о ОС
type WMIOS struct {
	os []Win32_OperatingSystem
	cs []Win32_ComputerSystem
}

type Win32_OperatingSystem struct {
	Caption        string
	BuildNumber    string
	SerialNumber   string
	InstallDate    string
	OSArchitecture string
	Version        string
}

type Win32_ComputerSystem struct {
	Domain   string
	UserName string
}

//Информация о общих ресурсах
type WMIShare struct {
	share []Win32_Share
}

type Win32_Share struct {
	Name        string
	Path        string
	Description string
	Caption     string
	Status      string
}

//Информация о автозагрузке
type WMIStartup struct {
	startup []Win32_StartupCommand
}

type Win32_StartupCommand struct {
	Caption string
	Command string
	User    string
}

//Информация о обновлениях
type WMIUpdate struct {
	update []Win32_QuickFixEngineering
}

type Win32_QuickFixEngineering struct {
	HotFixID    string
	Description string
	InstalledBy string
	InstalledOn string
	FixComments string
	Status      string
	Name        string
}

//Информация о программах
type WMIPrograms struct {
	programs []Win32_Product
}

type Win32_Product struct {
	Name            string
	Caption         string
	Description     string
	InstallDate     string
	InstallLocation string
	RegOwner        string
	RegCompany      string
	Vendor          string
	Version         string
}

/*
*  Информация
*             об
*                событиях
*                          ос
 */

//WMIEvent - События ОС
type WMIEvent struct {
	events []Win32_NTLogEvent
}

type Win32_NTLogEvent struct {
	User           string
	LogFile        string
	Message        string
	CategoryString string
	ComputerName   string
	Data           []uint8
	SourceName     string
	TimeWritten    string
}

/*
*  Информация
*             об
*                производительности
 */

//WMIPerfomance -  Данные о производительности
type WMIPerfomance struct {
	pref_disk      []Win32_PerfFormattedData_PerfDisk_PhysicalDisk
	perf_processor []Win32_PerfFormattedData_Counters_ProcessorInformation
	perf_mem       []Win32_PerfFormattedData_PerfOS_Memory
	perf_process   []Win32_PerfFormattedData_PerfProc_Process
}
type Win32_PerfFormattedData_PerfDisk_PhysicalDisk struct {
	PercentDiskTime       uint64
	DiskWriteBytesPersec  uint64
	DiskReadBytesPersec   uint64
	AvgDisksecPerTransfer uint32
	AvgDiskQueueLength    uint64
}

type Win32_PerfFormattedData_Counters_ProcessorInformation struct {
	PercentProcessorPerformance uint64
	PercentProcessorUtility     uint64
	ProcessorFrequency          uint64
}

type Win32_PerfFormattedData_PerfOS_Memory struct {
	PercentCommittedBytesInUse uint64
	AvailableMBytes            uint64
}

type Win32_PerfFormattedData_PerfProc_Process struct {
	IOReadOperationsPersec  uint64
	IOWriteOperationsPersec uint64
	Name                    string
	PercentProcessorTime    uint64
	ThreadCount             uint32
	VirtualBytesPeak        uint64
	IDProcess               uint32
}

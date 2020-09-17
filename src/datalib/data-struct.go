package datalib

type (
	Hardware struct {
		CPU     []CPU    `json:"cpu_list"`
		Board   Board    `json:"board"`
		HDDs    []HDD    `json:"hdd_list"`
		Volumes []Volume `json:"volumes_list"`
		NICs    []NIC    `json:"nic_list"`
		RAMs    []RAM    `json:"ram_list"`
	}

	CPU struct {
		Manufacturer  string `json:"manufacturer"`
		Name          string `json:"name"`
		MaxClockSpeed uint32 `json:"speed"`
		NumberOfCores uint32 `json:"numberCores"`
		ThreadCount   uint32 `json:"threadCount"`
		Caption       string `json:"caption"`
		Description   string `json:"description"`
	}

	Board struct {
		Manufacturer string `json:"manufacturer"`
		Model        string `json:"model"`
		Product      string `json:"product"`
		PartNumber   string `json:"partNumber"`
		SerialNumber string `json:"serialNumber"`
		Version      string `json:"version"`
		BIOS         BIOS   `json:"bios"`
	}
	BIOS struct {
		Manufacturer string `json:"manufacturer"`
		Name         string `json:"name"`
		Version      string `json:"version"`
	}
	HDD struct {
		SerialNumber  string `json:"serialNumber"`
		InterfaceType string `json:"interfaceType"`
		Size          uint64 `json:"size"`
		Partitions    uint32 `json:"partitions"`
		Model         string `json:"model"`
		Name          string `json:"name"`
	}
	Volume struct {
		Caption     string `json:"caption"`
		Description string `json:"description"`
		DeviceID    string `json:"deviceID"`
		FileSystem  string `json:"fileSystem"`
		FreeSpace   uint64 `json:"freeSpace"`
		Name        string `json:"name"`
		Size        string `json:"size"`
		VolumeName  string `json:"volumeName"`
	}
	NIC struct {
		Description          string   `json:"description"`
		DHCPEnabled          bool     `json:"dhcpEnabled"`
		DHCPServer           string   `json:"dhcpServer"`
		IPAddress            []string `json:"ipAddress_list"`
		DefaultIPGateway     []string `json:"defaultIPGateway_list"`
		DNSServerSearchOrder []string `json:"dnsServerSearchOrderlist"`
	}
	RAM struct {
		Capacity             uint64 `json:"capacity"`
		Speed                uint32 `json:"speed"`
		DeviceLocator        string `json:"deviceLocator"`
		PartNumber           string `json:"partNumber"`
		Manufacturer         string `json:"manufacturer"`
		Model                string `json:"model"`
		Name                 string `json:"name"`
		FormFactor           uint16 `json:"formFactor"`
		OtherIdentifyingInfo string `json:"otherIdentifyingInfo"`
	}
)

type (
	Software struct {
		OS       OS        `json:"os"`
		Programs []Program `json:"programs_list"`
		Updates  []Update  `json:"updates_list"`
		Startup  []Startup `json:"startup_list"`
		Shared   []Shared  `json:"shared_list"`
	}

	OS struct {
		Caption        string `json:"caption"`
		BuildNumber    string `json:"buildNumber"`
		SerialNumber   string `json:"serialNumber"`
		InstallDate    string `json:"installDate"`
		OSArchitecture string `json:"osArchitecture"`
		Version        string `json:"version"`
		Name           string `json:"name"`
		Domain         string `json:"domain"`
		UserName       string `json:"userName"`
	}
	Program struct {
		Name            string `json:"name"`
		Caption         string `json:"caption"`
		Description     string `json:"description"`
		InstallDate     string `json:"installDate"`
		InstallLocation string `json:"installLocation"`
		RegOwner        string `json:"regOwner"`
		RegCompany      string `json:"regCompany"`
		Vendor          string `json:"vendor"`
		Version         string `json:"version"`
	}
	Update struct {
		HotFixID    string `json:"hotFixID"`
		Description string `json:"description"`
		InstalledBy string `json:"installedBy"`
		InstalledOn string `json:"installedOn"`
		FixComments string `json:"fixComments"`
		Status      string `json:"status"`
		Name        string `json:"name"`
	}
	Startup struct {
		Caption  string `json:"caption"`
		Command  string `json:"command"`
		User     string `json:"user"`
		Name     string `json:"name"`
		Location string `json:"location"`
	}
	Shared struct {
		Name        string `json:"name"`
		Path        string `json:"path"`
		Description string `json:"description"`
		Caption     string `json:"caption"`
		Status      string `json:"status"`
	}
)

type (
	Perfomance struct {
		CPU       []Perf_cpu     `json:"cpu"`
		Processes []Perf_process `json:"processes"`
		HDD       []Perf_hdd     `json:"hdd"`
		RAM       []Perf_ram     `json:"ram"`
	}
	Perf_cpu struct {
		PercentProcessorPerformance uint64 `json:"percentProcessorPerformance"`
		PercentProcessorUtility     uint64 `json:"percentProcessorUtility"`
		ProcessorFrequency          uint64 `json:"processorFrequency"`
	}
	Perf_process struct {
		IOReadOperationsPersec  uint64 `json:"ioReadOperationsPersec"`
		IOWriteOperationsPersec uint64 `json:"ioWriteOperationsPersec"`
		Name                    string `json:"name"`
		PercentProcessorTime    uint64 `json:"percentProcessorTime"`
		ThreadCount             uint32 `json:"threadCount"`
		VirtualBytesPeak        uint64 `json:"virtualBytesPeak"`
		IDProcess               uint32 `json:"idProcess"`
	}
	Perf_hdd struct {
		PercentDiskTime       uint64 `json:"percentDiskTime"`
		DiskWriteBytesPersec  uint64 `json:"diskWriteBytesPersec"`
		DiskReadBytesPersec   uint64 `json:"diskReadBytesPersec"`
		AvgDisksecPerTransfer uint32 `json:"avgDisksecPerTransfer"`
		AvgDiskQueueLength    uint64 `json:"avgDiskQueueLength"`
	}
	Perf_ram struct {
		PercentCommittedBytesInUse uint64 `json:"percentCommittedBytesInUse"`
		AvailableMBytes            uint64 `json:"availableMBytes"`
	}
)

type (
	Events struct {
		List []Event `json:"events"`
	}
	Event struct {
		User           string  `json:"user"`
		LogFile        string  `json:"logFile"`
		Message        string  `json:"message"`
		CategoryString string  `json:"categoryString"`
		ComputerName   string  `json:"computerName"`
		Data           []uint8 `json:"data"`
		SourceName     string  `json:"sourceName"`
		TimeWritten    string  `json:"timeWritten"`
	}
)

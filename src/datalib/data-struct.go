package datalib

type (
	Hardware struct {
		CPU     []cpu    `json:"cpu_list"`
		Board   board    `json:"board"`
		HDDs    []hdd    `json:"hdd_list"`
		Volumes []volume `json:"volumes_list"`
		NICs    []nic    `json:"nic_list"`
		RAMs    []ram    `json:"ram_list"`
	}

	cpu struct {
		Manufacturer  string `json:"manufacturer"`
		Name          string `json:"name"`
		MaxClockSpeed uint32 `json:"speed"`
		NumberOfCores uint32 `json:"numberCores"`
		ThreadCount   uint32 `json:"threadCount"`
		Caption       string `json:"caption"`
		Description   string `json:"description"`
	}

	board struct {
		Manufacturer string `json:"manufacturer"`
		Model        string `json:"model"`
		Product      string `json:"product"`
		PartNumber   string `json:"partNumber"`
		SerialNumber string `json:"serialNumber"`
		Version      string `json:"version"`
		BIOS         bios   `json:"bios"`
	}
	bios struct {
		Manufacturer string `json:"manufacturer"`
		Name         string `json:"name"`
		Version      string `json:"version"`
	}
	hdd struct {
		SerialNumber  string `json:"serialNumber"`
		InterfaceType string `json:"interfaceType"`
		Size          uint64 `json:"size"`
		Partitions    uint32 `json:"partitions"`
		Model         string `json:"model"`
		Name          string `json:"name"`
	}
	volume struct {
		Caption     string `json:"caption"`
		Description string `json:"description"`
		DeviceID    string `json:"deviceID"`
		FileSystem  string `json:"fileSystem"`
		FreeSpace   uint64 `json:"freeSpace"`
		Name        string `json:"name"`
		Size        string `json:"size"`
		VolumeName  string `json:"volumeName"`
	}
	nic struct {
		Description          string   `json:"description"`
		DHCPEnabled          bool     `json:"dhcpEnabled"`
		DHCPServer           string   `json:"dhcpServer"`
		IPAddress            []string `json:"ipAddress_list"`
		DefaultIPGateway     []string `json:"defaultIPGateway_list"`
		DNSServerSearchOrder []string `json:"dnsServerSearchOrderlist"`
	}
	ram struct {
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
		OS       os        `json:"os"`
		Programs []program `json:"programs_list"`
		Updates  []update  `json:"updates_list"`
		Startup  []startup `json:"startup_list"`
		Shared   []shared  `json:"shared_list"`
	}

	os struct {
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
	program struct {
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
	update struct {
		HotFixID    string `json:"hotFixID"`
		Description string `json:"description"`
		InstalledBy string `json:"installedBy"`
		InstalledOn string `json:"installedOn"`
		FixComments string `json:"fixComments"`
		Status      string `json:"status"`
		Name        string `json:"name"`
	}
	startup struct {
		Caption  string `json:"caption"`
		Command  string `json:"command"`
		User     string `json:"user"`
		Name     string `json:"name"`
		Location string `json:"location"`
	}
	shared struct {
		Name        string `json:"name"`
		Path        string `json:"path"`
		Description string `json:"description"`
		Caption     string `json:"caption"`
		Status      string `json:"status"`
	}
)

type (
	Perfomance struct {
		CPU       []perf_cpu     `json:"cpu"`
		Processes []perf_process `json:"processes"`
		HDD       []perf_hdd     `json:"hdd"`
		RAM       []perf_ram     `json:"ram"`
	}
	perf_cpu struct {
		PercentProcessorPerformance uint64 `json:"percentProcessorPerformance"`
		PercentProcessorUtility     uint64 `json:"percentProcessorUtility"`
		ProcessorFrequency          uint64 `json:"processorFrequency"`
	}
	perf_process struct {
		IOReadOperationsPersec  uint64 `json:"ioReadOperationsPersec"`
		IOWriteOperationsPersec uint64 `json:"ioWriteOperationsPersec"`
		Name                    string `json:"name"`
		PercentProcessorTime    uint64 `json:"percentProcessorTime"`
		ThreadCount             uint32 `json:"threadCount"`
		VirtualBytesPeak        uint64 `json:"virtualBytesPeak"`
		IDProcess               uint32 `json:"idProcess"`
	}
	perf_hdd struct {
		PercentDiskTime       uint64 `json:"percentDiskTime"`
		DiskWriteBytesPersec  uint64 `json:"diskWriteBytesPersec"`
		DiskReadBytesPersec   uint64 `json:"diskReadBytesPersec"`
		AvgDisksecPerTransfer uint32 `json:"avgDisksecPerTransfer"`
		AvgDiskQueueLength    uint64 `json:"avgDiskQueueLength"`
	}
	perf_ram struct {
		PercentCommittedBytesInUse uint64 `json:"percentCommittedBytesInUse"`
		AvailableMBytes            uint64 `json:"availableMBytes"`
	}
)

type (
	Events struct {
		List []event `json:"events"`
	}
	event struct {
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

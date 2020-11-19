package datalib

type (
	Report struct {
		Hardware   Hardware   `json:"hardware"`
		Software   Software   `json:"software"`
		Perfomance Perfomance `json:"perfomance"`
		Events     Events     `json:"evens"`
	}
	Hardware struct {
		CPUs    []CPU    `json:"cpu_list"`
		Board   Board    `json:"board"`
		HDDs    []HDD    `json:"hdd_list"`
		Volumes []Volume `json:"volumes_list"`
		NICs    []NIC    `json:"nic_list"`
		RAMs    []RAM    `json:"ram_list"`
	}

	CPU struct {
		Manufacturer  string `json:"manufacturer" name:"Производитель"`
		Name          string `json:"name" name:"Наименование"`
		MaxClockSpeed uint32 `json:"speed" name:"Максимальная частота"`
		NumberOfCores uint32 `json:"numberCores" name:"Количество ядер"`
		ThreadCount   uint32 `json:"threadCount" name:"Количество потоков"`
		Caption       string `json:"caption" name:"Название"`
		Description   string `json:"description" name:"Описание"`
	}

	Board struct {
		Manufacturer string `json:"manufacturer" name:"Производитель"`
		Product      string `json:"product" name:"Название продукта"`
		SerialNumber string `json:"serialNumber" name:"Серийный номер"`
		Version      string `json:"version" name:"Версия"`
		BIOS         BIOS   `json:"bios"`
	}
	BIOS struct {
		Manufacturer string `json:"manufacturer" name:"Производитель BIOS"`
		Name         string `json:"name" name:"Наименование BIOS"`
		Version      string `json:"version" name:"Версия BIOS"`
	}
	HDD struct {
		SerialNumber  string `json:"serialNumber" name:"Серийный номер"`
		InterfaceType string `json:"interfaceType" name:"Тип интерфейса"`
		Size          uint64 `json:"size" name:"Объем"`
		Partitions    uint32 `json:"partitions" name:"Количество разделов"`
		Model         string `json:"model" name:"Модель"`
		Name          string `json:"name" name:"Название"`
	}
	Volume struct {
		Caption     string `json:"caption" name:"Метка тома"`
		Description string `json:"description" name:"Описание"`
		DeviceID    string `json:"deviceID" name:"ID"`
		FileSystem  string `json:"fileSystem" name:"Файловая система"`
		FreeSpace   uint64 `json:"freeSpace" name:"Объем свободного пространства"`
		Name        string `json:"name" name:"Название"`
		Size        string `json:"size" name:"Объем"`
		VolumeName  string `json:"volumeName" name:"Название тома"`
	}
	NIC struct {
		Description          string   `json:"description" name:"Описание"`
		DHCPEnabled          bool     `json:"dhcpEnabled" name:"Активность DHCP"`
		DHCPServer           string   `json:"dhcpServer" name:"DHCP сервер"`
		IPAddress            []string `json:"ipAddress_list" name:"IP"`
		DefaultIPGateway     []string `json:"defaultIPGateway_list" name:"Шлюз"`
		DNSServerSearchOrder []string `json:"dnsServerSearchOrderlist" name:"DNS"`
	}
	RAM struct {
		Capacity      uint64 `json:"capacity" name:"Объем"`
		Speed         uint32 `json:"speed" name:"Скорость"`
		DeviceLocator string `json:"deviceLocator" name:"Шина"`
		PartNumber    string `json:"partNumber" name:"Номер"`
		Manufacturer  string `json:"manufacturer" name:"Производитель"`
		Model         string `json:"model" name:"Модель"`
		Name          string `json:"name" name:"Наименование"`
		FormFactor    uint16 `json:"formFactor" name:"Форм фактор"`
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
		Caption        string `json:"caption" name:"Наименование"`
		BuildNumber    string `json:"buildNumber" name:"Номер сборки"`
		SerialNumber   string `json:"serialNumber" name:"Серийный номер"`
		InstallDate    string `json:"installDate" name:"Дата установки"`
		OSArchitecture string `json:"osArchitecture" name:"Архитектура ОС"`
		Version        string `json:"version" name:"Версия"`
		Name           string `json:"name" name:"Название"`
		Domain         string `json:"domain" name:"Домен"`
		UserName       string `json:"userName" name:"Пользователь"`
	}
	Program struct {
		Name            string `json:"name" name:"Наименование"`
		Description     string `json:"description" name:"Описание"`
		InstallDate     string `json:"installDate" name:"Дата установки"`
		InstallLocation string `json:"installLocation" name:"Каталог установки"`
		Vendor          string `json:"vendor" name:"Производитель"`
		Version         string `json:"version" name:"Версия"`
	}
	Update struct {
		HotFixID    string `json:"hotFixID" name:"ID"`
		Description string `json:"description" name:"Описание"`
		InstalledBy string `json:"installedBy" name:"Пользователь"`
		InstalledOn string `json:"installedOn" name:"Дата установки"`
	}
	Startup struct {
		Caption  string `json:"caption" name:"Название"`
		Command  string `json:"command" name:"Команда"`
		User     string `json:"user" name:"Пользователь"`
		Name     string `json:"name" name:"Наименование"`
		Location string `json:"location" name:"Расположение"`
	}
	Shared struct {
		Name        string `json:"name" name:"Имя"`
		Path        string `json:"path" name:"Путь"`
		Description string `json:"description" name:"Описание"`
		Caption     string `json:"caption" name:"Наименование"`
		Status      string `json:"status" name:"Статус"`
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
		PercentProcessorPerformance uint64 `json:"percentProcessorPerformance" name:"Процент производительости процессора"`
		PercentProcessorUtility     uint64 `json:"percentProcessorUtility" name:"Использовано процентов"`
		ProcessorFrequency          uint64 `json:"processorFrequency" name:"Частота"`
	}
	Perf_process struct {
		IOReadOperationsPersec  uint64 `json:"ioReadOperationsPersec" name:"Операций стени в сек"`
		IOWriteOperationsPersec uint64 `json:"ioWriteOperationsPersec" name:"Операций записи в сек"`
		Name                    string `json:"name" name:"Название"`
		PercentProcessorTime    uint64 `json:"percentProcessorTime" name:"испольование процессора"`
		ThreadCount             uint32 `json:"threadCount" name:"использовано потоков"`
		VirtualBytesPeak        uint64 `json:"virtualBytesPeak" name:"занято памят"`
		IDProcess               uint32 `json:"idProcess" name:"ID"`
	}
	Perf_hdd struct {
		PercentDiskTime       uint64 `json:"percentDiskTime" name:"Процент испьзования диска"`
		DiskWriteBytesPersec  uint64 `json:"diskWriteBytesPersec" name:"Записано в сек"`
		DiskReadBytesPersec   uint64 `json:"diskReadBytesPersec" name:"Прочитано в сек"`
		AvgDisksecPerTransfer uint32 `json:"avgDisksecPerTransfer" name:"Пеедано в сек"`
		AvgDiskQueueLength    uint64 `json:"avgDiskQueueLength" name:"Очередь"`
	}
	Perf_ram struct {
		PercentCommittedBytesInUse uint64 `json:"percentCommittedBytesInUse" name:"Задействовано процентов"`
		AvailableMBytes            uint64 `json:"availableMBytes" name:"своодно мегбайт"`
	}
)

type (
	Events struct {
		List []Event `json:"events"`
	}
	Event struct {
		User        string  `json:"user" name:"Пользователь"`
		LogFile     string  `json:"logFile" name:"Файл лога"`
		Message     string  `json:"message" name:"Сообщение"`
		Data        []uint8 `json:"data" name:"Данные"`
		TimeWritten string  `json:"timeWritten" name:"Время записи"`
	}
)

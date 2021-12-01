package hwparser

type HardwareInfo struct {
	BIOS                     BIOS                      `json:"bios"`
	System                   System                    `json:"system"`
	MainMemory               MainMemory                `json:"mainMemory"`
	FPUS                     []FPU                     `json:"fpus"`
	DMAControllers           []DMAController           `json:"dmaControllers"`
	PICS                     []PIC                     `json:"pics"`
	KeyboardControllers      []KeyboardController      `json:"keyboardControllers"`
	WLANControllers          []WLANController          `json:"wlanControllers"`
	SATAControllers          []SATAController          `json:"sataControllers"`
	PCIBridges               []PCIBridge               `json:"pciBridges"`
	ISABridges               []ISABridge               `json:"isaBridges"`
	CommunicationControllers []CommunicationController `json:"communicationControllers"`
	AudioDevices             []AudioDevice             `json:"audioDevices"`
	EthernetControllers      []EthernetController      `json:"ethernetControllers"`
	SMBuses                  []SMBus                   `json:"smBuses"`
	HostBridges              []HostBridge              `json:"hostBridges"`
	USBControllers           []USBController           `json:"usbControllers"`
	SerialControllers        []SerialController        `json:"serialControllers"`
	SystemPeripherals        []SystemPeripheral        `json:"systemPeripherals"`
	VGACompatibleControllers []VGACompatibleController `json:"vgaCompatibleControllers"`
	LCDMonitors              []LCDMonitor              `json:"lcdMonitors"`
	Disks                    []Disk                    `json:"disks"`
	Partitions               []Partition               `json:"partitions"`
	BluetoothDevices         []BluetoothDevice         `json:"bluetoothDevices"`
	UnclassifiedDevices      []UnclassifiedDevice      `json:"unclassifiedDevices"`
	Hubs                     []Hub                     `json:"hubs"`
	Keyboards                []Keyboard                `json:"keyboards"`
	USBMouses                []USBMouse                `json:"usbMouses"`
	CPUS                     []CPU                     `json:"cpus"`
	Loopbacks                []Loopback                `json:"loopbacks"`
	Ethernets                []Ethernet                `json:"ethernets"`
}

type BIOS struct {
	UniqueID      string `json:"uniqueId"`
	HardwareClass string `json:"hardwareClass"`
	ConfigStatus  string `json:"configStatus"`
}

type System struct {
	UniqueID      string `json:"uniqueId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	Formfactor    string `json:"formFactor"`
	ConfigStatus  string `json:"configStatus"`
}

type FPU struct {
	UniqueID      string   `json:"uniqueId"`
	HardwareClass string   `json:"hardwareClass"`
	Model         string   `json:"model"`
	IOPort        []string `json:"ioports"`
	ConfigStatus  string   `json:"configStatus"`
}

type DMAController struct {
	UniqueID      string   `json:"uniqueId"`
	HardwareClass string   `json:"hardwareClass"`
	Model         string   `json:"model"`
	IOPort        []string `json:"ioports"`
	DMA           int64    `json:"dma"`
	ConfigStatus  string   `json:"configStatus"`
}

type PIC struct {
	UniqueID      string   `json:"uniqueId"`
	HardwareClass string   `json:"hardwareClass"`
	Model         string   `json:"model"`
	IOPort        []string `json:"ioports"`
	ConfigStatus  string   `json:"configStatus"`
}

type KeyboardController struct {
	UniqueID      string   `json:"uniqueId"`
	HardwareClass string   `json:"hardwareClass"`
	Model         string   `json:"model"`
	IOPort        []string `json:"ioports"`
	ConfigStatus  string   `json:"configStatus"`
}

type MainMemory struct {
	UniqueID      string `json:"uniqueId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	MemoryRange   string `json:"memoryRange"`
	MemorySize    string `json:"memorySize"`
	ConfigStatus  string `json:"configStatus"`
}

type WLANController struct {
	UniqueID                string `json:"uniqueId"`
	ParentID                string `json:"parentId"`
	SysFSID                 string `json:"sysFsId"`
	SysFSBusID              string `json:"sysFsBusId"`
	HardwareClass           string `json:"hardwareClass"`
	Model                   string `json:"model"`
	Vendor                  string `json:"vendor"`
	Device                  string `json:"device"`
	SubVendor               string `json:"subVendor"`
	SubDevice               string `json:"subDevice"`
	Revision                string `json:"revision"`
	Driver                  string `json:"driver"`
	DriverModules           string `json:"driverModules"`
	DeviceFile              string `json:"deviceFile"`
	Features                string `json:"features"`
	MemoryRange             string `json:"memoryRange"`
	IRQ                     string `json:"irq"`
	HWAddress               string `json:"hwAddress"`
	PermanentHWAddress      string `json:"permanentHwAddress"`
	LinkDetected            string `json:"linkDetected"`
	WLANChannels            string `json:"wlanChannels"`
	WLANFrequences          string `json:"wlanFrequences"`
	WLANEncryptionModes     string `json:"wlanEncryptionModes"`
	WLANAuthenticationModes string `json:"wlanAuthenticationModes"`
	ModuleAlias             string `json:"moduleAlias"`
	ConfigStatus            string `json:"configStatus"`
	AttachedTo              string `json:"attachedTo"`
}

type SATAController struct {
	UniqueID      string   `json:"uniqueId"`
	SysFSID       string   `json:"sysFsId"`
	SysFSBusID    string   `json:"sysFsBusId"`
	HardwareClass string   `json:"hardwareClass"`
	Model         string   `json:"model"`
	Vendor        string   `json:"vendor"`
	Device        string   `json:"device"`
	SubVendor     string   `json:"subVendor"`
	SubDevice     string   `json:"subDevice"`
	Revision      string   `json:"revision"`
	Driver        string   `json:"driver"`
	DriverModules string   `json:"driverModules"`
	MemoryRange   []string `json:"memoryRanges"`
	IOPort        []string `json:"ioPorts"`
	IRQ           string   `json:"irq"`
	ModuleAlias   string   `json:"moduleAlias"`
	ConfigStatus  string   `json:'configStatus`
}

type MemoryController struct {
	UniqueID      string   `json:"uniqueId"`
	SysFSID       string   `json:"sysFsId"`
	SysFSBusID    string   `json:"sysFsBusId"`
	HardwareClass string   `json:"hardwareClass"`
	Model         string   `json:"model"`
	Vendor        string   `json:"vendor"`
	Device        string   `json:"device"`
	SubVendor     string   `json:"subVendor"`
	SubDevice     string   `json:"subDevice"`
	Revision      string   `json:"revision"`
	MemoryRange   []string `json:"memoryRanges"`
	ModuleAlias   string   `json:"moduleAlias"`
	ConfigStatus  string   `json:"configStatus"`
}

type Bridge struct {
	UniqueID      string `json:"uniqueId"`
	SysFSID       string `json:"sysFsId"`
	SysFSBusID    string `json:"sysFsBusId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	Vendor        string `json:"vendor"`
	Device        string `json:"device"`
	SubVendor     string `json:"subVendor"`
	SubDevice     string `json:"subDevice"`
	Revision      string `json:"revision"`
	Driver        string `json:"driver"`
	IRQ           string `json:"irq"`
	ModuleAlias   string `json:"moduleAlias"`
	ConfigStatus  string `json:"configStatus"`
}

type ISABridge struct {
	UniqueID      string `json:"uniqueId"`
	SysFSID       string `json:"sysFsId"`
	SysFSBusID    string `json:"sysFsBusId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	Vendor        string `json:"vendor"`
	Device        string `json:"device"`
	SubVendor     string `json:"subVendor"`
	SubDevice     string `json:"subDevice"`
	Revision      string `json:"revision"`
	Driver        string `json:"driver"`
	DriverModules string `json:"driverModules"`
	ModuleAlias   string `json:"moduleAlias"`
	ConfigStatus  string `json:"configStatus"`
}

type CommunicationController struct {
	UniqueID      string   `json:"uniqueId"`
	SysFSID       string   `json:"sysFsId"`
	SysFSBusID    string   `json:"sysFsBusId"`
	HardwareClass string   `json:"hardwareClass"`
	Model         string   `json:"model"`
	Vendor        string   `json:"vendor"`
	Device        string   `json:"device"`
	SubVendor     string   `json:"subVendor"`
	SubDevice     string   `json:"subDevice"`
	Revision      string   `json:"revision"`
	Driver        string   `json:"driver"`
	DriverModules string   `json:"driverModules"`
	MemoryRange   []string `json:"memoryRanges"`
	IRQ           string   `json:"irq"`
	ModuleAlias   string   `json:"moduleAlias"`
	ConfigStatus  string   `json:"configStatus"`
}

type AudioDevice struct {
	UniqueID      string   `json:"uniqueId"`
	SysFSID       string   `json:"sysFsId"`
	SysFSBusID    string   `json:"sysFsBusId"`
	HardwareClass string   `json:"hardwareClass"`
	Model         string   `json:"model"`
	Vendor        string   `json:"vendor"`
	Device        string   `json:"device"`
	SubVendor     string   `json:"subVendor"`
	SubDevice     string   `json:"subDevice"`
	Revision      string   `json:"revision"`
	Driver        string   `json:"driver"`
	DriverModules string   `json:"driverModules"`
	MemoryRange   []string `json:"memoryRanges"`
	IRQ           string   `json:"irq"`
	ModuleAlias   string   `json:"moduleAlias"`
	ConfigStatus  string   `json:"configStatus"`
}

type PCIBridge struct {
	UniqueID      string `json:"uniqueId"`
	SysFSID       string `json:"sysFsId"`
	SysFSBusID    string `json:"sysFsBusId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	Vendor        string `json:"vendor"`
	Device        string `json:"device"`
	SubVendor     string `json:"subVendor"`
	SubDevice     string `json:"subDevice"`
	Revision      string `json:"revision"`
	Driver        string `json:"driver"`
	IRQ           string `json:"irq"`
	ModuleAlias   string `json:"moduleAlias"`
	ConfigStatus  string `json:"configStatus"`
}

type EthernetController struct {
	UniqueID           string   `json:"uniqueId"`
	SysFSID            string   `json:"sysFsId"`
	SysFSBusID         string   `json:"sysFsBusId"`
	HardwareClass      string   `json:"hardwareClass"`
	Model              string   `json:"model"`
	Vendor             string   `json:"vendor"`
	Device             string   `json:"device"`
	SubVendor          string   `json:"subVendor"`
	SubDevice          string   `json:"subDevice"`
	Revision           string   `json:"revision"`
	Driver             string   `json:"driver"`
	DriverModules      string   `json:"driverModules"`
	DeviceFile         string   `json:"deviceFile"`
	MemoryRange        []string `json:"memoryRanges"`
	IRQ                string   `json:"irq"`
	HWAddress          string   `json:"hwAddress"`
	PermanentHWAddress string   `json:"permanentHwAddress"`
	LinkDetected       string   `json:"linkDetected"`
	ModuleAlias        string   `json:"moduleAlias"`
	ConfigStatus       string   `json:"configStatus"`
}

type SMBus struct {
	UniqueID      string   `json:"smBus"`
	SysFSID       string   `json:"sysFsId"`
	SysFSBusID    string   `json:"sysFsBusId"`
	HardwareClass string   `json:"hardwareClass"`
	Model         string   `json:"model"`
	Vendor        string   `json:"vendor"`
	Device        string   `json:"device"`
	SubVendor     string   `json:"subVendor"`
	SubDevice     string   `json:"subDevice"`
	Revision      string   `json:"revision"`
	Driver        string   `json:"driver"`
	DriverModules string   `json:"driverModules"`
	MemoryRange   []string `json:"memoryRanges"`
	IRQ           string   `json:"irq"`
	ModuleAlias   string   `json:"moduleAlias"`
	ConfigStatus  string   `json:"configStatus"`
}

type HostBridge struct {
	UniqueID      string `json:"uniqueId"`
	SysFSID       string `json:"sysFsId"`
	SysFSBusID    string `json"sysFsBusId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	Vendor        string `json:"vendor"`
	Device        string `json:"device"`
	SubVendor     string `json:"subVendor"`
	SubDevice     string `json:"subDevice"`
	Revision      string `json:"revision"`
	ModuleAlias   string `json:"moduleaLias"`
	ConfigStatus  string `json:"configStatus"`
}

type USBController struct {
	UniqueID      string   `json:"uniqueId"`
	SysFSID       string   `json:"sysFsId"`
	SysFSBusID    string   `json:"sysFsBusId"`
	HardwareClass string   `json:"hardwareClass"`
	Model         string   `json:"model"`
	Vendor        string   `json:"vendor"`
	Device        string   `json:"device"`
	SubVendor     string   `json:"subVendor"`
	SubDevice     string   `json:"subDevice"`
	Revision      string   `json:"revision"`
	Driver        string   `json:"driver"`
	DriverModules string   `json:"driverModules"`
	MemoryRange   []string `json:"memoryRanges"`
	IRQ           string   `json:"irq"`
	ModuleAlias   string   `json:"moduleAlias"`
	ConfigStatus  string   `json:"configStatus"`
}

type SerialController struct {
	UniqueID      string `json:"uniqueId"`
	SysFSID       string `json:"sysFsId"`
	SysFSBusID    string `json:"sysFsBusId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	Vendor        string `json:"vendor"`
	Device        string `json:"device"`
	SubVendor     string `json:"subVendor"`
	SubDevice     string `json:"subDevice"`
	Revision      string `json:"revision"`
	Driver        string `json:"driver"`
	MemoryRange   string `json:"memoryRange"`
	IRQ           string `json:"irq"`
	ModuleAlias   string `json:"moduleAlias"`
	ConfigStatus  string `json:"configStatus"`
}

type SystemPeripheral struct {
	UniqueID      string `json:"uniqueId"`
	ParentID      string `json:"parentId"`
	SysFSID       string `json:"sysFsId"`
	SysFSBusID    string `json:"sysFsBusId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	Vendor        string `json:"vendor"`
	Device        string `json:"device"`
	SubVendor     string `json:"subVendor"`
	SubDevice     string `json:"subDevice"`
	Revison       string `json:"revision"`
	Driver        string `json:"driver"`
	DriverModules string `json:"driverModules"`
	MemoryRange   string `json:"memoryRange"`
	IRQ           string `json:"irq"`
	ModuleAlias   string `json:"moduleAlias"`
	ConfigStatus  string `json:"configStatus"`
}

type VGACompatibleController struct {
	UniqueID      string   `json:"uniqueId"`
	SysFSID       string   `json:"sysFsId"`
	SysFSBusID    string   `json:"sysFsBusId"`
	HardwareClass string   `json:"hardwareClass"`
	Model         string   `json:"model"`
	Vendor        string   `json:"vendor"`
	Device        string   `json:"device"`
	SubVendor     string   `json:"subVendor"`
	SubDevice     string   `json:"subDevice"`
	Revision      string   `json:"revision"`
	Driver        string   `json:"driver"`
	DriverModules string   `json:"driverModules"`
	MemoryRange   []string `json:"memoryRanges"`
	IRQ           string   `json:"irq"`
	ModuleAlias   string   `json:"moduleAlias"`
	ConfigStatus  string   `json:"configStatus"`
}

type LCDMonitor struct {
	UniqueID          string   `json:"uniqueId"`
	ParentID          string   `json:"parentId"`
	HardwareClass     string   `json:"hardwareClass"`
	Model             string   `json:"model"`
	Vendor            string   `json:"vendor"`
	Device            string   `json:"device"`
	Resolution        []string `json:"resolutions"`
	Size              string   `json:"size"`
	YearOfManufacture int64    `json:"yearOfManufacture"`
	WeekOfManufacture int64    `json:"weekOfManufacture"`
	ConfigStatus      string   `json:"configStatus"`
	AttachedTo        string   `json:"attachedTo"`
}

type Disk struct {
	UniqueID        string `json:"uniqueId"`
	ParentID        string `json:"parentId"`
	SysFSID         string `json:"sysFsId"`
	SysFSBusID      string `json:"sysFsBusId"`
	SysFSDeviceLink string `json:"sysFsDeviceLink"`
	HardwareClass   string `json:"hardwareClass"`
	Model           string `json:"model"`
	Vendor          string `json:"vendor"`
	Device          string `json:"device"`
	Revision        string `json:"revision"`
	Driver          string `json:"driver"`
	DriverModules   string `json:"driverModules"`
	DeviceFile      string `json:"deviceFile"`
	DeviceFiles     string `json:"deviceFiles"`
	DeviceNumber    string `json:"deviceNumber"`
	BIOSID          string `json:"biosId"`
	DriveStatus     string `json:"driveStatus"`
	ConfigStatus    string `json:"configStatus"`
	AttachedTo      string `json:"attachedTo"`
}

type Partition struct {
	UniqueID      string `json:"uniqueId"`
	ParentID      string `json:"parentId"`
	SysFSID       string `json:"sysFsId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	DeviceFile    string `json:"deviceFile"`
	DeviceFiles   string `json:"deviceFiles"`
	ConfigStatus  string `json:"configStatus"`
	AttachedTo    string `json:"attachedTo"`
}

type BluetoothDevice struct {
	UniqueID      string `json:"uniqueId"`
	ParentID      string `json:"parentId"`
	SysFSID       string `json:"sysFsId"`
	SysFSBusID    string `json:"sysFsBusId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	Hotplug       string `json:"hotplug"`
	Vendor        string `json:"vendor"`
	Device        string `json:"device"`
	Driver        string `json:"driver"`
	Revision      string `json:"revision"`
	SerialID      string `json:"serialId"`
	Speed         string `json:"speed"`
	ModuleAlias   string `json:"moduleAlias"`
	ConfigStatus  string `json:"configStatus"`
	AttachedTo    string `json:"attachedTo"`
}

type UnclassifiedDevice struct {
	UniqueID      string `json:"uniqueId"`
	ParentID      string `json:"parentId"`
	SysFSID       string `json:"sysFsId"`
	SysFSBusID    string `json:"sysFsBusId"`
	HardwareClass string `json:"hardwareClassa"`
	Model         string `json"model"`
	Hotplug       string `json:"hotplug"`
	Vendor        string `json:"vendor"`
	Device        string `json:"device"`
	Revision      string `json:"revision"`
	Driver        string `json:"driver"`
	DriverModules string `json:"driverModules"`
	Speed         string `json:"speed"`
	ModuleAlias   string `json:"moduleAlias"`
	ConfigStatus  string `json:"configStatus"`
	AttachedTo    string `json:"attachedTo"`
}

type Hub struct {
	UniqueID      string `json:"uniqueId"`
	ParentID      string `json:"parentId"`
	SysFSID       string `json:"sysFsId"`
	SysFSBusID    string `json:"sysFsBusId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	Hotplug       string `json:"hotplug"`
	Vendor        string `json:"vendor"`
	Device        string `json:"device"`
	Driver        string `json:"driver"`
	DriverModules string `json:"driverModules"`
	Speed         string `json:"speed"`
	ModuleAlias   string `json:"moduleAlias"`
	ConfigStatus  string `json:"configStatus"`
	AttachedTo    string `json:"attachedTo"`
}

type Keyboard struct {
	UniqueID      string `json:"uniqueId"`
	ParentID      string `json:"parentId"`
	SysFSID       string `json:"sysFsId"`
	SysFSBusID    string `json:"sysFsBusId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	Hotplug       string `json:"hotplug"`
	Vendor        string `json:"vendor"`
	Device        string `json:"device"`
	Revision      string `json:"revision"`
	Driver        string `json:"driver"`
	DriverModules string `json:"driverModules"`
	CompatibleTo  string `json:"compatibleTo"`
	DeviceFile    string `json:"deviceFile"`
	DeviceFiles   string `json:"deviceFiles"`
	DeviceNumber  string `json:"deviceNumber"`
	Speed         string `json:"speed"`
	ConfigStatus  string `json:"configStatus"`
	AttachedTo    string `json:"attachedTo"`
}

type USBMouse struct {
	UniqueID      string `json:"uniqueId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	Vendor        string `json:"vendor"`
	Device        string `json:"device"`
	CompatibleTo  string `json:"compatibleTo"`
	DeviceFile    string `json:"deviceFile"`
	DeviceFiles   string `json:"deviceFiles"`
	DeviceNumber  string `json:"deviceNumber"`
	ConfigStatus  string `json:"configStatus"`
}

type CPU struct {
	UniqueID      string `json:"uniqueId"`
	HardwareClass string `json:"hardwareClass"`
	Arch          string `json:"arch"`
	Vendor        string `json:"vendor"`
	Model         string `json:"model"`
	Features      string `json:"features"`
	Clock         string `json:"clock"`
	BogoMips      string `json:"bogoMips"`
	Cache         string `json:"cache"`
	ConfigStatus  string `json:"configStatus"`
}

type Loopback struct {
	UniqueID      string `json:"uniqueId"`
	SysFSID       string `json:"sysFsId"`
	HardwareClass string `json:"hardwareClass"`
	Model         string `json:"model"`
	DeviceFile    string `json:"deviceFile"`
	LinkDetected  string `json:"linkDetected"`
	ConfigStatus  string `json:"configStatus"`
}

type Ethernet struct {
	UniqueID           string `json:"uniqueId"`
	SysFSID            string `json:"sysFsId"`
	SysFSDeviceLink    string `json:"sysFsDeviceLink"`
	HardwareClass      string `json:"hardwareClass"`
	Model              string `json:"model"`
	Driver             string `json:"driver"`
	DriverModule       string `json:"driverModule"`
	DeviceFile         string `json:"deviceFile"`
	HWAddress          string `json:"hwAddress"`
	PermanentHWAddress string `json:"permanentHwAddress"`
	LinkDetected       string `json:"linkDetected"`
	ConfigStatus       string `json:"configStatus"`
	AttachedTo         string `json:"attachedTo"`
}

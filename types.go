package hwparser

type HardwareInfo struct {
	BIOS                     BIOS                      `name:"BIOS"`
	System                   System                    `name:"System"`
	MainMemory               MainMemory                `name:"Main Memory"`
	FPUS                     []FPU                     `name:"FPU"`
	DMAControllers           []DMAController           `name:"DMA controller"`
	PICS                     []PIC                     `name:"PIC"`
	KeyboardControllers      []KeyboardController      `name:"Keyboard controller"`
	WLANControllers          []WLANController          `name:"WLAN controller"`
	SATAControllers          []SATAController          `name:"SATA controller"`
	PCIBridges               []PCIBridge               `name:"PCI bridge"`
	ISABridges               []ISABridge               `name:"ISA bridge"`
	CommunicationControllers []CommunicationController `name:"Communication controller"`
	AudioDevices             []AudioDevice             `name:"Audio device"`
	EthernetControllers      []EthernetController      `name:"Ethernet controller"`
	SMBuses                  []SMBus                   `name:"SMBus"`
	HostBridges              []HostBridge              `name:"Host bridge"`
	USBControllers           []USBController           `name:"USB Controller"`
	SerialControllers        []SerialController        `name:"Serial controller"`
	SystemPeripherals        []SystemPeripheral        `name:"System peripheral"`
	VGACompatibleControllers []VGACompatibleController `name:"VGA compatible controller"`
	LCDMonitors              []LCDMonitor              `name:"LCD Monitor"`
	Disks                    []Disk                    `name:"Disk"`
	Partitions               []Partition               `name:"Partition"`
	BluetoothDevices         []BluetoothDevice         `name:"Bluetooth Device"`
	UnclassifiedDevices      []UnclassifiedDevice      `name:"Unclassified device"`
	Hubs                     []Hub                     `name:"Hub"`
	Keyboards                []Keyboard                `name:"Keyboard"`
	USBMouses                []USBMouse                `name:"USB Mouse"`
	CPUS                     []CPU                     `name:"CPU"`
	Loopbacks                []Loopback                `name:"Loopback"`
	Ethernets                []Ethernet                `name:"Ethernet"`
}

type BIOS struct {
	UniqueID      string
	HardwareClass string
	ConfigStatus  string
}

type System struct {
	UniqueID      string
	HardwareClass string
	Model         string
	Formfactor    string
	ConfigStatus  string
}

type FPU struct {
	UniqueID      string
	HardwareClass string
	Model         string
	IOPort        []string
	ConfigStatus  string
}

type DMAController struct {
	UniqueID      string
	HardwareClass string
	Model         string
	IOPort        []string
	DMA           int64
	ConfigStatus  string
}

type PIC struct {
	UniqueID      string
	HardwareClass string
	Model         string
	IOPort        []string
	ConfigStatus  string
}

type KeyboardController struct {
	UniqueID      string
	HardwareClass string
	Model         string
	IOPort        []string
	ConfigStatus  string
}

type MainMemory struct {
	UniqueID      string
	HardwareClass string
	Model         string
	MemoryRange   string
	MemorySize    string
	ConfigStatus  string
}

type WLANController struct {
	UniqueID                string
	ParentID                string
	SysFSID                 string
	SysFSBusID              string
	HardwareClass           string
	Model                   string
	Vendor                  string
	Device                  string
	SubVendor               string
	SubDevice               string
	Revision                string
	Driver                  string
	DriverModules           string
	DeviceFile              string
	Features                string
	MemoryRange             string
	IRQ                     string
	HWAddress               string
	PermanentHWAddress      string
	LinkDetected            string
	WLANChannels            string
	WLANFrequences          string
	WLANEncryptionModes     string
	WLANAuthenticationModes string
	ModuleAlias             string
	ConfigStatus            string
	AttachedTo              string
}

type SATAController struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	Driver        string
	DriverModules string
	MemoryRange   []string
	IOPort        []string
	IRQ           string
	ModuleAlias   string
	ConfigStatus  string
}

type MemoryController struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	MemoryRange   []string
	ModuleAlias   string
	ConfigStatus  string
}

type Bridge struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	Driver        string
	IRQ           string
	ModuleAlias   string
	ConfigStatus  string
}

type ISABridge struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	Driver        string
	DriverModules string
	ModuleAlias   string
	ConfigStatus  string
}

type CommunicationController struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	Driver        string
	DriverModules string
	MemoryRange   []string
	IRQ           string
	ModuleAlias   string
	ConfigStatus  string
}

type AudioDevice struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClas  string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	Driver        string
	DriverModules string
	MemoryRange   []string
	IRQ           string
	ModuleAlias   string
	ConfigStatus  string
}

type PCIBridge struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	Driver        string
	IRQ           string
	ModuleAlias   string
	ConfigStatus  string
}

type EthernetController struct {
	UniqueID           string
	SysFSID            string
	SysFSBusID         string
	HardwareClass      string
	Model              string
	Vendor             string
	Device             string
	SubVendor          string
	SubDevice          string
	Revision           string
	Driver             string
	DriverModules      string
	DeviceFile         string
	MemoryRange        []string
	IRQ                string
	HWAddress          string
	PermanentHWAddress string
	LinkDetected       string
	ModuleAlias        string
	ConfigStatus       string
}

type SMBus struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	Driver        string
	DriverModules string
	MemoryRange   []string
	IRQ           string
	ModuleAlias   string
	ConfigStatus  string
}

type HostBridge struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	ModuleAlias   string
	ConfigStatus  string
}

type USBController struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	Driver        string
	DriverModules string
	MemoryRange   []string
	IRQ           string
	ModuleAlias   string
	ConfigStatus  string
}

type SerialController struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	Driver        string
	MemoryRange   string
	IRQ           string
	ModuleAlias   string
	ConfigStatus  string
}

type SystemPeripheral struct {
	UniqueID      string
	ParentID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revison       string
	Driver        string
	DriverModules string
	MemoryRange   string
	IRQ           string
	ModuleAlias   string
	ConfigStatus  string
}

type VGACompatibleController struct {
	UniqueID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	SubVendor     string
	SubDevice     string
	Revision      string
	Driver        string
	DriverModules string
	MemoryRange   []string
	IRQ           string
	ModuleAlias   string
	ConfigStatus  string
}

type LCDMonitor struct {
	UniqueID          string
	ParentID          string
	HardwareClass     string
	Model             string
	Vendor            string
	Device            string
	Resolution        []string
	Size              string
	YearOfManufacture int64
	WeekOfManufacture int64
	ConfigStatus      string
	AttachedTo        string
}

type Disk struct {
	UniqueID        string
	ParentID        string
	SysFSID         string
	SysFSBusID      string
	SysFSDeviceLink string
	HardwareClass   string
	Model           string
	Vendor          string
	Device          string
	Revision        string
	Driver          string
	DriverModules   string
	DeviceFile      string
	DeviceFiles     string
	DeviceNumber    string
	BIOSID          string
	DriveStatus     string
	ConfigStatus    string
	AttachedTo      string
}

type Partition struct {
	UniqueID      string
	ParentID      string
	SysFSID       string
	HardwareClass string
	Model         string
	DeviceFile    string
	DeviceFiles   string
	ConfigStatus  string
	AttachedTo    string
}

type BluetoothDevice struct {
	UniqueID      string
	ParentID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Hotplug       string
	Vendor        string
	Device        string
	Revision      string
	SerialID      string
	Speed         string
	ModuleAlias   string
	ConfigStatus  string
	AttachedTo    string
}

type UnclassifiedDevice struct {
	UniqueID      string
	ParentID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Hotplug       string
	Vendor        string
	Device        string
	Revision      string
	Driver        string
	DriverModules string
	Speed         string
	ModuleAlias   string
	ConfigStatus  string
	AttachedTo    string
}

type Hub struct {
	UniqueID      string
	ParentID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Hotplug       string
	Vendor        string
	Device        string
	Driver        string
	DriverModules string
	Speed         string
	ModuleAlias   string
	ConfigStatus  string
	AttachedTo    string
}

type Keyboard struct {
	UniqueID      string
	ParentID      string
	SysFSID       string
	SysFSBusID    string
	HardwareClass string
	Model         string
	Hotplug       string
	Vendor        string
	Device        string
	Revision      string
	Driver        string
	DriverModules string
	CompatibleTo  string
	DeviceFile    string
	DeviceFiles   string
	DeviceNumber  string
	Speed         string
	ConfigStatus  string
	AttachedTo    string
}

type USBMouse struct {
	UniqueID      string
	HardwareClass string
	Model         string
	Vendor        string
	Device        string
	CompatibleTo  string
	DeviceFile    string
	DeviceFiles   string
	DeviceNumber  string
	ConfigStatus  string
}

type CPU struct {
	UniqueID      string
	HardwareClass string
	Arch          string
	Vendor        string
	Model         string
	Features      string
	Clock         string
	BogoMips      string
	Cache         string
	ConfigStatus  string
}

type Loopback struct {
	UniqueID      string
	SysFSID       string
	HardwareClass string
	Model         string
	DeviceFile    string
	LinkDetected  string
	ConfigStatus  string
}

type Ethernet struct {
	UniqueID           string
	SysFSID            string
	SysFSDeviceLink    string
	HardwareClass      string
	Model              string
	Driver             string
	DriverModule       string
	DeviceFile         string
	HWAddress          string
	PermanentHWAddress string
	LinkDetected       string
	ConfigStatus       string
	AttachedTo         string
}

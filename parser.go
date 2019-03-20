package hwparser

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	endDebugPattern = `=*\send debug info\s=*`
	titlePattern    = `\d{2}:\s.+`
)

var (
	errPatternNotFound = fmt.Errorf("debug info string not found")
)

func GetHardwareInfo() *HardwareInfo {
	return parse()
}

func getRawHWData() ([]byte, error) {
	cmd := exec.Command("hwinfo")
	bytes, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	pattern := regexp.MustCompile(endDebugPattern)
	indexes := pattern.FindSubmatchIndex(bytes)

	if len(indexes) == 0 {
		return nil, errPatternNotFound
	}

	return bytes[indexes[1]:], nil
}

func parse() *HardwareInfo {
	hwData, err := getRawHWData()
	if err != nil {
		log.Fatal(err)
	}

	hwInfo := &HardwareInfo{}

	pattern := regexp.MustCompile(titlePattern)
	indexes := pattern.FindAllStringSubmatchIndex(string(hwData), -1)
	for i := 0; i < len(indexes)-1; i++ {
		dataStart := indexes[i][1]
		dataEnd := indexes[i+1][0]
		title := parseTitle(strings.TrimSpace(string(hwData[indexes[i][0]:indexes[i][1]])))
		data := strings.TrimSpace(string(hwData[dataStart:dataEnd]))

		switch title {

		case "BIOS":
			reflectStruct(&hwInfo.BIOS, data)
			break

		case "System":
			reflectStruct(&hwInfo.System, data)
			break

		case "Main Memory":
			reflectStruct(&hwInfo.MainMemory, data)
			break

		case "FPU":
			fpu := FPU{}
			reflectStruct(&fpu, data)
			hwInfo.FPUS = append(hwInfo.FPUS, fpu)
			break

		case "DMA controller":
			dma := DMAController{}
			reflectStruct(&dma, data)
			hwInfo.DMAControllers = append(hwInfo.DMAControllers, dma)
			break

		case "PIC":
			pic := PIC{}
			reflectStruct(&pic, data)
			hwInfo.PICS = append(hwInfo.PICS, pic)
			break

		case "Keyboard controller":
			kc := KeyboardController{}
			reflectStruct(&kc, data)
			hwInfo.KeyboardControllers = append(hwInfo.KeyboardControllers, kc)
			break

		case "WLAN controller":
			wc := WLANController{}
			reflectStruct(&wc, data)
			hwInfo.WLANControllers = append(hwInfo.WLANControllers, wc)
			break

		case "SATA controller":
			sc := SATAController{}
			reflectStruct(&sc, data)
			hwInfo.SATAControllers = append(hwInfo.SATAControllers, sc)
			break

		case "PCI bridge":
			pciBridge := PCIBridge{}
			reflectStruct(&pciBridge, data)
			hwInfo.PCIBridges = append(hwInfo.PCIBridges, pciBridge)
			break

		case "ISA bridge":
			isaBridge := ISABridge{}
			reflectStruct(&isaBridge, data)
			hwInfo.ISABridges = append(hwInfo.ISABridges, isaBridge)
			break

		case "Communication controller":
			cc := CommunicationController{}
			reflectStruct(&cc, data)
			hwInfo.CommunicationControllers = append(hwInfo.CommunicationControllers, cc)
			break

		case "Audio device":
			audio := AudioDevice{}
			reflectStruct(&audio, data)
			hwInfo.AudioDevices = append(hwInfo.AudioDevices, audio)
			break

		case "Ethernet controller":
			ec := EthernetController{}
			reflectStruct(&ec, data)
			hwInfo.EthernetControllers = append(hwInfo.EthernetControllers, ec)
			break

		case "SMBus":
			smb := SMBus{}
			reflectStruct(&smb, data)
			hwInfo.SMBuses = append(hwInfo.SMBuses, smb)
			break

		case "Host bridge":
			hb := HostBridge{}
			reflectStruct(&hb, data)
			hwInfo.HostBridges = append(hwInfo.HostBridges, hb)
			break

		case "USB Controller":
			uc := USBController{}
			reflectStruct(&uc, data)
			hwInfo.USBControllers = append(hwInfo.USBControllers, uc)
			break

		case "Serial controller":
			sc := SerialController{}
			reflectStruct(&sc, data)
			hwInfo.SerialControllers = append(hwInfo.SerialControllers, sc)
			break

		case "System peripheral":
			sp := SystemPeripheral{}
			reflectStruct(&sp, data)
			hwInfo.SystemPeripherals = append(hwInfo.SystemPeripherals, sp)
			break

		case "VGA compatible controller":
			vcc := VGACompatibleController{}
			reflectStruct(&vcc, data)
			hwInfo.VGACompatibleControllers = append(hwInfo.VGACompatibleControllers, vcc)
			break

		case "LCD Monitor":
			monitor := LCDMonitor{}
			reflectStruct(&monitor, data)
			hwInfo.LCDMonitors = append(hwInfo.LCDMonitors, monitor)
			break

		case "Disk":
			disk := Disk{}
			reflectStruct(&disk, data)
			hwInfo.Disks = append(hwInfo.Disks, disk)
			break

		case "Partition":
			partition := Partition{}
			reflectStruct(&partition, data)
			hwInfo.Partitions = append(hwInfo.Partitions, partition)
			break

		case "Bluetooth Device":
			bluetooth := BluetoothDevice{}
			reflectStruct(&bluetooth, data)
			hwInfo.BluetoothDevices = append(hwInfo.BluetoothDevices, bluetooth)
			break

		case "Unclassified device":
			unknown := UnclassifiedDevice{}
			reflectStruct(&unknown, data)
			hwInfo.UnclassifiedDevices = append(hwInfo.UnclassifiedDevices, unknown)
			break

		case "Hub":
			hub := Hub{}
			reflectStruct(&hub, data)
			hwInfo.Hubs = append(hwInfo.Hubs, hub)
			break

		case "Keyboard":
			keyboard := Keyboard{}
			reflectStruct(&keyboard, data)
			hwInfo.Keyboards = append(hwInfo.Keyboards, keyboard)
			break

		case "USB Mouse":
			mouse := USBMouse{}
			reflectStruct(&mouse, data)
			hwInfo.USBMouses = append(hwInfo.USBMouses, mouse)
			break

		case "CPU":
			cpu := CPU{}
			reflectStruct(&cpu, data)
			hwInfo.CPUS = append(hwInfo.CPUS, cpu)
			break

		case "Loopback":
			loopback := Loopback{}
			reflectStruct(&loopback, data)
			hwInfo.Loopbacks = append(hwInfo.Loopbacks, loopback)
			break

		case "Ethernet":
			ethernet := Ethernet{}
			reflectStruct(&ethernet, data)
			hwInfo.Ethernets = append(hwInfo.Ethernets, ethernet)
			break
		}
	}

	return hwInfo
}

func reflectStruct(target interface{}, data string) {
	values := reflect.ValueOf(target).Elem()

	lines := strings.Split(data, "\n")
	for i := 1; i < len(lines); i++ {
		spacelessLine := strings.Replace(lines[i], " ", "", -1)
		splitData := strings.Split(strings.TrimSpace(spacelessLine), ":")
		fieldName := splitData[0]
		fieldValue := splitData[1]
		structField := values.FieldByName(fieldName)
		if structField.Kind() == reflect.Invalid {
		} else {
			setValue(structField, fieldValue)
		}
	}
}

func setValue(field reflect.Value, value interface{}) {
	fieldType := field.Type()

	if fieldType.Kind() == reflect.Slice {
		newSlice := reflect.Append(field, reflect.ValueOf(value))
		field.Set(newSlice)
	} else {
		switch fieldType.Name() {
		case "string":
			field.SetString(value.(string))

		case "int64":
			parsedInt, err := strconv.ParseInt(value.(string), 10, 64)
			if err != nil {
				log.Fatalf(
					"failed to parse value. type: %s, value: %s, err: %s",
					fieldType.Name(),
					value,
					err,
				)
			}
			field.SetInt(parsedInt)
		}
	}

}

func parseTitle(title string) string {
	firstSplit := strings.Split(title, ":")[2]
	parenthesisSplit := strings.Split(firstSplit, "(")
	if len(parenthesisSplit) > 1 {
		return strings.Join(strings.Fields(parenthesisSplit[0])[1:], " ")
	} else {
		return strings.Join(strings.Fields(firstSplit)[1:], " ")
	}
}

package cpx

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gotmc/lxi"
)

const defaultIp = "192.168.0.102"

// lxiDeviceData structure to hold package data
type lxiDeviceData struct {
	ipAddress   string
	port        int
	isConnected bool
	device      *lxi.Device
}

// PowerSupplyInterface Interface for CPX package
type PowerSupplyInterface interface {
	Connect()
	IsConnected() bool
	Query(string) string
	GetVoltage() (float64, error)
	GetCurrent() (float64, error)
	Close()
}

// BenchPowerSupply constructor (optional)
func BenchPowerSupply(ip string, port int) PowerSupplyInterface {
	return &lxiDeviceData{ip, port, false, nil}
}

// Connect Attempts to make a TCP/IP connection with the bench power
// supply via the LXI (LAN Extension Interface) standard.
func (data *lxiDeviceData) Connect() {

	var err error
	address := fmt.Sprintf("TCPIP0::%s::%v::SOCKET", data.ipAddress, data.port)
	data.device, err = lxi.NewDevice(address)

	if err != nil {
		log.Printf("CPX error: %s", err)
	} else {
		fmt.Println("Connection established")
		data.isConnected = true
	}
}

//Close Cleanly end the connection to the bench power supply
func (data *lxiDeviceData) Close() {
	data.device.Close()
}

//IsConnected Check to see the status of the link between program and bench power supply
func (data *lxiDeviceData) IsConnected() bool {
	return data.isConnected
}

// Query Attempt to read a single telemetry type (i.e. volts, current, etc)
func (data *lxiDeviceData) Query(telemetryType string) string {

	telemetryValue, err := data.device.Query(telemetryType)

	if err != nil {
		log.Printf("Error reading: %v", err)
	}

	return telemetryValue
}

//GetVoltage Returns the voltage of the power supply as an int
func (data *lxiDeviceData) GetVoltage() (float64, error) {
	voltageString, err := data.device.Query("v1O?")
	if err != nil {
		return 0, err
	} else {
		// remove suffix and convert string to float64
		var voltage = strings.TrimSpace(voltageString)
		voltageTrimmedString := strings.TrimSuffix(voltage, "V")
		voltageFloat, _ := strconv.ParseFloat(voltageTrimmedString, 64)
		return voltageFloat, nil
	}
}

//GetVoltage Returns the voltage of the power supply as an int
func (data *lxiDeviceData) GetCurrent() (float64, error) {
	currentString, err := data.device.Query("i1O?")
	if err != nil {
		return 0, err
	} else {
		// remove suffix and convert string to float64
		var current = strings.TrimSpace(currentString)
		currentTrimmedString := strings.TrimSuffix(current, "I")
		currentFloat, _ := strconv.ParseFloat(currentTrimmedString, 64)
		return currentFloat, nil
	}
}

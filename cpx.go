package cpx

import (
	"fmt"
	"log"

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

//IsConnected Check to see the status of the link between program and bench power supply
func (data *lxiDeviceData) IsConnected() bool {
	return data.isConnected
}

// Query Attempt to read a single telemetry type (i.e. volts, current, etc)
func (data *lxiDeviceData) Query(telemetryType string) string {
	ws := fmt.Sprintf("%s?", telemetryType)

	telemetryValue, err := data.device.Query(ws)

	log.Printf("Completed %s query", ws)

	if err != nil {
		log.Printf("Error reading: %v", err)
	} else {
		log.Printf("Query %s? = %s", telemetryType, telemetryValue)
	}

	return telemetryValue
}

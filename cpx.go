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
	Connect() bool
	IsConnected() bool
}

// BenchPowerSupply constructor (optional)
func BenchPowerSupply(ip string, port int) PowerSupplyInterface {
	return &lxiDeviceData{ip, port, false, nil}
}

// Connect Attempts to make a TCP/IP connection with the bench power
// supply via the LXI (LAN Extension Interface) standard.
func (data *lxiDeviceData) Connect() bool {
	//TODO parse ip arguments into NewDevice()
	var err error
	data.device, err = lxi.NewDevice("TCPIP0::127.0.0.1::50001::SOCKET")

	connectionEstabilished := false
	if err != nil {
		log.Fatalf("CPX error: %s", err)
	}

	if data.device != nil {
		fmt.Println("Connection established")
		connectionEstabilished = true
	}
	return connectionEstabilished
}

//IsConnected Check to see the status of the link between program and bench power supply
func (data *lxiDeviceData) IsConnected() bool {
	return data.isConnected
}

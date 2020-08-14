package cpx

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
)

func TestConstructor(t *testing.T) {

	device := BenchPowerSupply("127.0.0.1", 50001)

	if device.IsConnected() == true {
		t.Fatalf("Expected False | Actual: %t\n", device.IsConnected())
	}
}

func TestConnection(t *testing.T) {
	cases := map[string]struct {
		ip       string
		port     int
		expected bool
	}{
		"loopback_std_port":     {ip: "127.0.0.1", port: 50001, expected: true},
		"loopback_non_std_port": {ip: "127.0.0.1", port: 50000, expected: false},
	}

	basicTCPServer("127.0.0.1:50001")

	for name, tc := range cases {

		device := BenchPowerSupply(tc.ip, tc.port)
		device.Connect()

		if device.IsConnected() != tc.expected {
			t.Fatalf("%s | Expected %v | Actual: %t\n", name, tc.expected, device.IsConnected())
		}
		//device.Close()
	}
}

func TestGetVoltage(t *testing.T) {
	complexTCPServer("127.0.0.1:50002")

	device := BenchPowerSupply("127.0.0.1", 50002)
	device.Connect()

	if device.IsConnected() != true {
		t.Fatal("Failed to connect to local test server")
	}

	voltage, _ := device.GetVoltage()
	if voltage != -0.0204 {
		t.Fatal("Returned incorrect voltage reading")
	}
}

func TestGetCurrent(t *testing.T) {
	complexTCPServer("127.0.0.1:50003")

	device := BenchPowerSupply("127.0.0.1", 50003)
	device.Connect()

	if device.IsConnected() != true {
		t.Fatal("Failed to connect to local test server")
	}

	voltage, _ := device.GetCurrent()
	if voltage != 0.303 {
		t.Fatal("Returned incorrect current reading")
	}
}

//basicTCPServer helper function to generate simple TCP server that can test whether a client
//connection can be established. No data is sent or received.
func basicTCPServer(tcpAddress string) {
	ln, err := net.Listen("tcp", tcpAddress)

	if err != nil {
		fmt.Printf("basicTcpServer - Listener Error: %s", err.Error())
	}

	go func() {
		defer ln.Close()
		_, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept Error: ", err)
			return
		}
	}()
}

//complexTCPServer helper function to generate simple TCP server that can test whether a client
//connection can be established. No data is sent or received.
func complexTCPServer(tcpAddress string) {
	ln, err := net.Listen("tcp", tcpAddress)

	if err != nil {
		fmt.Printf("basicTcpServer - Listener Error: %s", err.Error())
	}

	go func() {
		defer ln.Close()
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept Error: ", err)
			return
		}
		handleConnection(c)
	}()
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())

	netData, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	temp := strings.TrimSpace(netData)
	temp = strings.TrimSuffix(temp, "\n")
	if temp == "v1O?" {
		fakeVoltage := "-0.0204V\n"
		c.Write([]byte(fakeVoltage))
	} else if temp == "i1O?" {
		fakeCurrent := "0.303I\n"
		c.Write([]byte(fakeCurrent))
	}

	return
}

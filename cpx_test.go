package cpx

import (
	"fmt"
	"net"
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

	for name, tc := range cases {
		server := basicTCPServer("127.0.0.1:50001")

		device := BenchPowerSupply(tc.ip, tc.port)
		device.Connect()

		if device.IsConnected() != tc.expected {
			t.Fatalf("%s | Expected %v | Actual: %t\n", name, tc.expected, device.IsConnected())
		}

		_ = server
	}
}

//basicTCPServer helper function to generate simple TCP server that can test whether a client
//connection can be established. No data is sent or received.
func basicTCPServer(ip string) net.Conn {
	ln, err := net.Listen("tcp", "127.0.0.1:50001")

	if err != nil {
		fmt.Println("Listener Error")
	}

	var server net.Conn
	go func() {
		defer ln.Close()
		server, err = ln.Accept()
		if err != nil {
			fmt.Println("Accept Error: ", err)
			return
		}

	}()

	return server

}

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

//var server net.Conn

func TestConnection(t *testing.T) {

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

	_ = server

	device := BenchPowerSupply("127.0.0.1", 50001)
	device.Connect()

	if device.IsConnected() == true {
		t.Fatalf("Expected true | Actual: %t\n", device.IsConnected())
	}

}

func basicTcpServer(ip string) net.Conn {
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

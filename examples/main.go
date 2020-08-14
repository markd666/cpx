package main

import (
	"log"
	"os"
	"time"

	"github.com/markd666/cpx"
)

func main() {
	t := cpx.BenchPowerSupply("192.168.0.102", 9221)
	t.Connect()

	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(f)

	for {
		log.Println(t.GetVoltage())
		log.Println(t.GetCurrent())

		time.Sleep(100 * time.Millisecond)
	}
}

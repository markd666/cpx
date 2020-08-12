package main

import "github.com/markd666/cpx"

func main() {
	t := cpx.BenchPowerSupply("192.168.0.100", 9990)
	t.Connect()
}

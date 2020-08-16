# cpx
[![Go Report Card](https://goreportcard.com/badge/github.com/markd666/cpx)](https://goreportcard.com/report/github.com/markd666/cpx)
[![GoDoc](https://godoc.org/github.com/markd666/cpx?status.svg)](https://godoc.org/github.com/markd666/cpx)
[![Build Status](https://travis-ci.org/markd666/cpx.svg?branch=master)](https://travis-ci.org/markd666/cpx)

Go interface to the CPX Series Bench Power Supplies by TTi

## Install

go get -v github.com/markd666/cpx

or if you use go modules add to imports path and call

`go mod download'

## Use 

Basic example can be found in the /examples folder

First need to create a benchPowerSupply object providing ip address and port number of the power supply.

```golang
device := cpx.BenchPowerSupply("192.168.0.102", 9221)
```

Then attempt to make a TCP/IP connection to it. Making sure the power supply in question is ping-able from the machine running the code.

``` golang
device.Connect()
```

Data can be requested from the bench power supply via the following commands.

```golang
voltage, err := device.GetVoltage()
```
and 
```golang
current,err := device.GetCurrent()
```

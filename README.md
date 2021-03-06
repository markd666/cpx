# cpx
[![Go Report Card](https://goreportcard.com/badge/github.com/markd666/cpx)](https://goreportcard.com/report/github.com/markd666/cpx)
[![GoDoc](https://godoc.org/github.com/markd666/cpx?status.svg)](https://godoc.org/github.com/markd666/cpx)
[![Build Status](https://travis-ci.org/markd666/cpx.svg?branch=master)](https://travis-ci.org/markd666/cpx)
[![License Badge][license badge]][LICENSE]

Go interface to the CPX Series Bench Power Supplies by [AIM TTi](https://www.aimtti.com/product-category/dc-power-supplies/aim-cpxseries) 

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
## Examples

The examples folder contains a basic golang program that utalises the CPX library, connects to a bench power supply and logs voltage/current data to a file.

```golang
go run main.go
```

## Scripts

Included in the repository are python3 scripts to parse data generated by the example program. The scripts typically require the pandas and matplotlib libraries, which can be installed eith pip.

```python3
pip install pandas matplotlib
```
![Script_Plots](https://github.com/markd666/cpx/blob/master/scripts/example_capture.PNG?raw=true "Script Plots")

## Notes

* This code has only been tested on the CPX400SP series bench power supply.
* Only basic functionality is currently included in teh library. Sofar only gathering 


[LICENSE]: https://github.com/markd666/cpx/blob/master/LICENSE
[license badge]: https://img.shields.io/badge/license-MIT-blue
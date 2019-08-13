package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (addr IPAddr) String() string {
	var str string
	str += strconv.Itoa(int(addr[0])) + "."
	str += strconv.Itoa(int(addr[1])) + "."
	str += strconv.Itoa(int(addr[2])) + "."
	str += strconv.Itoa(int(addr[3]))

	return str
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

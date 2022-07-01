package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

//string interface which auotmaticaaly search string for printing
func (ip IPAddr) String() string {
	str := make([]string, len(ip))
	for i, val := range ip {
		str[i] = strconv.Itoa(int(val))
	}
	return strings.Join(str, ".")
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

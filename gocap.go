package main

import (
	"fmt"
	_ "github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	version := pcap.Version()
	fmt.Println(version)
	devices, _ := pcap.FindAllDevs()
	fmt.Println(devices)
}

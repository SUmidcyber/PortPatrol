package device

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Interface() {
	fmt.Println("Welcome, you will protect your server with Interface.")
	fmt.Print("Please enter your interface tour: ")

	// TODO: Implement interface tour
	var InterfaceType string
	fmt.Scanln(&InterfaceType)

	handle, err := pcap.OpenLive(InterfaceType, 1024, true, pcap.BlockForever)
	if err != nil {
		fmt.Printf("Error opening interface: %v\n", err)
		return
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}

}

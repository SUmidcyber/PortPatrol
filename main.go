package main

import (
	"fmt"

	"github.com/SUmidcyber/PortPatrol/device"
	"github.com/SUmidcyber/PortPatrol/firewall"
	"github.com/SUmidcyber/PortPatrol/packets"
)

func main() {
	fmt.Println("To monitor and analyze network traffic. [1] \n To close open ports [2]")
	fmt.Print("Your choice: ")
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		device.Interface()
	case "2":
		packets.Server_ip()
	}

	fmt.Print("Do you want to close ports from log file [1] or manually [2]?: ")
	var closeChoice string
	fmt.Scanln(&closeChoice)

	switch closeChoice {
	case "1":
		firewall.OpenFile()
	case "2":
		fmt.Print("Enter port to close: ")
		var port string
		fmt.Scanln(&port)
		firewall.ClosePort(port)
	}
}

package packets

import (
	"fmt"
	"os"
	"os/exec"
)

func Server_ip() {
	fmt.Println("Welcome Please enter your server IP to start the process.")
	fmt.Print("Server IP: ")
	var Serverips string
	fmt.Scanln(&Serverips)
	// The $LOG/PATH$ Command
	fmt.Print("$LOG/PATH$: ")
	var LogPats string
	fmt.Scanln(&LogPats)

	cmd := exec.Command("nmap", "-p-", "-Pn", "-v", Serverips) // Nmap Port scan
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Erorr %s\n", err)
		return
	}

	// Write output .yml
	file, err := os.Create(string(LogPats))

	if err != nil {
		fmt.Printf("Error creating file %s\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(string(output))
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
	}

	fmt.Printf("Output written to %s\n", LogPats) // You look LogPats
}

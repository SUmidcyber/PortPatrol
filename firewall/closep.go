package firewall

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func OpenFile() {
	fmt.Print("$LOG/PATH$: ")
	var LogFirewall string
	fmt.Scanln(&LogFirewall)

	file, err := os.Open(LogFirewall)
	if err != nil {
		fmt.Println("PATH of Log File is not correct...", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		port := strings.TrimSpace(scanner.Text())
		ClosePort(port)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("LOG file reading error...", err)
	}
}

func ClosePort(port string) {
	cmd := exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--dport", port, "-j", "DROP")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to close port %s: %v\n", port, err)
	} else {
		fmt.Printf("Port %s was closed successfully.\n", port)
	}
}

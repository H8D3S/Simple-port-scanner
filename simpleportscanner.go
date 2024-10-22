package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Define the target hostname or IP address
	hostname := "asuracomic.net"

	// Define the range of ports you want to scan
	startPort := 1
	endPort := 1024

	// Set a timeout for the connection attempts
	timeout := 1 * time.Second

	fmt.Printf("Starting scan on %s from port %d to %d\n", hostname, startPort, endPort)

	for port := startPort; port <= endPort; port++ {
		// Try to scan each port
		if scanPort(hostname, port, timeout) {
			fmt.Printf("Port %d is open\n", port)
		}
	}
}

// scanPort attempts to connect to the target hostname on the specified port
// It returns true if the port is open, false otherwise.
func scanPort(hostname string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		// If there's an error, it means the port is closed or filtered
		return false
	}
	// If no error, the port is open
	conn.Close() // Close the connection
	return true
}

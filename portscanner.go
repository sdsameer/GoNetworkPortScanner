package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// scanPort checks if a specific port on the target address is open.
func scanPort(target string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()

	fmt.Printf("Port %d is open\n", port)
}

// scanRange scans a range of ports on a given target.
func scanRange(target string, startPort, endPort int) {
	fmt.Printf("Scanning ports %d-%d on %s...\n", startPort, endPort, target)

	var wg sync.WaitGroup
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go scanPort(target, port, &wg)
	}

	wg.Wait()
	fmt.Println("Scanning completed.")
}

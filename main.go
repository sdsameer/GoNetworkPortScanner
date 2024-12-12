package main

import "fmt"

func main() {
	var target string
	var startPort, endPort int

	fmt.Println("Network Port Scanner")
	fmt.Print("Enter the target IP address or domain: ")
	fmt.Scanln(&target)

	fmt.Print("Enter the starting port: ")
	fmt.Scanln(&startPort)

	fmt.Print("Enter the ending port: ")
	fmt.Scanln(&endPort)

	scanRange(target, startPort, endPort)
}

package main

import "fmt"

// IPAddr represents an IPv4 address as an array of 4 bytes.
type IPAddr [4]byte

// String method implements the fmt.Stringer interface for the IPAddr type.
// It converts the IP address to a string in dotted quad format (e.g., "127.0.0.1").

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func stringers() {
	// A map to associate hostnames with their corresponding IP addresses.
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1}, // The loopback address for local testing.
		"googleDNS": {8, 8, 8, 8},   // Google's public DNS server address.
	}

	// Iterate over the map and print each hostname with its associated IP address.
	for name, ip := range hosts {
		// %v uses the String() method for IPAddr to format the IP as a string.
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// Exercise: Stringers
// IPAddr 型を実装してみましょう IPアドレスをドットで4つに区切った( dotted quad )表現で出力するため、 fmt.Stringer インタフェースを実装してください。
package main

import (
	"fmt"
)

// IPAddr presents IP address.
type IPAddr [4]byte

// String makes dotted quad representation of IP address as string.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
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

package main

import (
	"fmt"
	"larmic/sonnen-charger-api/api/client"
)

func main() {
	c := client.NewClient("10.0.40.200")

	fmt.Println(c.ReadSettings())
	fmt.Println(c.ReadNumberOfConnectors())
	fmt.Println(c.ReadConnector(1))
	// Read unknown connector for testing
	fmt.Println(c.ReadConnector(2))
}

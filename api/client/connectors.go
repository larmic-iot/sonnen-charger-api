package client

import (
	"fmt"
	"strconv"
)

const (
	ChargerConnectorsNumber     = 1020
	ChargerConnectorBaseAddress = 1022
)

type Connector struct {
	Type string
}

func (c *ChargerClient) ReadNumberOfConnectors() int {
	fmt.Println("### Reading charger connectors...")
	_ = c.client.Open()

	register := c.readRegister(ChargerConnectorsNumber, "Number of connectors", 1)

	_ = c.client.Close()

	return int(register[0])
}

func (c *ChargerClient) ReadConnector(number int) Connector {
	_ = c.client.Open()

	registerAddress := uint16(ChargerConnectorBaseAddress + (number-1)*100)

	register := c.readRegister(registerAddress, "Connector number "+strconv.Itoa(number), 1)

	_ = c.client.Close()

	var connectorType string

	if register[0] == 1 {
		connectorType = "SocketType2"
	} else if register[0] == 2 {
		connectorType = "CableType2"
	} else {
		connectorType = "Unknown"
	}

	return Connector{
		Type: connectorType,
	}
}

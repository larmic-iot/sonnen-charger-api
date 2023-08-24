package client

import (
	"fmt"
	"strconv"
)

const (
	ChargerConnectorsNumber      = 1020
	ChargerConnectorBaseAddress  = 1022
	ChargerConnectorPhaseAddress = 1023
)

type ConnectorType string

const (
	SocketType2 ConnectorType = "SocketType2"
	CableType2                = "CableType2"
	Unknown                   = "Unknown"
)

type Connector struct {
	Type   ConnectorType
	Phases int
}

func (c *ChargerClient) ReadNumberOfConnectors() int {
	fmt.Println("### Reading charger connectors...")
	_ = c.client.Open()

	register := c.readRegister(ChargerConnectorsNumber, "Number of connectors", 1)

	_ = c.client.Close()

	return int(register[0])
}

func (c *ChargerClient) ReadConnector(number int) Connector {
	return Connector{
		Type:   c.readConnectorType(number),
		Phases: c.readNumberOfPhases(number),
	}
}

func (c *ChargerClient) readConnectorType(number int) ConnectorType {
	_ = c.client.Open()

	registerAddress := uint16(ChargerConnectorBaseAddress + (number-1)*100)
	register := c.readRegister(registerAddress, "Connector number "+strconv.Itoa(number), 1)

	_ = c.client.Close()

	var connectorType ConnectorType

	if register[0] == 1 {
		connectorType = SocketType2
	} else if register[0] == 2 {
		connectorType = CableType2
	} else {
		connectorType = Unknown
	}

	return connectorType
}

func (c *ChargerClient) readNumberOfPhases(number int) int {
	_ = c.client.Open()

	registerAddress := uint16(ChargerConnectorPhaseAddress + (number-1)*100)
	register := c.readRegister(registerAddress, "Number of phases "+strconv.Itoa(number), 1)

	_ = c.client.Close()

	return int(register[0])
}

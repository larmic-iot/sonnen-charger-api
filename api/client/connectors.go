package client

import (
	"fmt"
	"strconv"
)

const (
	ChargerConnectorsNumber      = 1020
	ChargerConnectorBaseAddress  = 1022
	ChargerConnectorPhaseAddress = 1023
	ChargerConnectorL1Address    = 1024
	ChargerConnectorL2Address    = 1025
	ChargerConnectorL3Address    = 1026
)

type ConnectorType string

const (
	SocketType2 ConnectorType = "SocketType2"
	CableType2                = "CableType2"
	Unknown                   = "Unknown"
)

type ConnectorPhases struct {
	L1 int
	L2 int
	L3 int
}

type Connector struct {
	Type           ConnectorType
	NumberOfPhases int
	Phases         ConnectorPhases
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
		Type:           c.readConnectorType(number),
		NumberOfPhases: c.readNumberOfPhases(number),
		Phases:         c.readConnectorPhases(number),
	}
}

func (c *ChargerClient) readConnectorType(connectorNumber int) ConnectorType {
	_ = c.client.Open()

	registerAddress := uint16(ChargerConnectorBaseAddress + +getConnectorOffset(connectorNumber))
	register := c.readRegister(registerAddress, "Connector number "+strconv.Itoa(connectorNumber), 1)

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

func (c *ChargerClient) readNumberOfPhases(connectorNumber int) int {
	_ = c.client.Open()

	registerAddress := uint16(ChargerConnectorPhaseAddress + getConnectorOffset(connectorNumber))
	register := c.readRegister(registerAddress, "Number of phases of connector "+strconv.Itoa(connectorNumber), 1)

	_ = c.client.Close()

	return int(register[0])
}

func (c *ChargerClient) readConnectorPhases(connectorNumber int) ConnectorPhases {
	_ = c.client.Open()

	registerAddressL1 := uint16(ChargerConnectorL1Address + getConnectorOffset(connectorNumber))
	registerL1 := c.readRegister(registerAddressL1, "L1 connected phase of connector "+strconv.Itoa(connectorNumber), 1)

	registerAddressL2 := uint16(ChargerConnectorL2Address + getConnectorOffset(connectorNumber))
	registerL2 := c.readRegister(registerAddressL2, "L2 connected phase of connector "+strconv.Itoa(connectorNumber), 1)

	registerAddressL3 := uint16(ChargerConnectorL3Address + getConnectorOffset(connectorNumber))
	registerL3 := c.readRegister(registerAddressL3, "L3 connected phase of connector "+strconv.Itoa(connectorNumber), 1)

	_ = c.client.Close()

	return ConnectorPhases{
		L1: int(registerL1[0]),
		L2: int(registerL2[0]),
		L3: int(registerL3[0]),
	}
}

func getConnectorOffset(connectorNumber int) int {
	return (connectorNumber - 1) * 100
}

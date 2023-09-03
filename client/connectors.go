package client

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
)

const (
	ChargerConnectorStatus       = 0
	ChargerConnectorsNumber      = 1020
	ChargerConnectorBaseAddress  = 1022
	ChargerConnectorPhaseAddress = 1023
	ChargerConnectorL1Address    = 1024
	ChargerConnectorL2Address    = 1025
	ChargerConnectorL3Address    = 1026
	ChargerConnectorMaxCurrent   = 1028
)

type ConnectorType string
type ConnectorStatus string

const (
	SocketType2 ConnectorType = "SocketType2"
	CableType2                = "CableType2"
	Unknown                   = "Unknown"
)

const (
	Available                  ConnectorStatus = "Available"
	ConnectTheCable                            = "Connect the cable"
	WaitingForVehicleToRespond                 = "Waiting for vehicle to respond"
	Charging                                   = "Charging"
	VehicleHasPausedCharging                   = "Vehicle has paused charging"
	EVSEHasPausedCharging                      = "EVSE has paused charging"
	ChargingHasBeeEnded                        = "Charging has been ended"
	ChargingFault                              = "Charging fault"
	UnpausingCharging                          = "Unpausing charging"
	Unavailable                                = "Unavailable"
	UnknownStatus                              = "Unknown"
)

type ConnectorPhases struct {
	L1 int
	L2 int
	L3 int
}

type Connector struct {
	Type               ConnectorType
	Status             ConnectorStatus
	NumberOfPhases     int
	Phases             ConnectorPhases
	MaxCurrentInAmpere int
}

func (c *ChargerClient) ReadNumberOfConnectors() int {
	fmt.Println("### Reading charger connectors...")

	register := c.readRegister(ChargerConnectorsNumber, "Number of connectors", 2)

	return int(register[0])
}

func (c *ChargerClient) ReadConnector(number int) Connector {
	return Connector{
		Type:               c.readConnectorType(number),
		Status:             c.readConnectorStatus(number),
		NumberOfPhases:     c.readNumberOfPhases(number),
		Phases:             c.readConnectorPhases(number),
		MaxCurrentInAmpere: c.readConnectorMaxInA(number),
	}
}

func (c *ChargerClient) readConnectorType(connectorNumber int) ConnectorType {
	registerAddress := uint16(ChargerConnectorBaseAddress + getConnectorOffset(connectorNumber))
	register := c.readBytes(registerAddress, "Connector number "+strconv.Itoa(connectorNumber), 1, true)

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
	registerAddress := uint16(ChargerConnectorPhaseAddress + getConnectorOffset(connectorNumber))
	register := c.readRegister(registerAddress, "Number of phases of connector "+strconv.Itoa(connectorNumber), 1)

	return int(register[0])
}

func (c *ChargerClient) readConnectorPhases(connectorNumber int) ConnectorPhases {
	registerAddressL1 := uint16(ChargerConnectorL1Address + getConnectorOffset(connectorNumber))
	registerL1 := c.readRegister(registerAddressL1, "L1 connected phase of connector "+strconv.Itoa(connectorNumber), 1)

	registerAddressL2 := uint16(ChargerConnectorL2Address + getConnectorOffset(connectorNumber))
	registerL2 := c.readRegister(registerAddressL2, "L2 connected phase of connector "+strconv.Itoa(connectorNumber), 1)

	registerAddressL3 := uint16(ChargerConnectorL3Address + getConnectorOffset(connectorNumber))
	registerL3 := c.readRegister(registerAddressL3, "L3 connected phase of connector "+strconv.Itoa(connectorNumber), 1)

	return ConnectorPhases{
		L1: int(registerL1[0]),
		L2: int(registerL2[0]),
		L3: int(registerL3[0]),
	}
}

func (c *ChargerClient) readConnectorMaxInA(connectorNumber int) int {
	registerAddress := uint16(ChargerConnectorMaxCurrent + getConnectorOffset(connectorNumber))
	register := c.readBytes(registerAddress, "Max current in ampere of connector "+strconv.Itoa(connectorNumber), 2, false)

	cMaxCurrent := binary.BigEndian.Uint32(register)
	fcMaxCurrent := math.Float32frombits(cMaxCurrent)

	return int(fcMaxCurrent)
}

func (c *ChargerClient) readConnectorStatus(connectorNumber int) ConnectorStatus {
	registerAddress := uint16(ChargerConnectorStatus + getConnectorOffset(connectorNumber))
	register := c.readBytes(registerAddress, "Connector number "+strconv.Itoa(connectorNumber), 1, true)

	var status ConnectorStatus

	if register[0] == 1 {
		status = Available
	} else if register[0] == 2 {
		status = ConnectTheCable
	} else if register[0] == 3 {
		status = WaitingForVehicleToRespond
	} else if register[0] == 4 {
		status = Charging
	} else if register[0] == 5 {
		status = VehicleHasPausedCharging
	} else if register[0] == 6 {
		status = EVSEHasPausedCharging
	} else if register[0] == 7 {
		status = ChargingHasBeeEnded
	} else if register[0] == 8 {
		status = ChargingFault
	} else if register[0] == 9 {
		status = UnpausingCharging
	} else if register[0] == 10 {
		status = Unavailable
	} else {
		status = UnknownStatus
	}

	return status
}

func getConnectorOffset(connectorNumber int) int {
	return (connectorNumber - 1) * 100
}

package client

import (
	"bytes"
	"fmt"
	"github.com/simonvetter/modbus"
	"time"
)

type ChargerClient struct {
	Ip     string
	client *modbus.ModbusClient
}

func NewClient(ip string) *ChargerClient {
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://" + ip + ":502",
		Timeout: 1 * time.Second,
	})

	if err != nil {
		fmt.Println("Generic error", err)
		return nil
	}

	return &ChargerClient{
		Ip:     ip,
		client: client,
	}
}

func (c *ChargerClient) readBytesAsString(registerAddress uint16, registerName string, quantity uint16) string {
	register, err := c.client.ReadBytes(registerAddress, quantity, modbus.INPUT_REGISTER)

	if err != nil {
		fmt.Printf("[%d] %s: failed with error '%s' \n", registerAddress, registerName, err)
		return ""
	}

	// remove 0 bytes
	n := bytes.Index(register[:], []byte{0})
	value := register[:n]

	fmt.Printf("[%d] %s: %s \n", registerAddress, registerName, value)
	return string(value)
}

func (c *ChargerClient) readRegister(registerAddress uint16, registerName string, quantity uint16) []uint16 {
	register, err := c.client.ReadRegisters(registerAddress, quantity, modbus.INPUT_REGISTER)

	if err != nil {
		fmt.Printf("[%d] %s: failed with error '%s' \n", registerAddress, registerName, err)
		return make([]uint16, 0)
	}

	fmt.Printf("[%d] %s: %s \n", registerAddress, registerName, register)
	return register
}

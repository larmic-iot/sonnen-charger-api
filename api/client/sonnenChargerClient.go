package client

import (
	"bytes"
	"fmt"
	modbus2 "github.com/goburrow/modbus"
	"time"
)

type ChargerClient struct {
	Ip            string
	modbusHandler *modbus2.TCPClientHandler
}

func NewClient(ip string) *ChargerClient {
	handler := modbus2.NewTCPClientHandler(ip + ":502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0xFF
	//handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	defer handler.Close()

	if err != nil {
		fmt.Println("Generic error", err)
		return nil
	}

	return &ChargerClient{
		Ip:            ip,
		modbusHandler: handler,
	}
}

func (c *ChargerClient) readBytesAsString(registerAddress uint16, registerName string, quantity uint16) string {
	client := modbus2.NewClient(c.modbusHandler)
	results, err := client.ReadInputRegisters(registerAddress, quantity)

	if err != nil {
		fmt.Printf("[%d] %s: failed with error '%s' \n", registerAddress, registerName, err)
		return ""
	}

	// remove 0 bytes
	n := bytes.Index(results[:], []byte{0})
	value := results[:n]

	fmt.Printf("[%d] %s: %s \n", registerAddress, registerName, value)
	return string(value)
}

func (c *ChargerClient) readBytes(registerAddress uint16, registerName string, quantity uint16) []byte {
	client := modbus2.NewClient(c.modbusHandler)
	results, err := client.ReadInputRegisters(registerAddress, quantity)

	if err != nil {
		fmt.Printf("[%d] %s: failed with error '%s' \n", registerAddress, registerName, err)
		return []byte("")
	}

	fmt.Printf("[%d] %s: %s \n", registerAddress, registerName, results)
	return results
}

func (c *ChargerClient) readRegister(registerAddress uint16, registerName string, quantity uint16) []byte {
	client := modbus2.NewClient(c.modbusHandler)
	results, err := client.ReadInputRegisters(registerAddress, quantity)

	if err != nil {
		fmt.Printf("[%d] %s: failed with error '%s' \n", registerAddress, registerName, err)
		return []byte{0}
	}

	var nonZeroBytes []byte
	for _, value := range results {
		if value != 0 {
			nonZeroBytes = append(nonZeroBytes, value)
		}
	}

	if nonZeroBytes == nil {
		return []byte{0}
	}

	fmt.Printf("[%d] %s: %d \n", registerAddress, registerName, nonZeroBytes)
	return nonZeroBytes
}

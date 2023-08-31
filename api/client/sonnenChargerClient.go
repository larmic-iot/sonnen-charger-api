package client

import (
	"bytes"
	modbus2 "github.com/goburrow/modbus"
	"log"
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
		log.Println("Generic error", err)
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
		log.Printf("[%d] %s: failed with error '%s'", registerAddress, registerName, err)
		return ""
	}

	// remove 0 bytes
	n := bytes.Index(results[:], []byte{0})
	value := results[:n]

	log.Printf("[%d] %s: %s", registerAddress, registerName, results)
	return string(value)
}

func (c *ChargerClient) readBytes(registerAddress uint16, registerName string, quantity uint16, removeZeroByte bool) []byte {
	client := modbus2.NewClient(c.modbusHandler)
	results, err := client.ReadInputRegisters(registerAddress, quantity)

	if err != nil {
		log.Printf("[%d] %s: failed with error '%s'", registerAddress, registerName, err)
		return []byte(nil)
	}

	if removeZeroByte {
		var nonZeroBytes = removeZeroBytes(results)
		log.Printf("[%d] %s: %s", registerAddress, registerName, nonZeroBytes)
		return nonZeroBytes
	}

	log.Printf("[%d] %s: %s", registerAddress, registerName, results)
	return results
}

func (c *ChargerClient) readRegister(registerAddress uint16, registerName string, quantity uint16) []byte {
	client := modbus2.NewClient(c.modbusHandler)
	results, err := client.ReadInputRegisters(registerAddress, quantity)

	if err != nil {
		log.Printf("[%d] %s: failed with error '%s'", registerAddress, registerName, err)
		return []byte{0}
	}

	var nonZeroBytes = removeZeroBytes(results)

	log.Printf("[%d] %s: %s", registerAddress, registerName, nonZeroBytes)
	return nonZeroBytes
}

// remove all zero bytes from array and add at least one zero if array is empty
func removeZeroBytes(value []byte) []byte {
	var nonZeroBytes []byte
	for _, value := range value {
		if value != 0 {
			nonZeroBytes = append(nonZeroBytes, value)
		}
	}

	if nonZeroBytes == nil {
		return []byte{0}
	}

	return nonZeroBytes
}

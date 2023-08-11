package client

import "fmt"

const (
	ChargerSettingsSerialNumber = 990
	ChargerSettingsModel        = 1000
	ChargerSettingsHwVersion    = 1010
	ChargerSettingsSwVersion    = 1015
)

type ChargerSettings struct {
	SerialNumber string
	Model        string
	HWVersion    string
	SWVersion    string
}

func (c *ChargerClient) ReadSettings() ChargerSettings {
	fmt.Println("### Reading charger settings...")
	_ = c.client.Open()

	serialNumber := c.readBytesAsString(ChargerSettingsSerialNumber, "Serial Number", 10)
	model := c.readBytesAsString(ChargerSettingsModel, "Model", 20)
	hwVersion := c.readBytesAsString(ChargerSettingsHwVersion, "HW version", 10)
	swVersion := c.readBytesAsString(ChargerSettingsSwVersion, "SW version", 10)

	_ = c.client.Close()

	return ChargerSettings{
		SerialNumber: serialNumber,
		Model:        model,
		HWVersion:    hwVersion,
		SWVersion:    swVersion,
	}
}

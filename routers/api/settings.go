package api

import (
	"github.com/gin-gonic/gin"
	"larmic/sonnen-charger-api/client"
	"net/http"
)

type chargerSettingsDto struct {
	SerialNumber string `json:"serialNumber"`
	Model        string `json:"model"`
	HWVersion    string `json:"hardwareVersion"`
	SWVersion    string `json:"softwareVersion"`
	Connectors   int    `json:"numberOfConnectors"`
}

func GetSettings(c *gin.Context) {
	charger := client.NewClient("10.0.40.200")

	settings := charger.ReadSettings()
	connectors := charger.ReadNumberOfConnectors()

	c.IndentedJSON(http.StatusOK, chargerSettingsDto{
		SerialNumber: settings.SerialNumber,
		Model:        settings.Model,
		HWVersion:    settings.HWVersion,
		SWVersion:    settings.SWVersion,
		Connectors:   connectors,
	})
}

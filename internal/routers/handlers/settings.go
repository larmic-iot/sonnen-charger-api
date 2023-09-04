package handlers

import (
	"github.com/gin-gonic/gin"
	"larmic/sonnen-charger-api/internal/client"
	"net/http"
)

type chargerSettingsDto struct {
	SerialNumber string `json:"serialNumber"`
	Model        string `json:"model"`
	HWVersion    string `json:"hardwareVersion"`
	SWVersion    string `json:"softwareVersion"`
	Connectors   int    `json:"numberOfConnectors"`
}

func GetSettings(chargerIp string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		charger := client.NewClient(chargerIp)

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
	return gin.HandlerFunc(fn)
}

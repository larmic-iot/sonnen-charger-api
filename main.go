package main

import (
	"fmt"
	"larmic/sonnen-charger-api/api/client"
)

const (
	// real time values
	ADDR_CONNNECTOR_STATUS_BASE                               = 0
	ADDR_MEASURED_VEHICLE_NUMBER_OF_PHASES_BASE               = 1
	ADDR_EV_MAX_PHASE_CURRENT_BASE                            = 2
	ADDR_TARGET_CURRENT_FROM_POWER_MGM_OR_MODBUS_BASE         = 4
	ADDR_FREQUENCY_BASE                                       = 6
	ADDR_L_N_VOLTAGE_L1_BASE                                  = 8
	ADDR_L_N_VOLTAGE_L2_BASE                                  = 10
	ADDR_L_N_VOLTAGE_L3_BASE                                  = 12
	ADDR_CURENT_L1_BASE                                       = 14
	ADDR_CURENT_L2_BASE                                       = 16
	ADDR_CURENT_L3_BASE                                       = 18
	ADDR_ACTIVE_POWER_L1_BASE                                 = 20
	ADDR_ACTIVE_POWER_L2_BASE                                 = 22
	ADDR_ACTIVE_POWER_L3_BASE                                 = 24
	ADDR_ACTIVE_POWER_TOTAL_BASE                              = 26
	ADDR_POWER_FACTOR_BASE                                    = 28
	ADDR_TOTAL_IMPORTED_ACTIVE_ENERGY_IN_RUNNING_SESSION_BASE = 30
	ADDR_RUNNING_SESSION_DURATION_BASE                        = 32
	ADDR_RUNNING_SESSION_DEPARTURE_TIME_BASE                  = 36
	ADDR_RUNNING_SESSION_ID_BASE                              = 40
	ADDR_EV_MAX_POWER_BASE                                    = 44
	ADDR_EV_PLANNED_ENERGY_BASE                               = 46

	AddrNumberOfConnectors = 1020
)

func main() {
	c := client.NewClient("10.0.40.200")

	fmt.Println(c.ReadSettings())
	fmt.Println(c.ReadNumberOfConnectors())
	fmt.Println(c.ReadConnector(1))
}

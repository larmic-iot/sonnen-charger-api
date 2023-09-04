package main

import (
	"larmic/sonnen-charger-api/internal/routers"
	"log"
	"os"
)

func main() {
	log.Println("Hello sonnen-charger-api!")

	ip := os.Getenv("SONNEN_CHARGER_IP")

	if ip == "" {
		log.Fatal("Environment variable SONNEN_CHARGER_IP is not set!")
	}

	routersInit := routers.InitRouter(ip)

	_ = routersInit.Run()
}

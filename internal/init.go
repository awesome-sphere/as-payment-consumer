package internal

import (
	"fmt"

	"github.com/awesome-sphere/as-payment-consumer/utils"
)

var SEATING_SERVICE string
var BOOKING_SERVICE string

func initializeService(service, defaultPort, path string) string {
	HOST := utils.GetenvOr(service+"_HOST", "localhost")
	PORT := utils.GetenvOr(service+"_PORT", defaultPort)
	PATH := utils.GetenvOr(service+"_PATH", path)
	return fmt.Sprintf("http://%s:%s%s", HOST, PORT, PATH)
}

func InitializeInternalServices() {
	SEATING_SERVICE = initializeService("SEATING_SERVICE", "9004", "/seating/update-status")
	BOOKING_SERVICE = initializeService("BOOKING_SERVICE", "9009", "/booking/update-status")
}

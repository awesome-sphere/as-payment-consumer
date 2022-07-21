package internal

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	outputserializer "github.com/awesome-sphere/as-payment-consumer/internal/output-serializer"
	"github.com/awesome-sphere/as-payment-consumer/kafka/interfaces"
)

func notifySeatingService(val interfaces.UpdateOrderMessageInterface, seatID int) {
	body := outputserializer.SeatingServiceSerializer{
		TimeSlotID: int64(val.TimeSlotId),
		TheaterID:  int64(val.TheaterId),
		SeatID:     seatID,
		Status:     val.Status,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err.Error())
		return
	}
	http.Post(SEATING_SERVICE, "application/json", bytes.NewBuffer(jsonBody))
}

func generateTicket(val interfaces.UpdateOrderMessageInterface, seatID int) {
	// TODO: generate ticket
	// TODO: send to authen
}

func notifyBookingService(val interfaces.UpdateOrderMessageInterface, seatID int) {
	// TODO: wait for booking service to be ready
}

func NotifyOtherServices(val interfaces.UpdateOrderMessageInterface) {
	for _, seatID := range val.SeatNumber {
		notifySeatingService(val, seatID)
		notifyBookingService(val, seatID)
		generateTicket(val, seatID)
	}
}

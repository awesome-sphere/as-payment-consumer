package db

import (
	"log"

	"github.com/awesome-sphere/as-payment-consumer/db/models"
)

func CreateUserHistory(user_id int, time_slot_id int, theater_id int, seat_number []int, price int) {
	event := models.Order{
		UserID:     int64(user_id),
		TimeSlotID: int64(time_slot_id),
		TheaterID:  int64(theater_id),
		Price:      int64(price),
		Status:     models.Awaiting,
	}
	err := DB.Create(&event).Error
	if err != nil {
		log.Printf("Failed to update user history: %v", err.Error())
		return
	} else {
		for _, elt := range seat_number {
			history := models.OrderSeats{SeatID: int64(elt), Order: event, OrderID: event.ID}
			err := DB.Create(&history).Error
			if err != nil {
				log.Printf("Failed to update booking history: %v", err.Error())
				return
			}
		}
		log.Printf("Successfully updating %d's purchase history", user_id)
	}
}

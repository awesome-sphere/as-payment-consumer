package db

import (
	"log"

	"github.com/awesome-sphere/as-payment-consumer/db/models"
)

type orderStruct struct {
	OrderID int64  `json:"order_id"`
	UserID  int64  `json:"user_id"`
	Status  string `json:"status"`
}

func UpdateUserHistory(UserID, TimeSlotID, TheaterID, SeatID int, status models.OrderStatus) bool {
	var order orderStruct

	tx := DB.Model(&models.Order{}).
		Select("orders.id, orders.user_id, orders.status, orders.time_slot_id, orders.theater_id, order_seats.order_id, order_seats.seat_id").
		Where("time_slot_id = ? AND theater_id = ? AND status = ?", TimeSlotID, TheaterID, models.Awaiting).
		Joins("JOIN order_seats ON order_seats.order_id = orders.id").
		Where("order_seats.seat_id = ?", SeatID).
		First(&models.Order{}).
		Scan(&order)
	if tx.Error != nil {
		log.Printf("Failed to update user history: %v\n", tx.Error.Error())
		return false
	}

	if int(order.UserID) != UserID {
		log.Printf("This ticket doesn't belong to this user\n")
		return false
	}

	tx = DB.Model(&models.Order{}).Where("id = ?", order.OrderID).Update("status", status)
	if tx.Error != nil {
		log.Printf("Failed to update user history: %v\n", tx.Error.Error())
		return false
	}
	return status == models.Paid
}

package models

type OrderStatus string

const (
	Paid     OrderStatus = "PAID"
	Awaiting OrderStatus = "AWAITING"
	Canceled OrderStatus = "CANCELED"
)

type Order struct {
	ID         int64       `json:"id" gorm:"primaryKey;autoincrement;not null"`
	UserID     int64       `json:"user_id" gorm:"not null"`
	TimeSlotID int64       `json:"time_slot_id" gorm:"not null"`
	TheaterID  int64       `json:"theater_id" gorm:"not null"`
	Price      int64       `json:"price"`
	Status     OrderStatus `json:"order_status" sql:"type:order_status"`
}

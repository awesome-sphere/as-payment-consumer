package models

type OrderSeats struct {
	ID      int64 `json:"id" gorm:"primaryKey;autoincrement;not null"`
	SeatID  int64 `json:"seat_id" gorm:"not null"`
	Order   Order `gorm:"foreignKey:OrderID"`
	OrderID int64 `json:"order_id" gorm:"not null"`
}

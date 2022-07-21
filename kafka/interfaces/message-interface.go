package interfaces

type CreateOrderMessageInterface struct {
	UserID     int   `json:"user_id"`
	TimeSlotId int   `json:"time_slot_id"`
	TheaterId  int   `json:"theater_id"`
	SeatNumber []int `json:"seat_number"`
	Price      int   `json:"price"`
}

type UpdateOrderMessageInterface struct {
	UserID     int    `json:"user_id"`
	TimeSlotId int    `json:"time_slot_id"`
	TheaterId  int    `json:"theater_id"`
	SeatNumber []int  `json:"seat_number"`
	Status     string `json:"status"`
}

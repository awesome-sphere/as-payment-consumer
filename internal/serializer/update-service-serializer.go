package serializer

type SeatingServiceSerializer struct {
	TimeSlotID int64  `json:"time_slot_id" binding:"required"`
	TheaterID  int64  `json:"theater_id" binding:"required"`
	SeatID     int    `json:"seat_id" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

type BookingServiceSerializer struct {
	TimeSlotID int64 `json:"time_slot_id" binding:"required"`
	TheaterID  int64 `json:"theater_id" binding:"required"`
	SeatID     []int `json:"seat_id" binding:"required"`
}

package outputserializer

type SeatingServiceSerializer struct {
	TimeSlotID int64  `json:"time_slot_id"`
	TheaterID  int64  `json:"theater_id"`
	SeatID     int    `json:"seat_id"`
	Status     string `json:"status"`
}

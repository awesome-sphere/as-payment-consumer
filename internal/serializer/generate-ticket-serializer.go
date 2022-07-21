package serializer

type GenerateTicketSerializer struct {
	UserID    int64  `json:"user_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Location  string `json:"location" binding:"required"`
	Duration  int    `json:"duration" binding:"required"`
	Date      string `json:"date" binding:"required"`
	Time      string `json:"time" binding:"required"`
	SeatID    string `json:"seat_id" binding:"required"`
	TheaterID int    `json:"theater_id" binding:"required"`
}

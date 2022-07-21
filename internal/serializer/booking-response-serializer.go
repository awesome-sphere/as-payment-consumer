package serializer

import "time"

type BookingResponseSerializer struct {
	Location string    `json:"location" binding:"required"`
	Date     time.Time `json:"time" binding:"required"`
	MovieID  int64     `json:"movie_id" binding:"required"`
}

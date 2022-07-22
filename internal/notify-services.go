package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/awesome-sphere/as-payment-consumer/internal/serializer"
	"github.com/awesome-sphere/as-payment-consumer/kafka/interfaces"
)

func notifySeatingService(val interfaces.UpdateOrderMessageInterface, seatID int) {
	body := serializer.SeatingServiceSerializer{
		TimeSlotID: int64(val.TimeSlotId),
		TheaterID:  int64(val.TheaterId),
		SeatID:     seatID,
		Status:     val.Status,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err.Error())
	}
	http.Post(SEATING_SERVICE, "application/json", bytes.NewBuffer(jsonBody))
}

func notifyBookingService(val interfaces.UpdateOrderMessageInterface) (*http.Response, []byte) {
	body := serializer.BookingServiceSerializer{
		TimeSlotID: int64(val.TimeSlotId),
		TheaterID:  int64(val.TheaterId),
		SeatID:     val.SeatNumber,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err.Error())
	}
	resp, _ := http.Post(BOOKING_SERVICE, "application/json", bytes.NewBuffer(jsonBody))
	return resp, jsonBody
}

func decodeResponse(resp *http.Response, jsonBody []byte, target interface{}, service string) error {
	if resp == nil || resp.StatusCode != 200 {
		log.Printf("Failed to notify %s service: %v", service, resp.StatusCode)
		return fmt.Errorf("Failed to notify %s service: %v", service, resp.StatusCode)
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func getDate(time time.Time) string {
	year, month, day := time.Date()
	return fmt.Sprintf("%d/%v/%d", day, month, year)
}

func getTime(time time.Time) string {
	return fmt.Sprintf("%d:%d", time.Hour(), time.Minute())
}

func NotifyOtherServices(val interfaces.UpdateOrderMessageInterface) {
	bookingResp := serializer.BookingResponseSerializer{}
	resp, jsonBody := notifyBookingService(val)
	err := decodeResponse(resp, jsonBody, &bookingResp, "booking")
	if err != nil {
		return
	}

	movieResp := serializer.MovieResponseSerializer{}
	resp, _ = http.Get(fmt.Sprintf("%s/%d", MOVIE_SERVICE, bookingResp.MovieID))
	err = decodeResponse(resp, jsonBody, &movieResp, "movie")
	if err != nil {
		print(err.Error())
		return
	}
	seatNumbers := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(val.SeatNumber)), ", "), "[]")
	body := serializer.GenerateTicketSerializer{
		UserID:   int64(val.UserID),
		Title:    movieResp.Movie.Title,
		Location: bookingResp.Location,
		Duration: movieResp.Movie.Duration,
		// FIXME: Find a proper way to format date time
		Date:      getDate(bookingResp.Date),
		Time:      getTime(bookingResp.Date),
		SeatID:    seatNumbers,
		TheaterID: int(val.TheaterId),
	}

	jsonBody, err = json.Marshal(body)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err.Error())
	}

	http.Post(AUTHEN_SERVICE, "application/json", bytes.NewBuffer(jsonBody))

	for _, seatID := range val.SeatNumber {
		notifySeatingService(val, seatID)
	}
}

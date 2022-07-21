package internal

import (
	"github.com/awesome-sphere/as-payment-consumer/utils"
)

var SEATING_SERVICE string
var BOOKING_SERVICE string
var AUTHEN_SERVICE string
var MOVIE_SERVICE string

func InitializeInternalServices() {
	SEATING_SERVICE = utils.GetenvOr("SEATING_SERVICE", "http://localhost:9004/seating/update-status")
	BOOKING_SERVICE = utils.GetenvOr("SEATING_SERVICE", "http://localhost:9009/booking/buy-seat")
	AUTHEN_SERVICE = utils.GetenvOr("AUTHEN_SERVICE", "http://localhost:9001/authen/update-hist")
	MOVIE_SERVICE = utils.GetenvOr("MOVIE_SERVICE", "http://localhost:9002/movie/get-movie")
}

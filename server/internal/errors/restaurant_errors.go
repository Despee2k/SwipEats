package errors

import "errors"

var (
	ErrNoRestaurantsFound       = errors.New("no restaurants found")
	ErrGroupRestaurantNotFound  = errors.New("group restaurant not found")
)
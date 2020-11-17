package ordertime

import "time"

// OrderTimer is the interface for order timer
type OrderTimer interface {
	GetNextWeekInterval() (time.Time, time.Time)
}

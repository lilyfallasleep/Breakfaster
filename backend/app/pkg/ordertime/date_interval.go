package ordertime

import (
	"time"
)

// OrderTimer is the timer type for getting order date interval
type OrderTimer struct{}

func (svc *OrderTimer) getNextMonday() time.Time {
	t := time.Now()
	cur := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	if cur.Weekday() == time.Monday {
		cur = cur.AddDate(0, 0, 1)
	}
	for cur.Weekday() != time.Monday {
		cur = cur.AddDate(0, 0, 1)
	}
	return cur
}

// GetNextWeekInterval service returns the next week time interval
func (svc *OrderTimer) GetNextWeekInterval() (time.Time, time.Time) {
	start := svc.getNextMonday()
	end := start.AddDate(0, 0, 4)
	return start, end
}

// NewOrderTimer is the factory for OrderTimer instance
func NewOrderTimer() *OrderTimer {
	return &OrderTimer{}
}

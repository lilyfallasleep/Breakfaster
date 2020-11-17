package ordertime

import (
	"time"
)

// OrderTimerImpl is the timer type for getting order date interval
type OrderTimerImpl struct{}

func getNextMonday() time.Time {
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
func (svc *OrderTimerImpl) GetNextWeekInterval() (time.Time, time.Time) {
	start := getNextMonday()
	end := start.AddDate(0, 0, 4)
	return start, end
}

// NewOrderTimer is the factory for OrderTimerImpl
func NewOrderTimer() OrderTimer {
	return &OrderTimerImpl{}
}

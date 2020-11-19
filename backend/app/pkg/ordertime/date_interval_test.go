package ordertime

import (
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
)

func Test_getNextMonday(t *testing.T) {
	tests := []struct {
		name    string
		curDate time.Time
		want    time.Time
	}{
		{
			name:    "Get next Monday when today is monday",
			curDate: time.Date(2020, 11, 2, 0, 0, 0, 0, time.Local),
			want:    time.Date(2020, 11, 9, 0, 0, 0, 0, time.Local),
		},
		{
			name:    "Get next Monday when today is not monday",
			curDate: time.Date(2020, 11, 3, 0, 0, 0, 0, time.Local),
			want:    time.Date(2020, 11, 9, 0, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monkey.Patch(time.Now, func() time.Time {
				return tt.curDate
			})
			if got := getNextMonday(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNextMonday() = %v, want %v", got, tt.want)
			}
			monkey.Unpatch(time.Now)
		})
	}
}

func TestOrderTimerImpl_GetNextWeekInterval(t *testing.T) {
	tests := []struct {
		name  string
		svc   *OrderTimerImpl
		want  time.Time
		want1 time.Time
	}{
		{
			name:  "Get next week date interval",
			svc:   &OrderTimerImpl{},
			want:  getNextMonday(),
			want1: getNextMonday().AddDate(0, 0, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.svc.GetNextWeekInterval()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderTimerImpl.GetNextWeekInterval() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OrderTimerImpl.GetNextWeekInterval() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewOrderTimer(t *testing.T) {
	tests := []struct {
		name string
		want OrderTimer
	}{
		{
			name: "A new order timer",
			want: &OrderTimerImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderTimer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderTimer() = %v, want %v", got, tt.want)
			}
		})
	}
}

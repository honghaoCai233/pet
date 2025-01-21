package xtime

import (
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"time"
)

const (
	ISO8601Format = "2006-01-02T15:04:05Z"
)
const (
	UnixSecond int64 = 1
	UnixMinute       = 60 * UnixSecond
	UnixHour         = 60 * UnixMinute
	UnixDay          = 24 * UnixHour
	UnixWeek         = 7 * UnixDay
)

type PeriodTimeStamp struct {
	Start int64
	End   int64
}

type PeriodTime struct {
	Start time.Time
	End   time.Time
}

func FormatGMTISO8601(t time.Time) string {
	return t.UTC().Format(ISO8601Format)
}

func FormatSecondsDuration(t float64) string {
	seconds := int(t)
	ms := int(t*1000) % 1000
	sec := seconds % 60
	seconds = seconds / 60
	m := seconds % 60
	h := seconds / 60
	return fmt.Sprintf("%02d:%02d:%02d.%03d", h, m, sec, ms)
}

func StartOfHour() time.Time {
	return StartOfLastNHour(0)
}

func StartOfLastHour() time.Time {
	return StartOfLastNHour(1)
}

func StartOfLastNHour(n uint) time.Time {
	return datetime.BeginOfHour(time.Now().Add(-time.Duration(n) * time.Hour))
}

func StartOfMinute() time.Time {
	return StartOfLastNMinute(0)
}

func StartOfLastMinute() time.Time {
	return StartOfLastNMinute(1)
}

func StartOfLastNMinute(n int) time.Time {
	return datetime.BeginOfMinute(time.Now().Add(-time.Duration(n) * time.Minute))
}

func TimePtrToUnix(t *time.Time) int64 {
	if t == nil {
		return 0
	}
	return t.Unix()
}

func IsSameDay(t1, t2 time.Time) bool {
	return datetime.BeginOfDay(t1).Equal(datetime.BeginOfDay(t2))
}

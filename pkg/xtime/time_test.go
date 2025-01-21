package xtime

import "testing"

func TestStartOfLastHour(t *testing.T) {
	hour := StartOfLastHour()
	t.Log(hour.String())
	t.Log(hour.Format("2006-01-02 15"))
}

func TestStartOfLastNHour(t *testing.T) {
	hour := StartOfLastNHour(4*24 + 8)
	t.Log(hour.String())
	t.Log(hour.Format("2006-01-02 15"))
}

func TestStartOfLastMinute(t *testing.T) {
	minute := StartOfLastMinute()
	t.Log(minute)
}

func TestStartOfLastNMinute(t *testing.T) {
	minute := StartOfLastNMinute(5)
	t.Log(minute)
}

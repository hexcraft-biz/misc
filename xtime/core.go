package xtime

import (
	"database/sql/driver"
	"errors"
	"time"
)

//================================================================
// Time
//================================================================
type Time struct {
	time.Time
}

func NowUTC() Time {
	t := time.Now().UTC()
	return Time{Time: t}
}

func (t Time) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(time.RFC3339)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, time.RFC3339)
	b = append(b, '"')
	return b, nil
}

func (t Time) Value() (driver.Value, error) {
	return t.Time, nil
}

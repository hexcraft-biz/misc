package xtime

import (
	"database/sql/driver"
	"time"
)

// ================================================================
// Time
// ================================================================
type Time time.Time

func NowUTC() Time {
	return Time(time.Now().UTC())
}

func MysqlMin() Time {
	return Time(time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC))
}

func MysqlMax() Time {
	return Time(time.Date(9999, time.December, 31, 23, 59, 59, 999999999, time.UTC))
}

func (t Time) Before(u Time) bool {
	return time.Time(t).Before(time.Time(u))
}

func (t Time) After(u Time) bool {
	return time.Time(t).After(time.Time(u))
}

func (t Time) Add(d time.Duration) Time {
	return Time(time.Time(t).Add(d))
}

func (t Time) MarshalJSON() ([]byte, error) {

	/*
		tt := time.Time(t)
		stamp := []byte(fmt.Sprintf("%q", tt.Format(time.RFC3339)))
		return stamp, nil
	*/

	return time.Time(t).MarshalJSON()

	/*
		if y := time.Time(t).Year(); y < 0 || y >= 10000 {
			return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
		}

		b := make([]byte, 0, len(time.RFC3339)+2)
		b = append(b, '"')
		b = time.Time(t).AppendFormat(b, time.RFC3339)
		b = append(b, '"')
		return b, nil
	*/
}

func (t *Time) UnmarshalJSON(data []byte) error {
	return (*time.Time)(t).UnmarshalJSON(data)
}

func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

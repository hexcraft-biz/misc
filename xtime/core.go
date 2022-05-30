package xtime

import (
	"time"
)

type TimeRFC3339 time.Time

func (ft *TimeRFC3339) UnmarshalJSON(data []byte) error {
	if t, err := time.ParseInLocation(time.RFC3339, string(data[:]), time.UTC); err != nil {
		return err
	} else {
		*ft = TimeRFC3339(t)
		return nil
	}
}

func (ft TimeRFC3339) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(ft).UTC().Truncate(time.Second).Format(time.RFC3339)), nil
}

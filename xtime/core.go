package xtime

import (
	"encoding/json"
	"time"
)

type TimeRFC3339 time.Time

func (ft *TimeRFC3339) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	} else if t, err := time.ParseInLocation(time.RFC3339, s, time.UTC); err != nil {
		return err
	} else {
		*ft = TimeRFC3339(t)
		return nil
	}
}

func (ft TimeRFC3339) String() string {
	return time.Time(ft).String()
}

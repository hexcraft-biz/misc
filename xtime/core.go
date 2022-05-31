package xtime

import (
	"encoding/json"
	"fmt"
	"time"
)

type TimeRFC3339 time.Time

func (ft *TimeRFC3339) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data[:], &s); err != nil {
		return err
	} else if t, err := time.ParseInLocation(time.RFC3339, s, time.UTC); err != nil {
		return err
	} else {
		*ft = TimeRFC3339(t)
		return nil
	}
}

func (ft TimeRFC3339) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(ft).UTC().Truncate(time.Second).Format(time.RFC3339))), nil
}

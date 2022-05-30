package xtime

import (
	"time"
)

type RFC3339Time time.Time

func (ft *RFC3339Time) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	} else if t, err := time.ParseInLocation(time.RFC3339, s, time.UTC); err != nil {
		return err
	} else {
		*ft = RFC3339Time(t)
		return nil
	}
}

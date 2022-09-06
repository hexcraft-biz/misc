package xtime

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

//================================================================
//
//================================================================
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

func (ft TimeRFC3339) Value() (driver.Value, error) {
	return time.Time(ft).Format("2006-01-02T15:04:05Z"), nil
}

//================================================================
//
//================================================================
type NullTimeRFC3339 sql.NullTime

func (ntf *NullTimeRFC3339) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data[:], &s); err != nil {
		return err
	} else if s == "" {
		return nil
	} else if t, err := time.ParseInLocation(time.RFC3339, s, time.UTC); err != nil {
		return err
	} else {
		*ntf = NullTimeRFC3339(sql.NullTime{Valid: true, Time: t})
		return nil
	}
}

func (ntf NullTimeRFC3339) MarshalJSON() ([]byte, error) {
	if ntf.Valid {
		return []byte(fmt.Sprintf("\"%s\"", ntf.Time.UTC().Truncate(time.Second).Format(time.RFC3339))), nil
	} else {
		return []byte("\"\""), nil
	}
}

func (ntf NullTimeRFC3339) Value() (driver.Value, error) {
	return sql.NullTime(ntf).Value()
}

func (ntf *NullTimeRFC3339) Scan(value interface{}) error {
	return (*sql.NullTime)(ntf).Scan(value)
}

func ValidateTypeNullTimeRFC3339(field reflect.Value) interface{} {
	if field.Type() == reflect.TypeOf(NullTimeRFC3339{}) {
		nilRef := time.Time{}
		if field.Interface().(NullTimeRFC3339).Time == nilRef {
			return nil
		} else if val, err := field.Interface().(NullTimeRFC3339).Value(); err != nil {
			return nil
		} else {
			return val
		}
	}

	return nil
}

func ValidateTypeTimeRFC3339(field reflect.Value) interface{} {
	if field.Type() == reflect.TypeOf(TimeRFC3339{}) {
		nilRef := time.Time{}
		if field.Interface().(time.Time) == nilRef {
			return nil
		} else if val, err := field.Interface().(TimeRFC3339).Value(); err != nil {
			return nil
		} else {
			return val
		}
	}

	return nil
}

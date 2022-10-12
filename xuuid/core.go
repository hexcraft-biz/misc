package xuuid

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

//================================================================
// UUID
//================================================================
type UUID uuid.UUID

func (xuuid *UUID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data[:], &s); err != nil {
		return err
	} else if s == "" {
		return nil
	} else if u, err := uuid.Parse(s); err != nil {
		return err
	} else {
		*xuuid = UUID(u)
		return nil
	}
}

func (xuuid *UUID) Scan(src interface{}) error {
	return (*uuid.UUID)(xuuid).Scan(src)
}

func (xuuid UUID) Value() (driver.Value, error) {
	return uuid.UUID(xuuid).MarshalBinary()
}

func (xuuid UUID) MarshalBinary() ([]byte, error) {
	return uuid.UUID(xuuid).MarshalBinary()
}

func (xuuid *UUID) UnmarshalBinary(data []byte) error {
	return (*uuid.UUID)(xuuid).UnmarshalBinary(data)
}

func (xuuid *UUID) UnmarshalText(data []byte) error {
	return (*uuid.UUID)(xuuid).UnmarshalText(data)
}

func (xuuid UUID) MarshalText() ([]byte, error) {
	return uuid.UUID(xuuid).MarshalText()
}

func (xuuid UUID) String() string {
	return uuid.UUID(xuuid).String()
}

//================================================================
// Wildcard
//================================================================
type WildcardType uint8

const (
	WildcardTypeUndef WildcardType = iota
	WildcardTypeString
	WildcardTypeXUUID
)

type Wildcard struct {
	Type WildcardType
	Val  interface{}
}

func (w *Wildcard) UnmarshalJSON(data []byte) error {
	w.Type = WildcardTypeUndef
	if len(data) <= 0 {
		return nil
	}

	var s string
	if err := json.Unmarshal(data[:], &s); err != nil {
		return err
	}

	if u, err := uuid.Parse(s); err == nil {
		w.Type = WildcardTypeXUUID
		w.Val = UUID(u)
	} else {
		w.Type = WildcardTypeString
		w.Val = s
	}

	return nil
}

func (w *Wildcard) Scan(src interface{}) error {
	w.Type = WildcardTypeUndef
	switch src := src.(type) {
	case nil:
		return nil

	case string:
		if src == "" {
			w.Type = WildcardTypeString
			return nil
		}

		if u, err := uuid.Parse(src); err == nil {
			w.Type = WildcardTypeXUUID
			w.Val = UUID(u)
		} else {
			w.Type = WildcardTypeString
			w.Val = src
		}

	case []byte:
		if len(src) == 0 {
			return nil
		}

		if len(src) != 16 {
			return w.Scan(string(src))
		}

		w.Type = WildcardTypeXUUID
		u := uuid.UUID{}
		copy((u)[:], src)
		w.Val = UUID(u)

	default:
		return fmt.Errorf("Wildcard: unable to scan type %T", src)
	}

	return nil
}

func (w Wildcard) Value() (driver.Value, error) {
	switch w.Type {
	case WildcardTypeString:
		return w.Val.(string), nil
	case WildcardTypeXUUID:
		return w.Val.(UUID).Value()
	default:
		return nil, fmt.Errorf("Wildcard: Invalid value %v.", w.Val)
	}
}

func (w Wildcard) MarshalBinary() ([]byte, error) {
	switch w.Type {
	case WildcardTypeString:
		return json.Marshal(w.Val.(string))
	case WildcardTypeXUUID:
		return w.Val.(UUID).MarshalBinary()
	default:
		return nil, fmt.Errorf("Wildcard: Failed to MarshalBinary on %v.", w.Val)
	}
}

func (w *Wildcard) UnmarshalBinary(data []byte) error {
	if len(data) <= 0 {
		return nil
	}

	var (
		u uuid.UUID
		s string
	)

	if err := u.UnmarshalBinary(data); err == nil {
		w.Type = WildcardTypeXUUID
		w.Val = UUID(u)
		return nil
	}

	if err := json.Unmarshal(data, &s); err == nil {
		w.Type = WildcardTypeString
		w.Val = s
		return nil
	}

	return fmt.Errorf("Wildcard: Failed to UnmarshalBinary.")
}

func (w *Wildcard) UnmarshalText(data []byte) error {
	var u uuid.UUID

	if err := u.UnmarshalText(data); err == nil {
		w.Type = WildcardTypeXUUID
		w.Val = UUID(u)
	} else {
		w.Type = WildcardTypeString
		w.Val = string(data[:])
	}

	return nil
}

func (w Wildcard) MarshalText() ([]byte, error) {
	switch w.Type {
	case WildcardTypeString:
		return []byte(w.Val.(string)), nil
	case WildcardTypeXUUID:
		return w.Val.(UUID).MarshalText()
	default:
		return nil, fmt.Errorf("Wildcard: Failed to MarshalText on %v.", w.Val)
	}
}

func (w Wildcard) String() string {
	switch w.Type {
	case WildcardTypeString:
		return w.Val.(string)
	case WildcardTypeXUUID:
		return w.Val.(UUID).String()
	default:
		return ""
	}
}

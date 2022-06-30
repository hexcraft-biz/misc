package xuuid

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/google/uuid"
)

type UUID uuid.UUID

func (xuuid *UUID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data[:], &s); err != nil {
		return err
	} else if u, err := uuid.Parse(s); err != nil {
		return err
	} else {
		*xuuid = UUID(u)
		return nil
	}
}

func (xuuid UUID) MarshalJSON() ([]byte, error) {
	return xuuid.MarshalBinary()
}

func (xuuid *UUID) Scan(src interface{}) error {
	return (*uuid.UUID)(xuuid).Scan(src)
}

func (xuuid UUID) Value() (driver.Value, error) {
	return uuid.UUID(xuuid).Value()
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

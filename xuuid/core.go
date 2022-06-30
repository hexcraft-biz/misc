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
	return []byte(uuid.UUID(xuuid).String()), nil
}

func (xuuid *UUID) Scan(src interface{}) error {
	u := (*uuid.UUID)(xuuid)
	return u.Scan(src)
}

func (xuuid UUID) Value() (driver.Value, error) {
	u := uuid.UUID(xuuid)
	return u.Value()
}

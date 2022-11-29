package xuuid

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/google/uuid"
)

// ================================================================
// UUID
// ================================================================
type UUID uuid.UUID

func New() UUID {
	return UUID(uuid.New())
}

func (xu *UUID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data[:], &s); err != nil {
		return err
	} else if s == "" {
		return nil
	} else if u, err := uuid.Parse(s); err != nil {
		return err
	} else {
		*xu = UUID(u)
		return nil
	}
}

func (xu *UUID) UnmarshalBinary(data []byte) error {
	return (*uuid.UUID)(xu).UnmarshalBinary(data)
}

func (xu UUID) MarshalBinary() ([]byte, error) {
	return uuid.UUID(xu).MarshalBinary()
}

func (xu *UUID) UnmarshalText(data []byte) error {
	return (*uuid.UUID)(xu).UnmarshalText(data)
}

func (xu UUID) MarshalText() ([]byte, error) {
	return uuid.UUID(xu).MarshalText()
}

func (xu UUID) String() string {
	return uuid.UUID(xu).String()
}

func (xu *UUID) Scan(src interface{}) error {
	return (*uuid.UUID)(xu).Scan(src)
}

func (xu UUID) Value() (driver.Value, error) {
	return xu.MarshalBinary()
}

// ================================================================
// Wildcard
// ================================================================
type Wildcard []byte

func (w *Wildcard) UnmarshalJSON(data []byte) error {
	u, err := uuid.Parse(string(data))
	if err != nil {
		// string
		*w = data
		return nil
	}

	if b, err := u.MarshalBinary(); err != nil {
		return err
	} else {
		// uuid
		*w = b
		return nil
	}
}

func (w *Wildcard) UnmarshalBinary(data []byte) error {
	u, err := uuid.FromBytes(data)
	if err != nil {
		// string
		*w = data
		return nil
	}

	if b, err := u.MarshalBinary(); err != nil {
		return err
	} else {
		// uuid
		*w = b
		return nil
	}
}

func (w Wildcard) MarshalBinary() ([]byte, error) {
	if u, err := uuid.FromBytes(w); err != nil {
		// string
		return w, nil
	} else {
		// uuid
		return u.MarshalBinary()
	}
}

func (w *Wildcard) UnmarshalText(data []byte) error {
	u, err := uuid.Parse(string(data))
	if err != nil {
		// string
		*w = data
		return nil
	}

	if b, err := u.MarshalBinary(); err != nil {
		return err
	} else {
		// uuid
		*w = b
		return nil
	}
}

func (w Wildcard) MarshalText() ([]byte, error) {
	if u, err := uuid.FromBytes(w); err != nil {
		// string
		return w, nil
	} else {
		// uuid
		return u.MarshalText()
	}
}

func (w Wildcard) String() string {
	if u, err := uuid.FromBytes(w); err != nil {
		// string
		return string(w)
	} else {
		// uuid
		return u.String()
	}
}

func (w Wildcard) Value() (driver.Value, error) {
	if u, err := uuid.FromBytes(w); err != nil {
		// string
		return string(w), nil
	} else {
		// uuid
		return UUID(u).Value()
	}
}

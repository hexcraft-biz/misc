package xuuid

import (
	"encoding/json"
	"github.com/google/uuid"
)

type UUID uuid.UUID

func (fu *UUID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data[:], &s); err != nil {
		return err
	} else if u, err := uuid.Parse(s); err != nil {
		return err
	} else {
		*fu = UUID(u)
		return nil
	}
}

func (fu UUID) MarshalJSON() ([]byte, error) {
	return []byte(uuid.UUID(fu).String()), nil
}

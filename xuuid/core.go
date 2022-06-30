package xuuid

import (
	"encoding/json"
	"github.com/google/uuid"
)

type XUUID uuid.UUID

func (fu *XUUID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data[:], &s); err != nil {
		return err
	} else if u, err := uuid.Parse(s); err != nil {
		return err
	} else {
		*fu = XUUID(u)
		return nil
	}
}

func (fu XUUID) MarshalJSON() ([]byte, error) {
	return []byte(uuid.UUID(fu).String()), nil
}

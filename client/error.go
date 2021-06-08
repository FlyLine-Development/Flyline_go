package client

import (
	"fmt"
)

type Error struct {
	DisplayMessage string `json:"display_message"`

	StatusCode int
}

func (e Error) Error() string {
	return fmt.Sprintf("Plaid Error - http status: %d", e.StatusCode)
}

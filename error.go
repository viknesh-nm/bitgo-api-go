package bitgo

import (
	"encoding/json"
)

// Error holds the error field listed from the server
type Error struct {
	ErrMessage string `json:"error"`
	RequestID  string `json:"requestId"`
	Name       string `json:"name"`
}

// Error returns a string representing the error, satisfying the error interface.
func (e Error) Error() string {
	data, err := json.Marshal(e)
	if err != nil {
		return "Unable to Marshal"
	}
	return string(data)
}

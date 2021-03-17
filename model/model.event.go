package model

import "encoding/json"

type (
	EventData struct {
		UserID string      `json:"user_id"`
		Name   string      `json:"name"`
		Data   interface{} `json:"data"`
	}
)

func (_ *EventData) FromJson(b []byte) EventData {
	var e EventData
	if err := json.Unmarshal(b, &e); err != nil {
		return e
	}
	return e
}

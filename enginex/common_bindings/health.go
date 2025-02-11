package common_bindings

import (
	"bytes"
	"fmt"
)

type HealthResponse struct {
	Status        string         `json:"status"`
	StatusMessage string         `json:"status_message"`
	Items         [] *HealthItem `json:"items"`
	EpochTime     int64          `json:"epoch_time"`
	StatusTime    string         `json:"status_time"`
}

type HealthItem struct {
	ItemName string `json:"item_name"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

func (h HealthResponse) String() string {
	buffer := &bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("Status: %s, Status: %s\n", h.Status, h.StatusMessage))
	for _, item := range h.Items {
		buffer.WriteString(item.String() + "\n")
	}
	return buffer.String()
}

func (hi HealthItem) String() string {
	return fmt.Sprintf("Item Name: %s, Status: %s, Message: %s", hi.ItemName, hi.Status, hi.Message)
}

package entity

import "time"

type Device struct {
	ID        string    `json:"id"`
	Key       string    `json:"key"`
	UpdatedAt time.Time `json:"updatedAt"`
	// Token     string    `json:"token"`
}

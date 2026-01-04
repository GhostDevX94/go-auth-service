package model

import "time"

type RefreshToken struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	IsRevoked bool      `json:"is_revoked"`
}

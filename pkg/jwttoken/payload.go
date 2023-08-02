package jwttoken

import (
	"github.com/google/uuid"
	"time"
)

type JWTPayload struct {
	ID        uuid.UUID `yaml:"id"`
	UserID    int64     `yaml:"user_id"`
	IssuedAt  time.Time `yaml:"issued_at"`
	ExpiredAt time.Time `yaml:"expired_at"`
}

func (j *JWTPayload) Valid() error {
	if time.Now().After(j.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}

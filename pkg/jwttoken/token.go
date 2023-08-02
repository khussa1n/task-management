package jwttoken

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

type JWTToken struct {
	secretkey string
}

func New(secretkey string) *JWTToken {
	return &JWTToken{
		secretkey: secretkey,
	}
}

func (j *JWTToken) CreatToken(userID int64, duration time.Duration) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("uuid newrandom error: %w", err)
	}

	payload := &JWTPayload{
		ID:        id,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(j.secretkey))
}

func (j *JWTToken) ValidateToken(token string) (*JWTPayload, error) {
	keyFanc := func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if ok {
			return []byte(j.secretkey), nil
		}

		return nil, ErrInvalidToken
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JWTPayload{}, keyFanc)
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(validationErr, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*JWTPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

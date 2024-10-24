package tokenutil

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/onattech/invest/models"
)

func CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error) {
	// Calculate the expiration time
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

	// Create JWT claims
	claims := &models.JwtCustomClaims{
		Name: user.Name,
		ID:   user.ID.String(), // Convert uuid.UUID to string
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)), // New format for expiry
		},
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error) {
	// Calculate the expiration time and convert it to *jwt.NumericDate
	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry)))

	// Create JWT claims for refresh token
	claimsRefresh := &models.JwtCustomRefreshClaims{
		ID: user.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	// Sign the token with the secret
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return rt, nil
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims["id"].(string), nil
}

package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func genToken(userID string, role string, issuer string, secret string, expire int, tokenType string) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(time.Second * time.Duration(expire))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"type":    tokenType,
		"role":    role,
		"iss":     issuer,
		"iat":     now.Unix(),
		"nbf":     now.Unix(),
		"exp":     exp.Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, exp, err
}

func GenAccessToken(userID string, role string, issuer string, secret string, expire int) (string, time.Time, error) {
	return genToken(userID, role, issuer, secret, expire, "access")
}

func GenerateTokenByTime(userID string, iss string, secret string, d time.Duration) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(d)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"iss":     iss,
		"iat":     now.Unix(),
		"nbf":     now.Unix(),
		"exp":     exp.Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, exp, err
}

func ParseToken(tokenString string, secret string) (string, string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		//return "", jwt.ErrHashUnavailable
		return "", "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		var userID string
		var role string
		var tokenType string

		if v, ok := claims["user_id"]; ok {
			if s, ok := v.(string); ok {
				userID = s
			}
		}

		if v, ok := claims["role"]; ok {
			if s, ok := v.(string); ok {
				role = s
			}
		}

		if v, ok := claims["type"]; ok {
			if s, ok := v.(string); ok {
				tokenType = s

			}
		}

		return userID, role, tokenType, nil
	}

	return "", "", "", jwt.ErrTokenInvalidClaims
}

func GenRefreshToken(userID string, role string, issuer string, secret string, expire int) (string, time.Time, error) {
	return genToken(userID, role, issuer, secret, expire, "refresh")
}

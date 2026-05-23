package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userID string, role string, issuer string, secret string) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(24 * time.Hour * 365)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"iss":     issuer,
		"iat":     now.Unix(),
		"nbf":     now.Unix(),
		"exp":     exp.Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, exp, err
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

func ParseToken(tokenString string, secret string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		//return "", jwt.ErrHashUnavailable
		return "", "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		var userID string
		var role string

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

		return userID, role, nil
		//if userID, ok := claims["user_id"]; ok {
		//
		//	if i, ok := userID.(string); ok {
		//		return i, nil
		//	}
		//
		//}
	}

	return "", "", jwt.ErrTokenInvalidClaims
}

func RefreshToken(tokenString string, iss, secret string) (string, time.Time, error) {
	userID, role, err := ParseToken(tokenString, secret)
	if err != nil {
		return "", time.Now(), err
	}

	return GenerateToken(userID, role, iss, secret)
}

func GetUserIDByCtx(ctx context.Context) string {
	// JWT 中间件解析后，通常会将值以 json.Number 的类型存入 context
	jsonUid := ctx.Value("user_id")
	return jsonUid.(string)
}

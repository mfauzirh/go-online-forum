package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CreateToken handle JWT creation
//
// Parameters:
//   - id: User identifier owned by login user
//   - username: username owned by login user
//   - secretKey: key that will be used when signing JWT
//
// Returns:
//   - token (string): JWT token that created
//   - error: returns error if signing token failed
func CreateToken(id int64, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"exp":      time.Now().Add(10 * time.Minute).Unix(),
		},
	)

	key := []byte(secretKey)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// ValidateToken handle JWT validation
//
// Parameters:
//   - tokenStr: JWT in string format
//   - secretKey: secret key that has been used before when signing token in token creation
//
// Returns:
//   - user id of logged user
//   - username of logged user
//   - error if failed when parsing with claims or token is not valid
func ValidateToken(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	return int64(claims["id"].(float64)), claims["username"].(string), nil
}

func ValidateTokenWithoutExpiry(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	return int64(claims["id"].(float64)), claims["username"].(string), nil
}

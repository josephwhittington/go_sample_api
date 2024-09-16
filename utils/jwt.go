package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// probably should be stored in an environment variable
const secretKey = "TfQYcOZX0jOsfL6vJXW1yVZeB3B5UOLImzivC29OWLhx2Z14wrkTw4QbglcA20Jo"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// validate the signing method, this syntax of go's type checking system is called type assertion
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// email := claims["email"].(string)
	userId := claims["userId"].(float64) // this might be a no-go in production code, should look into this
	userIdInt64 := int64(userId)

	return userIdInt64, nil
}

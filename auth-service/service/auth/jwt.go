package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"auth-service/types"
	"auth-service/utils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"

type JWTClaims struct {
	UserID    string `json:"userID"`
	ExpiredAt int64  `json:"expiredAt"`
	jwt.RegisteredClaims
}

var secret = "0190b7df-1c44-7bbb-a76d-9bd11efe77ac"

func CreateJWT(userID string) (string, error) {
	expiration := time.Second * time.Duration(3000)

	claims := JWTClaims{
		UserID:    userID,
		ExpiredAt: time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetTokenFromRequest(r)

		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("falied to validate token: %v", err)
			unauthorized(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			unauthorized(w)
			return
		}

		claims := token.Claims.(*JWTClaims)
		userID := claims.UserID

		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			unauthorized(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(tokenString *jwt.Token) (interface{}, error) {
		if _, ok := tokenString.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tokenString.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func unauthorized(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
}

func GetUserIDFromContext(ctx context.Context) string {
	userID, ok := ctx.Value(UserKey).(string)
	if !ok {
		return ""
	}

	return userID
}

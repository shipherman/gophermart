package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/shipherman/gophermart/internal/db"
	mid "github.com/shipherman/gophermart/internal/models"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	User string
}

type Authenticator struct {
	Client db.DBClientInt
}

const tockenExpiration = time.Hour * 3
const sercretKey = "supersecretkey"

var ErrUserDoesNotExist = errors.New("no such user")

func NewAuthenticator(dbclient db.DBClientInt) Authenticator {
	return Authenticator{Client: dbclient}
}

func buildJWTString(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tockenExpiration)),
		},
		User: user,
	})
	tokenString, err := token.SignedString([]byte(sercretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getUser(tokenString string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(sercretKey), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return claims.User, fmt.Errorf("invalid token")
	}

	return claims.User, nil
}

func (a *Authenticator) Auth(u, p string) (jwt string, err error) {
	exist, err := a.Client.SelectUserExistence(u, p)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", ErrUserDoesNotExist
	}

	return buildJWTString(u)
}

func (a *Authenticator) CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if authenticated
		// Return 401 if not
		JWT := r.Header.Get("Authorization")
		if JWT == "" {
			http.Error(w, "AccessDenied", http.StatusUnauthorized)
			return
		}

		// Get user
		user, err := getUser(JWT)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Add user as context parameter
		r = r.WithContext(context.WithValue(r.Context(), mid.UserCtxKey{}, user))

		next.ServeHTTP(w, r)
	})

}

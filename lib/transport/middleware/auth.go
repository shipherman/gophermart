package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/shipherman/gophermart/lib/db"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	User string
}

const TOKEN_EXP = time.Hour * 3
const SECRET_KEY = "supersecretkey"

func buildJWTString(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
		},
		User: user,
	})
	tokenString, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getUser(tokenString string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return claims.User, fmt.Errorf("invalid token")
	}

	return claims.User, nil
}

func Auth(u, p string) (jwt string, err error) {
	exist, _ := db.SelectUserExistence(u, p)
	if !exist {
		return "", fmt.Errorf("no such user")
	}

	return buildJWTString(u)
}

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if authenticated
		// Return 401 if not
		JWT := r.Header.Get("Authorization")
		if JWT == "" {
			http.Error(w, "AccessDenied", http.StatusUnauthorized)
			return
		}

		JWTarr := strings.Split(JWT, " ")

		if JWTarr[0] != "Bearer" {
			http.Error(w, "Auth method shoud be Bearer", http.StatusUnauthorized)
			return
		}

		_, err := getUser(JWTarr[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

}

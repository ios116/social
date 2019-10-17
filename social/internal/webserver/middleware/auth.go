package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

// SessionContext data in session
type SessionContext struct {
	ID interface{}
}

// SessionContextKey session context key
type SessionContextKey string

// SessionKey key of session context
var SessionKey = SessionContextKey("session")

// SigningKey solt
var SigningKey = []byte("salt")

// Claims fo session
type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func sessionCheck(r *http.Request) (string, error) {
	// if no Auth cookie is set then return a 404 not found page
	cookie, err := r.Cookie("Auth")
	if err != nil {
		return "", err
	}

	// Return a token using the cookie
	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure token's signature wasn't changed
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return SigningKey, nil
	})

	if err != nil {
		return "", err
	}
	// Grab the tokens claims and pass it into the original request
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		//fmt.Println("claim", claims.Username)
		return claims.ID, nil
	}
	return "", err
}

// SessionMiddleware middleware for http request
func SessionMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := SessionContext{}
		if id, err := sessionCheck(r); err == nil {
			params.ID = id
		}
		ctx := context.WithValue(r.Context(), SessionKey, params)
		r = r.WithContext(ctx)
		inner.ServeHTTP(w, r)
	})
}

// SetToken set the token in the cookie if the password is correct
func SetToken(w http.ResponseWriter, id string, hours time.Duration) error {
	mySigningKey := SigningKey
	// expire the token and cookie
	expireToken := time.Now().Add(24 * time.Hour * hours).Unix()
	expireCookie := time.Now().Add(24 * time.Hour * hours)
	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return err
	}
	cookie := http.Cookie{Name: "Auth", Value: ss, Expires: expireCookie, HttpOnly: true, Path: "/"}
	http.SetCookie(w, &cookie)
	return nil
}

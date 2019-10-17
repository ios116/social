package auth

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)
// check auth
// set auth
// delete auth

type SessionProvider interface {
	GetSession(r *http.Request) (string, error)
	SetSession(w http.ResponseWriter, id string) error
	DeleteSession(w http.ResponseWriter) error
}

type Manager struct {
	SessionKey  string
	MaxLifetime time.Duration
}

// DeleteSession delete auth cookies
 func (s *Manager)  DeleteSession(w http.ResponseWriter) error {
	 deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now(), Path: "/"}
	 http.SetCookie(w, &deleteCookie)
 	return nil
 }

// GetSession returning a user id as the string in cookies or returning error
func (s *Manager)  GetSession(r *http.Request) (string, error) {
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

// SetSession set the token in the cookie if the password is correct
func (s *Manager) SetSession(w http.ResponseWriter, id string) error {
	mySigningKey := SigningKey
	// expire the token and cookie
	expireToken := time.Now().Add(24 * time.Hour * s.MaxLifetime).Unix()
	expireCookie := time.Now().Add(24 * time.Hour * s.MaxLifetime)
	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return err
	}
	cookie := http.Cookie{Name: s.SessionKey, Value: ss, Expires: expireCookie, HttpOnly: true, Path: "/"}
	http.SetCookie(w, &cookie)
	return nil
}

type Context struct {
	ID int64
	Login string
	Email string
}

// SigningKey salt
var SigningKey = []byte("salt")

// Claims fo auth
type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}





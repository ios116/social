package webserver

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type SessionProvider interface {
	GetSession(r *http.Request) (SessionContext, error)
	SetSession(w http.ResponseWriter, user SessionContext) error
	DeleteSession(w http.ResponseWriter)
}

type SessionContext struct {
	ID    int64
	Login string
}

type SessionManager struct {
	SessionKey  string
	MaxLifetime time.Duration
}

func NewSessionManager(sessionKey string, maxLifetime time.Duration) *SessionManager {
	return &SessionManager{SessionKey: sessionKey, MaxLifetime: maxLifetime}
}

// DeleteSession delete auth cookies
func (s *SessionManager) DeleteSession(w http.ResponseWriter) {
	deleteCookie := http.Cookie{Name: s.SessionKey, Value: "none", Expires: time.Now(), Path: "/"}
	http.SetCookie(w, &deleteCookie)
}

// GetSession returning a user id as the string in cookies or returning error
func (s *SessionManager) GetSession(r *http.Request) (SessionContext, error) {
	// if no Auth cookie is set then return a 404 not found page
	sess := SessionContext{}
	cookie, err := r.Cookie(s.SessionKey)
	if err != nil {
		return sess, err
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
		return sess, err
	}
	// Grab the tokens claims and pass it into the original request
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		//fmt.Println("claim", claims.Username)
		return SessionContext{
			ID:    claims.ID,
			Login: claims.Login,
		}, nil
	}
	return sess, err
}

// SetSession set the token in the cookie if the password is correct
func (s *SessionManager) SetSession(w http.ResponseWriter, user SessionContext) error {
	mySigningKey := SigningKey
	// expire the token and cookie
	expireToken := time.Now().Add(24 * time.Hour * s.MaxLifetime).Unix()
	expireCookie := time.Now().Add(24 * time.Hour * s.MaxLifetime)
	claims := Claims{
		user.ID,
		user.Login,
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

// SigningKey salt
var SigningKey = []byte("salt")

// Claims fo auth
type Claims struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	jwt.StandardClaims
}

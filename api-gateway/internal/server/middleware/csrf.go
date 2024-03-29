package middleware


import (
"crypto/sha256"
"encoding/base64"
log "github.com/sirupsen/logrus"
"io"
)

const (
	CSRFHeader = "X-CSRF-Token"
	csrfSalt = "KbWaoi5xtDC3GEfBa9ovQdzOzXsuVU9I"
)

// Create CSRF token

func MakeToken(sid string) string {
	hash := sha256.New()
	_, err := io.WriteString(hash, csrfSalt+sid)
	if err != nil {
		log.Errorf("Make CSRF Token: %v", err)
	}
	token := base64.RawStdEncoding.EncodeToString(hash.Sum(nil))
	return token
}

// Validate CSRF token

func ValidateToken(token string, sid string) bool {
	trueToken := MakeToken(sid)
	return token == trueToken
}
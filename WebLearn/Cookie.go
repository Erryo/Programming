package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
)

var secretKey []byte

func WriteCookie(w http.ResponseWriter, cookie http.Cookie) {
	cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))

	// Check the total length of the cookie contents.
	// error if it's more than 4096 bytes.
	if len(cookie.String()) < 4096 {
		// Write the cookie as normal.
		http.SetCookie(w, &cookie)
	}
}

func Read(r *http.Request, name string) string {
	// Read the cookie as normal.
	cookie, err := r.Cookie(name)
	if err != nil {
		return ""
	}

	// Decode the base64-encoded cookie value. If the cookie didn't contain a
	// valid base64-encoded value, this operation will fail and we return an
	// ErrInvalidValue error.
	value, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return ""
	}

	// Return the decoded cookie value.
	return string(value)
}

func SetLoggedCookie(w http.ResponseWriter, r *http.Request, Value string) {
	cookie := http.Cookie{
		Name:     "UserStatus",
		Value:    Value,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}
	fmt.Println("SetLoggedCookie:secretKey: ", secretKey)

	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(cookie.Name))
	mac.Write([]byte(cookie.Value))
	signature := mac.Sum(nil)
	// Prepend the cookie value with the HMAC signature.
	cookie.Value = string(signature) + cookie.Value
	WriteCookie(w, cookie)
}

func GetLoggedCookie(w http.ResponseWriter, r *http.Request) bool {
	cookieValue := Read(r, "UserStatus")

	fmt.Println("GetLoggedCookie:cookieValue: ", cookieValue)
	if cookieValue == "" {
		return false
	}
	if len(cookieValue) < sha256.Size {
		return false
	}
	signature := cookieValue[:sha256.Size]
	value := cookieValue[sha256.Size:]

	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte("UserStatus"))
	mac.Write([]byte(value))
	expectedSignature := mac.Sum(nil)

	// Check that the recalculated signature matches the signature we received
	// in the cookie. If they match, we can be confident that the cookie name
	// and value haven't been edited by the client.
	if !hmac.Equal([]byte(signature), expectedSignature) {
		return false
	}
	if LogInData.Username == "" {
		return false
	}
	fmt.Println("GetLoggedCookie:designed value", value)
	return value == LogInData.Username
}

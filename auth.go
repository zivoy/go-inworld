package goinworld

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

func GetAuthorization(host string, apiKey ApiKey) (string, error) {
	host = strings.Replace(host, ":443", "", 1)                            // should be -1 since there is a chance for it different from the js version
	const method = "ai.inworld.studio.v1alpha.Tokens/GenerateSessionToken" //cant get it from the proto file like in the js version, so i am also not bothering slicing it

	datetime := getTimeFormat(time.Now())
	nonceBytes := make([]byte, 8)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		return "", err
	}
	nonce := fmt.Sprintf("%x", nonceBytes)[1:]

	signature := signSignature(apiKey.Secret,
		datetime,
		host,
		method,
		nonce,
	)

	return fmt.Sprintf("IW1-HMAC-SHA256 ApiKey=%s,DateTime=%s,Nonce=%s,Signature=%s", apiKey.Key, datetime, nonce, signature), nil
}

func signSignature(key string, params ...string) string {
	s := []byte("IW1" + key)
	for _, param := range params {
		h := hmac.New(sha256.New, s)
		h.Write([]byte(param))
		s = h.Sum(nil)
	}

	h := hmac.New(sha256.New, s)
	h.Write([]byte("iw1_request"))
	return hex.EncodeToString(h.Sum(nil))
}

func getTimeFormat(t time.Time) string {
	t = t.UTC()
	return fmt.Sprintf("%.4d%.2d%.2d%.2d%.2d%.2d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

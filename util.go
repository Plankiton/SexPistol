package api

import (
    "golang.org/x/crypto/bcrypt"
    "crypto/sha1"
    "fmt"
    "io"
)

type Response struct {
    Message   string       `json:"message,omitempty"`
    Type      string       `json:"type,omitempty"`
    Data      interface{}  `json:"data,omitempty"`
}

type Request struct {
    Token   string             `json:"auth,omitempty"`
    Data    map[string]string  `json:"data,omitempty"`
}

func ToHash(s string) string {
    h := sha1.New()
    io.WriteString(h, s)
    return fmt.Sprintf("%x", h.Sum(nil))
}
func ToPassHash(s string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
    return string(hash), err
}

func CheckPass(p []byte, s string) (error) {
    err := bcrypt.CompareHashAndPassword(p, []byte(s))
    return err
}

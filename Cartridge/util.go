package SexDB
import (
    "crypto/sha1"
    "io"

    "fmt"
    "golang.org/x/crypto/bcrypt"
)

// Utility Database Function to create sha1 hash
func ToHash(s string) string {
    h := sha1.New()
    io.WriteString(h, s)
    return fmt.Sprintf("%x", h.Sum(nil))
}

// Utility Database Function to create password bcrypt hash with sault
func ToPassHash(s string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
    return string(hash), err
}

// Utility Database Function to validate password bcrypt hash
func CheckPass(p []byte, s string) (error) {
    err := bcrypt.CompareHashAndPassword(p, []byte(s))
    return err
}


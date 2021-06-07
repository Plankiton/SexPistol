package SexDB
import (
    "crypto/sha1"
    "io"

    "fmt"
    "golang.org/x/crypto/bcrypt"
)

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


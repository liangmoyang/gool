package str

import (
	"crypto/sha1"
	"encoding/hex"
)

// ToSha1  sha1 encryption
func ToSha1(s string) string {

	hash := sha1.New()

	hash.Write([]byte(s))

	bs := hash.Sum(nil)

	sign := hex.EncodeToString(bs)

	return sign
}

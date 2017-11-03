// sha1.go
package small

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(key string) string {
	h := sha1.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}

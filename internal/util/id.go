package util

import (
	"crypto/md5"
	"encoding/hex"
)

// GetId returns the md5 hash of the source string, which is used as the id of the error
func GetId(source string) string {
	hasher := md5.New()
	hasher.Write([]byte(source))
	return hex.EncodeToString(hasher.Sum(nil))
}

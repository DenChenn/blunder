package util

import (
	"crypto/md5"
	"encoding/hex"
)

func GetId(source string) string {
	hasher := md5.New()
	hasher.Write([]byte(source))
	return hex.EncodeToString(hasher.Sum(nil))
}

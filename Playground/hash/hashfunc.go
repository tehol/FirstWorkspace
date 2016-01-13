package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func HashString(s string) (hashString string) {

	reader := strings.NewReader(s)
	size := reader.Size()
	data := make([]byte, size)
	reader.Read(data)
	hash := sha256.Sum256(data)
	fmt.Println(hash)
	hashString = EncodeBase64(hash[:])
	return
}

func EncodeBase64(data []byte) string {
	hashString := base64.StdEncoding.EncodeToString(data)
	return hashString
}

package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
)

//签名
func Sign(params map[string]string, appKey string )  string {
	keys := []string{}
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	strSign := ""
	for _, key := range keys {
		if key == "sign" || key == "s" {
			continue
		}
		v := params[key]
		if v == "" {
			continue
		}
		strSign += fmt.Sprintf("%v#", v)
	}
	strSign += appKey
	return hashSha1(strSign)
}

func hashSha1(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum([]byte("")))
}
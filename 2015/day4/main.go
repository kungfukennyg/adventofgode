package day4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

func md5WithPrefix(secretKey, prefix string) int {
	for i := range math.MaxInt64 {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", secretKey, i)))
		hex := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hex, prefix) {
			return i
		}
	}
	panic("Hash not found")
}

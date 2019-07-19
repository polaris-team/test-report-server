package random

import (
	"math/rand"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func Token() string {
	u1 := uuid.NewV4()
	return strings.ReplaceAll(u1.String(), "-", "")
}

func RandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

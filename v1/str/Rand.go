package str

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// RandString generate random string
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())

	rs := make([]string, length)

	for i := 0; i < length; i++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}

	return strings.Join(rs, "")
}

// RandStringUpper generate random upper string
func RandStringUpper(length int) string {

	bytes := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < length; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}

	return string(bytes)
}

package cryptox

import (
	"crypto/md5"
	"fmt"
)

func GetMd5(strings ...string) string {
	h := md5.New()
	for _, s := range strings {
		_, _ = h.Write([]byte(s))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

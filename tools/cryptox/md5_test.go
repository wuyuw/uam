package cryptox

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	res := GetMd5("hello", "go")
	fmt.Println(res)
}

package rsa

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var publicPem = `
-----BEGIN RSA Public Key-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3wXT60A2mIZ0DUFmjT2N
dsw50FokUazD2QMWi13tUr0TiOXB4P5T2JRDbVeNmufIrZLBz56K6d0lyDB73fcO
+KJ411eNSGZOupr4b+tXpJ7BsZ79TrhugyYMuw4xggCoycfZSR9SJ+Aso9EO/sOg
qalNlz+uJZ8hq7ysB2aqPpj8CUwGIuHDAb0rkLXWYN1exjPRmC00SU6M171JdU89
qESIFW7Ei9TRBfGP6rd1ZH2bkA6npbPVWK3kjoJlDbi+m9u6T1EcjBq3JIJx1b4P
eIM2YYDdR1dv5NRrmTmsFsxB7hOJztY9RG9g1sAx7m+90rlv/XvOmHoj1conLGUv
zQIDAQAB
-----END RSA Public Key-----
`

var privatePem = `
-----BEGIN RSA Private Key-----
MIIEogIBAAKCAQEA3wXT60A2mIZ0DUFmjT2Ndsw50FokUazD2QMWi13tUr0TiOXB
4P5T2JRDbVeNmufIrZLBz56K6d0lyDB73fcO+KJ411eNSGZOupr4b+tXpJ7BsZ79
TrhugyYMuw4xggCoycfZSR9SJ+Aso9EO/sOgqalNlz+uJZ8hq7ysB2aqPpj8CUwG
IuHDAb0rkLXWYN1exjPRmC00SU6M171JdU89qESIFW7Ei9TRBfGP6rd1ZH2bkA6n
pbPVWK3kjoJlDbi+m9u6T1EcjBq3JIJx1b4PeIM2YYDdR1dv5NRrmTmsFsxB7hOJ
ztY9RG9g1sAx7m+90rlv/XvOmHoj1conLGUvzQIDAQABAoIBACrQLpd5s0FihkLJ
LEuu5kpI+ExEEbbQKKSvUBOfC2EXxPlByg9MI4JvK+aAqUF0f3S6uJQHxnkQqCEf
FZhNxkT6w6HrP8cHRNPTzh+GGUQT6fEUKWKES0rH8ieymNRxFfXudIryBU58XXVx
O6Syn1QSmT+QzPiR7N/QD2I4VjQ/xFlF236lhIJ0rGqs9qxRT7c7+nB9svtOHS5k
k3rWS3FNz2qKV2l7DwbwwinY976vbLbxLbP6JugrTYpE5d2OMa2JazTeKtlmuhcW
jpkMypnA1ohG6SdxgJf9S+rIHFz1Iiu2dj6SuDADgyiOg585kuG2Xa0iHeqWjYTE
WQ6SDqECgYEA8zRBOxzb5r3E3oPQM6M1USnXbqj058pzec9QNslTonrOIx2PJQvs
zmjM2yoZBBbcRTZuhX+yyiGvZOnsREpXKRKbz+sSwxg47LDNwIZhEuRIuVazVAlX
c3wAmrQOr/f52jD5uvyOg/zYNkM/zig6yWKKCbMSq3qTJt9OTfCj3MkCgYEA6sG/
eA3LrpBqB4N/pMpQNaUTfFb455v+CvwucaYGgc/GRD3GOsJ0mwV+Nya+W1WVONZ7
4VYhYR4/HCD8b4rAR5Z+jWpCTX8p2py26r2PlEigTgPqizVwu/otKqY1zp7c11nz
p+91y/wqgPPggr/asDR/vxmj8TqqPytPbqGTMOUCgYApilFiAWnmHZ/Uyfrz9vqS
ZG0xr5Y1STU0Jx7yXKz2Ybd39AKRN1o5X1kuTiB7vFPfVo7GKqulLt/AgtwiRfhh
QZZvix1nSWnfs8tRCSLnkSqCzbZPslDHnvSTeBHSKK76f8cIEz9ceAGOMypg0ipI
X5ZoVbfopkUgLKA5W9MBUQKBgFpyX3S/y/PrzA4tCebR0+l3OnSzhZ6mqVBOLQ64
atVk2fy82D0XYpm/mgthsAG8jYuih4QgDSg/4QzTYK8RBFgQkZ2mjPkSv2ts6cSz
WDhHawvj0l/kLRUfpHtEIoMDDg8ipw/S4M3A0Bdy3tNBW957u6RDGrj8Y0+HPklf
kzhJAoGAWme8FgUbQTIffjZ8I+1uTMnWsC/YMF0fwckaHw5fCZaPBYptHEwlkQxb
qRrsXaOVN+xYauCgvTY2jpanfSSEQrNG/AxhhqZ5SuE8lvd5DK0tUQbGqFV8vbZo
61NcE8bV1XNVY+PUuJWBkB7D9GI3Y3VO+6CfCJlap/JA8WpBJow=
-----END RSA Private Key-----
`

func TestRSA(t *testing.T) {

	Convey("测试RSA加解密", t, func() {
		caseStr := "hello world"
		message := []byte(caseStr)
		//加密
		cipherText := RSAEncrypt(message, publicPem)
		// fmt.Println("加密后为：", string(cipherText))
		//解密
		plainText := RSADecrypt(cipherText, privatePem)
		// fmt.Println("解密后为：", string(plainText))
		So(string(plainText), ShouldEqual, caseStr)
	})
}

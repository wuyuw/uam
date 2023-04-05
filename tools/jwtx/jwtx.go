package jwtx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

const JWTRedisPrefix = "uamadmin:jwt:"

// GetJwtToken 生成JWT Token
func GetJwtToken(secretKey string, iat, seconds int64, custom map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range custom {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// GetCtxValueInt64 从ctx中获取int64类型的值
func GetCtxValueInt64(ctx context.Context, key string) (int64, error) {
	valIface := ctx.Value(CtxKey(key))
	if valIface == nil {
		return 0, fmt.Errorf("key不存在: %s", key)
	}
	jsonNum, ok := valIface.(json.Number)
	if !ok {
		return 0, fmt.Errorf("类型转换错误, key: %s value: %v", key, valIface)
	}
	val, err := jsonNum.Int64()
	if err != nil {
		return 0, fmt.Errorf("UID类型转换错误, key: %s value: %v err: %s", key, jsonNum, err)
	}
	return val, nil
}

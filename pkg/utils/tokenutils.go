/*
Package utils comment
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
SPDX-License-Identifier: Apache-2.0
*/
package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secret = "nacos-secret"

func GetSecret() string {
	return secret
}

// NewSignedToken new signed token
func NewSignedToken(userId int64, username string, userType string, uuid string, expireHour int) (string, error) {
	claims := &JwtClaims{
		UserId:   userId,
		UserName: username,
		UserType: userType,
		Uuid:     uuid,
	}
	claims.IssuedAt = time.Now().Unix()
	// 不会过期
	if expireHour == 0 {
		claims.ExpiresAt = 0
	} else {
		claims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(expireHour)).Unix()
	}
	return toSignedToken(claims)
}

// toSignedToken to signed token
func toSignedToken(claims *JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// LoadJwtClaims load jwt claims
func LoadJwtClaims(tokenText string) (JwtMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenText, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	// 强制类型转换，类似于Java中的instance of
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, fmt.Errorf("can not load token")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}
	return JwtClaimsToJwtMapClaims(claims), nil
}

// JwtClaims jwt claims
type JwtClaims struct {
	jwt.StandardClaims
	UserId   int64  // 用户id
	UserName string // 用户名
	UserType string // 用户类型
	Uuid     string // 用户uuid
}

type JwtMapClaims map[string]interface{}

func (j JwtMapClaims) GetUserName() string {
	return j["user_name"].(string)
}

func (j JwtMapClaims) GetUserId() int64 {
	return j["user_id"].(int64)
}

func (j JwtMapClaims) GetUserType() string {
	return j["user_type"].(string)
}

func (j JwtMapClaims) GetUuid() string {
	return j["uuid"].(string)
}

func JwtClaimsToJwtMapClaims(claims *JwtClaims) JwtMapClaims {
	var jwtMapClaims = JwtMapClaims{}
	jwtMapClaims["user_id"] = claims.UserId
	jwtMapClaims["user_name"] = claims.UserName
	jwtMapClaims["user_type"] = claims.UserType
	jwtMapClaims["uuid"] = claims.Uuid
	jwtMapClaims["exp"] = claims.ExpiresAt
	jwtMapClaims["iat"] = claims.IssuedAt
	return jwtMapClaims
}

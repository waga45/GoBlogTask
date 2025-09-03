package component

import (
	"GoBlogServer/utils"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"sync"
	"time"
)

// payload
type Claims struct {
	UserId   int    `json:"id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// 管理器
type JwtManager struct {
	secretKey   string
	expire      time.Duration
	aesInstance *utils.AESEncryptor
}

// 单例
var (
	onceJwt     sync.Once
	instanceJwt *JwtManager
)

func GetJwtInstance() *JwtManager {
	if instanceJwt == nil {
		panic("JwtManager未完成初始化")
	}
	return instanceJwt
}

// 初始化
func InitJwtManager(secretKey string, expire int) *JwtManager {
	onceJwt.Do(func() {
		instanceJwt = &JwtManager{
			secretKey: secretKey,
			expire:    time.Duration(expire) * time.Hour,
		}
		b, _ := utils.NewAESEncryptor("*2025BotASK_BLOG")
		instanceJwt.aesInstance = b
	})
	return instanceJwt
}

func (j *JwtManager) GenToken(id int, username string) (string, error) {
	expirationTime := time.Now().Add(j.expire * time.Hour)
	claims := &Claims{
		UserId:   id,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "GoBlog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(j.secretKey))
	bs, _ := j.aesInstance.EncryptCBC([]byte(tokenString))
	return string(bs), nil
}

func (j *JwtManager) ParserToken(tokenString string) (*Claims, error) {
	if len(tokenString) == 0 {
		return nil, errors.New("请输入Token")
	}
	tokenBs, _ := j.aesInstance.DecryptCBC(tokenString)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(string(tokenBs), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const SECRET = "taoshihan"

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserId   uint   `json:"userId"`
	RoleId   uint   `json:"roleId"`
	jwt.StandardClaims
}

// GenerateToken 派发token
func GenerateToken(username, password string, roleId uint, userId uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	// 待加密的结构体
	claims := Claims{
		username,
		password,
		userId,
		roleId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 生效时间
			Issuer:    "gaoding-sheji",   // 签发人
		},
	}
	// 使用HS256进行对结构体加密
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 转换tokenstring
	token, err := tokenClaims.SignedString([]byte(SECRET))

	return token, err
}

// ParseToken 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// ParseTokenRequest 解析token参数获取解密参数
func ParseTokenRequest(c *gin.Context) (*Claims, error) {
	var err error
	// 获取存在的token
	token := c.GetHeader("Authorization")
	authInfo, err := ParseToken(token)
	if err != nil {
		return nil, err
	}
	return authInfo, nil
}

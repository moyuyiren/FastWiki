package Token

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// token结构体信息
type MyClaims struct {
	Userid          int64 `json:"userid"`
	Userpermissions int8  `json:"userpermissions"`
	jwt.RegisteredClaims
}

// 秘钥
var mySercet = []byte("FastWiki")

// Token存在时间
const TokenExpireDuration = time.Hour * 168 * time.Duration(1)

// genToken生成token
func GenToken(Userid int64, userpermissions int8) (string, error) {
	c := MyClaims{
		Userid:          Userid,
		Userpermissions: userpermissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                          // 生效时间
			Issuer:    "my-project",                                            //签发人
		},
	}
	//使用签名方法创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用秘钥加密token
	atoken, err := token.SignedString(mySercet)
	return atoken, err
}

// ParseToken解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySercet, nil
	})

	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")

}

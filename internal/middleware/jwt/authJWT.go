package jwt

import (
	"SummerProject/common"
	"SummerProject/internal/middleware"
	"errors"
	gojwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"time"
)

var jwtKey []byte

func init() {
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

type Claims struct {
	Uid int
	gojwt.StandardClaims
}

// Award 生成Token
func Award(uid *int) (string, error) {
	// 过期时间
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Uid: *uid, //MUser.ID由uint转换为int
		StandardClaims: gojwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	// 生成token
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*gojwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := gojwt.ParseWithClaims(tokenStr, claims, func(t *gojwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, claims, err
}
func AuthJWT() middleware.Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("Authorization")
			_, claim, err := ParseToken(tokenStr)
			if err != nil {
				common.Error(w, errors.New("请先登录！"))
				return
			}
			log.Println("AuthJWT: 执行操作的用户Uid:", claim.Uid)
			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

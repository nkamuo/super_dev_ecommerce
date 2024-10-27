package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/superdev/ecommerce/gateway/config"
)

type JWTMiddleware struct {
	JWTSecret string
	conf      *config.Config
}

func NewJWTMiddleware(conf *config.Config) *JWTMiddleware {
	return &JWTMiddleware{JWTSecret: conf.JWTSecret}
}

func (m *JWTMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "missing or invalid token",
			})
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.JWTSecret), nil
		})

		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "invalid token",
			})
			return
		}
		c.Next()
	}
}

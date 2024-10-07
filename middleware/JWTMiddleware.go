package middleware

import (
	"fmt"
	"os"
	"pos-webapp/helpers"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func checkJWT(middlewareAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Token string di dapat dari header request Authorization
		authHeader := c.Request.Header.Get("Authorization")
		// Mengambil token dari "Bearer <token>"
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			// Parse takes the token string and a function for looking up the key. The latter is especially
			// Useful if you use multiple keys for your application. The standard is to use 'kid' in the
			// Head of the token to identify which key to use, but the parsed token (head and claims) is provided
			// to the callback, providing flexibility.
			token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signin method: %v", token.Header["alg"])
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				RoleId := claims["role_id"]
				c.Set("jwt_user_id", claims["user_id"])
				c.Set("jwt_role_id", claims["role_id"])
				c.Set("jwt_business_id", claims["business_id"])
				c.Set("jwt_warehouse_id", claims["warehouse_id"])

				if middlewareAdmin && RoleId == 0 {
					helpers.ElorResponse(c, "Akses dibatasi, anda tidak memiliki akses ini!")
					c.Abort()
					return
				}
			} else {
				helpers.ElorWithData(c, "Token tidak valid!", err)
				c.Abort()
				return
			}
		} else {
			helpers.ElorResponse(c, "Autorisasi diperlukan untuk mengakses endpoint ini!")
			c.Abort()
			return
		}
	}
}

func IsAuth() gin.HandlerFunc {
	return checkJWT(true)
}

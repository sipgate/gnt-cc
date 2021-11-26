package auth

import (
	"gnt-cc/config"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const identityKey = "id"

type Credentials struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GetMiddleware() (ginJWTMiddleware *jwt.GinJWTMiddleware) {
	ginJWTMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "gnt-cc",
		Key:         []byte(config.Get().JwtSigningKey),
		Timeout:     config.Get().JwtExpire,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Username: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Credentials
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if validateUser(userID, password) {
				return &User{
					Username: userID,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// for now we just validate if the data is valid User struct
			if _, ok := data.(*User); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		SendCookie:     true,
		CookieHTTPOnly: true,
		SecureCookie:   true,
		CookieName:     "jwt",
		TokenLookup:    "cookie:jwt",
		CookieSameSite: http.SameSiteLaxMode,

		TimeFunc: time.Now,
	})

	if err != nil {
		log.Errorf("Error initializing JWT middleware: %s", err)
	}

	return
}

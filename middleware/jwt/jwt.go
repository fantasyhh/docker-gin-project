package jwt

import (
	"github.com/fantasyhh/drizzle/models"
	"github.com/fantasyhh/drizzle/pkg/e"
	"log"
	"net/http"
	"time"

	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "identity"

// User demo
type User struct {
	UserName string
}

// AuthMiddleware 为自定义的jwt认证中间件，认证过程走这里
func AuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "baird jwt",
		Key:         []byte("gaga"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		// 用户名密码认证，返回一个data给PayloadFunc
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Auth
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			isAuth, err := models.CheckAuth(username, password)
			if err != nil {
				return "", jwt.ErrFailedAuthentication
			}
			if isAuth {
				return &User{
					UserName: username,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// 上面的认证成功后拿到认证后的数据嵌入claims中，随后库内部生成token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		// 生成token成功后，返程登录成功信息
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, gin.H{
				"msg":    "用户认证成功！",
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		// 请求api时，取出token中的数据然后拿到用户信息
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		// 查看取出的用户信息是否属于被授权了这个api的使用
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		//在登录授权用户时发生任何错误，或者在请求中没有传递令牌或无效令牌时，将发生以下情况
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"msg":    e.ErrAuthFail,
				"detail": message,
			})
		},
		// 注销成功
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "注销成功",
				"data": nil,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}

//var AuthM  = AuthMiddleware()

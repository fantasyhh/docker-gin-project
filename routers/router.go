package routers

import (
	"github.com/fantasyhh/drizzle/middleware/jwt"
	"github.com/fantasyhh/drizzle/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
    // swagger doc
	_ "github.com/fantasyhh/drizzle/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/fantasyhh/drizzle/routers/api"
	v1 "github.com/fantasyhh/drizzle/routers/api/v1"
)

// InitRouter initialize all router
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	// token auth
	authMiddleware := jwt.AuthMiddleware()
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/logout", authMiddleware.LogoutHandler)
	r.GET("/refresh_token", authMiddleware.RefreshHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", api.PING)
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(authMiddleware.MiddlewareFunc())
	{
		apiv1.GET("/products", v1.ProductList)
		apiv1.GET("/products/:pk", v1.ProductRetrieve)
		apiv1.POST("/products", v1.ProductCreate)
		apiv1.PATCH("/products/:pk", v1.ProductUpdate)
		apiv1.DELETE("/products/:pk", v1.ProductDestroy)

		apiv1.GET("/productlines", v1.ProductLineList)
		apiv1.GET("/productlines/:pk", v1.ProductLineRetrieve)
		apiv1.POST("/productlines", v1.ProductLineCreate)
		apiv1.PATCH("/productlines/:pk", v1.ProductLineUpdate)
		apiv1.DELETE("/productlines/:pk", v1.ProductLineDestroy)
		apiv1.GET("/productlines/:pk/products", v1.ProductInLineList)

	}

	return r
}

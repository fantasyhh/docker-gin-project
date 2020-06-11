package main

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"

	"github.com/fantasyhh/drizzle/models"
	"github.com/fantasyhh/drizzle/pkg/setting"
	"github.com/fantasyhh/drizzle/routers"
)

func init() {
	setting.Setup()
	models.Setup()

}

func TestAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Add /example route via handler function to the gin instancego
	handler := routers.InitRouter()
	// Create httpexpect instance
	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		//Printers: []httpexpect.Printer{
		//	httpexpect.NewCurlPrinter(t),
		//	httpexpect.NewDebugPrinter(t, false),
		//},
	})

	// *************test ping*************
	e.GET("/ping").
		Expect().
		Status(http.StatusOK).JSON().Object().ValueEqual("message", "pong")

	// *************test admin login*************
	e.POST("/login").WithJSON(map[string]string{"username": "errorUser", "password": "errorUser"}).
		Expect().
		Status(http.StatusUnauthorized)
	r := e.POST("/login").WithJSON(map[string]string{"username": "admin", "password": "admin"}).
		Expect().
		Status(http.StatusOK).JSON().Object()
	token := r.Value("token").String().Raw()
	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	// *************test productLine*************
	auth.GET("/api/v1/productlines").
		Expect().
		Status(http.StatusOK)

	linePostData := models.ProductLine{ProductLine: "Planes", TextDescription: "this is a test line"}
	auth.POST("/api/v1/productlines").WithJSON(map[string]string{"bad": "json"}).
		Expect().
		Status(http.StatusBadRequest) //无法绑定json
	auth.POST("/api/v1/productlines").WithJSON(linePostData).
		Expect().
		Status(http.StatusBadRequest) //已经存在相同主键无法创建或未知错误
	linePostData.ProductLine = "testLine"
	auth.POST("/api/v1/productlines").WithJSON(linePostData).
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Equal(linePostData)

	auth.GET("/api/v1/productlines/gagaga").
		Expect().
		Status(http.StatusNotFound) //不存在的主键或未知错误,下同
	auth.GET("/api/v1/productlines/TestLine").
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Equal(linePostData)

	linePostData.TextDescription = "this is a  updated test line"
	auth.PATCH("/api/v1/productlines/gagaga").WithJSON(map[string]string{"textDescription": linePostData.TextDescription}).
		Expect().
		Status(http.StatusNotFound)
	auth.PATCH("/api/v1/productlines/TestLine").WithJSON(map[string]string{"textDescription": linePostData.TextDescription}).
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Equal(linePostData)

	auth.DELETE("/api/v1/productlines/gagaga").
		Expect().
		Status(http.StatusNotFound)
	auth.DELETE("/api/v1/productlines/TestLine").
		Expect().
		Status(http.StatusNoContent)

	// *************test Product*************
	auth.GET("/api/v1/products").
		Expect().
		Status(http.StatusOK)

	productPostData := models.Product{
		ProductCode:        "S10_1678",
		ProductName:        "testPP",
		ProductLine:        "Motorcycles",
		ProductScale:       "1:10",
		ProductVendor:      "fantasyhh",
		ProductDescription: "this is a test product",
		QuantityInStock:    8888,
		BuyPrice:           66.66,
		MSRP:               33.3,
	}

	auth.POST("/api/v1/products").WithJSON(map[string]string{"bad": "json"}).
		Expect().
		Status(http.StatusBadRequest) //无法绑定json
	auth.POST("/api/v1/products").WithJSON(productPostData).
		Expect().
		Status(http.StatusBadRequest) //已经存在相同主键无法创建或未知错误
	productPostData.ProductCode = "Test_007"
	auth.POST("/api/v1/products").WithJSON(productPostData).
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Equal(productPostData)

	auth.GET("/api/v1/products/gagaga").
		Expect().
		Status(http.StatusNotFound) //不存在的主键或未知错误,下同
	auth.GET("/api/v1/products/Test_007").
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Equal(productPostData)

	productPostData.ProductDescription = "this is a  updated test product"
	auth.PATCH("/api/v1/products/gagaga").WithJSON(map[string]string{"ProductDescription": productPostData.ProductDescription}).
		Expect().
		Status(http.StatusNotFound)
	auth.PATCH("/api/v1/products/Test_007").WithJSON(map[string]string{"ProductDescription": productPostData.ProductDescription}).
		Expect().
		Status(http.StatusOK).JSON().Object().Value("data").Equal(productPostData)

	// test product in line
	auth.GET("/api/v1/productlines/Motorcycles/products").
		Expect().
		Status(http.StatusOK)

	auth.DELETE("/api/v1/products/gagaga").
		Expect().
		Status(http.StatusNotFound)
	auth.DELETE("/api/v1/products/Test_007").
		Expect().
		Status(http.StatusNoContent)


	// *************test logout*************
	auth.GET("/logout").
		Expect().
		Status(http.StatusOK)

	// *************test guest login*************

	r = e.POST("/login").WithJSON(map[string]string{"username": "guest", "password": "guest"}).
		Expect().
		Status(http.StatusOK).JSON().Object()
	token = r.Value("token").String().Raw()
	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})

	// *************test guest no permissions  for api *************
	auth.GET("/api/v1/productlines").
		Expect().
		Status(http.StatusForbidden)
	// end
}

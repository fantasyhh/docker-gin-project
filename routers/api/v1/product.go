package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fantasyhh/drizzle/models"
	"github.com/fantasyhh/drizzle/pkg/app"
	"github.com/fantasyhh/drizzle/pkg/e"
	"github.com/fantasyhh/drizzle/pkg/setting"
	"github.com/fantasyhh/drizzle/pkg/util"
)

// ProductList godoc
// @Summary List products
// @Description get all products
// @Tags Product
// @Accept json
// @Produce json
// @Param page query integer false "page num"
// @Success 200 {object} app.JSONResult
// @Failure 400 {object} app.FailJSONResult
// @Router /api/v1/products  [get]
func ProductList(c *gin.Context) {
	appG := app.Gin{C: c}
	//maps := make(map[string]interface{})

	products, err := models.ProductList(util.GetPage(c), setting.AppSetting.PageSize)
	if err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrListProductFail, err)
		return
	}

	appG.Response(http.StatusOK, products)

}

// ProductRetrieve godoc
// @Summary Retrieve a product
// @Description Retrieve a product by pk
// @Tags Product
// @Accept json
// @Produce json
// @Param pk path string true "pk is ProductCode"
// @Success 200 {object} app.JSONResult
// @Failure 400 {object} app.FailJSONResult
// @Failure 404 {object} app.FailJSONResult
// @Router /api/v1/products/{pk}  [get]
func ProductRetrieve(c *gin.Context) {
	// pk == primary key
	pk := c.Param("pk")
	appG := app.Gin{C: c}

	if isExist := models.ExistProduct(pk); isExist == false {
		existErr := errors.New("不存在的主键或未知错误")
		appG.FailResponse(http.StatusNotFound, e.ErrRetrieveProductFail, existErr)
		return
	}

	product, err := models.ProductRetrieve(pk)
	if err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrRetrieveProductFail, err)
		return
	}

	appG.Response(http.StatusOK, product)
}

// ProductCreate godoc
// @Summary Create a product
// @Description Create one product by json product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body models.Product  true "create product"
// @Success 200 {object} app.JSONResult
// @Failure 400 {object} app.FailJSONResult
// @Router /api/v1/products  [post]
func ProductCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	var jsonData models.Product

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrCreateProductFail, err)
		return
	}

	if isExist := models.ExistProduct(jsonData.ProductCode); isExist == true {
		existErr := errors.New("已经存在相同主键无法创建或未知错误")
		appG.FailResponse(http.StatusBadRequest, e.ErrRetrieveProductFail, existErr)
		return
	}

	p, err := models.ProductCreate(jsonData)
	if err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrCreateProductFail, err)
		return
	}

	appG.Response(http.StatusOK, p)
}

// ProductUpdate godoc
// @Summary Update a product
// @Description Update a product by json product
// @Tags Product
// @Accept json
// @Produce json
// @Param pk path string true "pk is ProductCode"
// @Param product body models.Product true "partial update product"
// @Success 200 {object} app.JSONResult
// @Failure 400 {object} app.FailJSONResult
// @Failure 404 {object} app.FailJSONResult
// @Router /api/v1/products/{pk}  [patch]
func ProductUpdate(c *gin.Context) {
	pk := c.Param("pk")
	appG := app.Gin{C: c}

	if isExist := models.ExistProduct(pk); isExist == false {
		existErr := errors.New("不存在的主键或未知错误")
		appG.FailResponse(http.StatusNotFound, e.ErrUpdateProductFail, existErr)
		return
	}

	var jsonData map[string]interface{}
	decoder := json.NewDecoder(c.Request.Body)

	if err := decoder.Decode(&jsonData); err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrUpdateProductFail, err)
		return
	}

	p, err := models.ProductUpdate(pk, jsonData)
	if err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrUpdateProductFail, err)
		return
	}

	appG.Response(http.StatusOK, p)
}

// ProductDestroy godoc
// @Summary Destroy a product
// @Description Destroy a product by pk
// @Tags Product
// @Accept json
// @Produce json
// @Param pk path string true "pk is ProductCode"
// @Success 204 {object} app.JSONResult
// @Failure 404 {object} app.FailJSONResult
// @Router /api/v1/products/{pk}  [delete]
func ProductDestroy(c *gin.Context) {
	pk := c.Param("pk")
	appG := app.Gin{C: c}

	if isExist := models.ExistProduct(pk); isExist == false {
		existErr := errors.New("不存在的主键或未知错误")
		appG.FailResponse(http.StatusNotFound, e.ErrDestroyProductFail, existErr)
		return
	}

	if err := models.ProductDestroy(pk); err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrDestroyProductFail, err)
		return
	}

	appG.Response(http.StatusNoContent, nil)
}

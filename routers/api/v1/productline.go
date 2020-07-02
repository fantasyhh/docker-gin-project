package v1

import (
	"encoding/json"
	"errors"
	"github.com/fantasyhh/drizzle/models"
	"github.com/fantasyhh/drizzle/pkg/app"
	"github.com/fantasyhh/drizzle/pkg/e"
	"github.com/fantasyhh/drizzle/pkg/setting"
	"github.com/fantasyhh/drizzle/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ProductLineList godoc
// @Summary List productLines
// @Description get all productLines
// @Tags ProductLine
// @Accept json
// @Produce json
// @Param page query integer false "page num"
// @Success 200 {object} app.JSONResult
// @Failure 400 {object} app.FailJSONResult
// @Router /api/v1/productlines  [get]
func ProductLineList(c *gin.Context) {
	appG := app.Gin{C: c}

	products, err := models.ProductLineList(util.GetPage(c), setting.AppSetting.PageSize)
	if err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrListProductLineFail, err)
		return
	}

	appG.Response(http.StatusOK, products)

}

// ProductLineRetrieve godoc
// @Summary Retrieve a productLine
// @Description Retrieve a productLine by pk
// @Tags ProductLine
// @Accept json
// @Produce json
// @Param pk path string true "pk is ProductLine"
// @Success 200 {object} app.JSONResult
// @Failure 400 {object} app.FailJSONResult
// @Failure 404 {object} app.FailJSONResult
// @Router /api/v1/productlines/{pk}  [get]
func ProductLineRetrieve(c *gin.Context) {
	// pk == primary key
	pk := c.Param("pk")
	appG := app.Gin{C: c}

	if isExist := models.ExistProductLine(pk); isExist == false {
		existErr := errors.New("不存在的主键或未知错误")
		appG.FailResponse(http.StatusNotFound, e.ErrRetrieveProductLineFail, existErr)
		return
	}

	product, err := models.ProductLineRetrieve(pk)
	if err != nil {
		appG.FailResponse(http.StatusNotFound, e.ErrRetrieveProductLineFail, err)
		return
	}

	appG.Response(http.StatusOK, product)
}

// ProductLineCreate godoc
// @Summary Create a productLine
// @Description Create one productLine by json productLine
// @Tags ProductLine
// @Accept json
// @Produce json
// @Param product body models.ProductLine  true "create productLine"
// @Success 200 {object} app.JSONResult
// @Failure 400 {object} app.FailJSONResult
// @Router /api/v1/productlines  [post]
func ProductLineCreate(c *gin.Context) {
	appG := app.Gin{C: c}
	var jsonData models.ProductLine

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrCreateProductLineFail, err)
		//appG.FailResponse(http.StatusBadRequest, err.Error())
		return
	}

	if isExist := models.ExistProductLine(jsonData.ProductLine); isExist == true {
		existErr := errors.New("已经存在相同主键无法创建或未知错误")
		appG.FailResponse(http.StatusBadRequest, e.ErrCreateProductLineFail, existErr)
		return
	}

	p, err := models.ProductLineCreate(jsonData)
	if err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrCreateProductLineFail, err)
		//appG.FailResponse(http.StatusBadRequest, err.Error())
		return
	}

	appG.Response(http.StatusOK, p)
}

// ProductLineUpdate godoc
// @Summary Update a productLine
// @Description Update a productLine by json productLine
// @Tags ProductLine
// @Accept json
// @Produce json
// @Param pk path string true "pk is ProductLine"
// @Param product body models.ProductLine true "partial update productLine"
// @Success 200 {object} app.JSONResult
// @Failure 400 {object} app.FailJSONResult
// @Failure 404 {object} app.FailJSONResult
// @Router /api/v1/productlines/{pk}  [patch]
func ProductLineUpdate(c *gin.Context) {
	pk := c.Param("pk")
	appG := app.Gin{C: c}

	if isExist := models.ExistProductLine(pk); isExist == false {
		existErr := errors.New("不存在的主键或未知错误")
		appG.FailResponse(http.StatusNotFound, e.ErrUpdateProductLineFail, existErr)
		return
	}

	var jsonData map[string]interface{}
	decoder := json.NewDecoder(c.Request.Body)

	if err := decoder.Decode(&jsonData); err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrUpdateProductLineFail, err)
		return
	}

	p, err := models.ProductLineUpdate(pk, jsonData)
	if err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrUpdateProductLineFail, err)
		return
	}

	appG.Response(http.StatusOK, p)
}

// ProductLineDestroy godoc
// @Summary Destroy a productLine
// @Description Destroy a productLine by pk
// @Tags ProductLine
// @Accept json
// @Produce json
// @Param pk path string true "pk is ProductLine"
// @Success 204 {object} app.JSONResult
// @Failure 404 {object} app.FailJSONResult
// @Router /api/v1/productlines/{pk}  [delete]
func ProductLineDestroy(c *gin.Context) {
	pk := c.Param("pk")
	appG := app.Gin{C: c}

	if isExist := models.ExistProductLine(pk); isExist == false {
		existErr := errors.New("不存在的主键或未知错误")
		appG.FailResponse(http.StatusNotFound, e.ErrDestroyProductLineFail, existErr)
		return
	}

	if err := models.ProductLineDestroy(pk); err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrDestroyProductLineFail, err)
		return
	}

	appG.Response(http.StatusNoContent, nil)
}

// ProductInLineList godoc
// @Summary Product In Line List
// @DescriptionGet all products in a productLine
// @Tags ProductLine
// @Accept json
// @Produce json
// @Param pk path string true "pk is ProductLine"
// @Success 200 {object} app.JSONResult
// @Failure 400 {object} app.FailJSONResult
// @Failure 404 {object} app.FailJSONResult
// @Router /api/v1/productlines/{pk}/products  [get]
func ProductInLineList(c *gin.Context) {
	pk := c.Param("pk")
	appG := app.Gin{C: c}

	if isExist := models.ExistProductLine(pk); isExist == false {
		existErr := errors.New("不存在的主键或未知错误")
		appG.FailResponse(http.StatusNotFound, e.ErrDestroyProductLineFail, existErr)
		return
	}

	products, err := models.ProductInLine(pk)
	if err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrListProductInLineFail, err)
		return
	}

	appG.Response(http.StatusOK, products)
}

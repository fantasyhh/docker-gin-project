package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fantasyhh/drizzle/pkg/app"
	"github.com/fantasyhh/drizzle/pkg/e"
	"github.com/fantasyhh/drizzle/pkg/upload"
)
// UploadImage represent upload image handler
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}

	//f, err := fhs[0].Open() ; return f, fhs[0], err ===>  file == image.Open()
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrUploadFail, err)
		return
	}

	if image == nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrUploadFail, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		err = errors.New("上传文件格式不准确")
		appG.FailResponse(http.StatusBadRequest, e.ErrUploadFail, err)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrUploadFail, err)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		appG.FailResponse(http.StatusBadRequest, e.ErrUploadFail, err)
		return
	}

	// https://stackoverflow.com/questions/4083702/posting-a-file-and-associated-data-to-a-restful-webservice-preferably-as-json
	// 返回了image_url之后将此属性传入productline的body json中做到关联
	appG.Response(http.StatusOK, map[string]string{
		"image_url":      upload.GetImageFullURL(imageName),
		"image_save_url": savePath + imageName,
	})
}

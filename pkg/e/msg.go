package e

import "errors"

var (
	// ErrListProductFail indicates fail to list product
	ErrListProductFail = errors.New("展示所有产品失败")
	// ErrRetrieveProductFail indicates fail to retrieve product
	ErrRetrieveProductFail = errors.New("取回单个产品失败")
	// ErrCreateProductFail indicates fail to create product
	ErrCreateProductFail = errors.New("创建产品失败")
	// ErrUpdateProductFail indicates fail to update product
	ErrUpdateProductFail = errors.New("更新产品失败")
	// ErrDestroyProductFail indicates fail to destroy product
	ErrDestroyProductFail = errors.New("删除产品失败")

	// ErrListProductLineFail indicates fail to list productline
	ErrListProductLineFail = errors.New("展示所有产线失败")
	// ErrRetrieveProductLineFail indicates fail to retrieve productline
	ErrRetrieveProductLineFail = errors.New("取回单条产线失败")
	// ErrCreateProductLineFail indicates fail to create productline
	ErrCreateProductLineFail = errors.New("创建产线失败")
	// ErrUpdateProductLineFail indicates fail to update productline
	ErrUpdateProductLineFail = errors.New("更新产线失败")
	// ErrDestroyProductLineFail indicates fail to destroy productline
	ErrDestroyProductLineFail = errors.New("删除产线失败")
	// ErrListProductInLineFail indicates fail to list all products productline
	ErrListProductInLineFail = errors.New("获取产线中的所有产品失败")

	// ErrAuthFail indicates fail to login
	ErrAuthFail = errors.New("用户认证失败")

	// ErrUploadFail indicates fail to upload file
	ErrUploadFail = errors.New("上传文件失败")

	//ErrUnknownError = errors.New("未知错误")
)

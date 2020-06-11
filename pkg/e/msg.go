package e

// 一些错误短语
const (
	ListProductFail     = "ListProductFail"
	RetrieveProductFail = "RetrieveProductFail"
	CreateProductFail   = "CreateProductFail"
	UpdateProductFail   = "UpdateProductFail"
	DestroyProductFail  = "DestroyProductFail"

	ListProductLineFail     = "ListProductLineFail"
	RetrieveProductLineFail = "RetrieveProductLineFail"
	CreateProductLineFail   = "CreateProductLineFail"
	UpdateProductLineFail   = "UpdateProductLineFail"
	DestroyProductLineFail  = "DestroyProductLineFail"
	ListProductInLineFail   = "ListProductInLineFail"

	AuthFail      = "AuthFail"
	LoginSuccess  = "LoginSuccess"
	LogoutSuccess = "LogoutSuccess"

	UploadFail    = "UploadFail"
	UnknownError  = "UnknownError"
	CommonSuccess = "Ok"
)

// FlagMsgs 表示错误短语对应的信息
var FlagMsgs = map[string]string{
	ListProductFail:     "展示所有产品失败",
	RetrieveProductFail: "取回单个产品失败",
	CreateProductFail:   "创建产品失败",
	UpdateProductFail:   "更新产品失败",
	DestroyProductFail:  "删除产品失败",

	ListProductLineFail:     "展示所有产线失败",
	RetrieveProductLineFail: "取回单条产线失败",
	CreateProductLineFail:   "创建产线失败",
	UpdateProductLineFail:   "更新产线失败",
	DestroyProductLineFail:  "删除产线失败",
	ListProductInLineFail:   "获取产线中的所有产品失败",

	AuthFail:      "用户认证失败",
	LoginSuccess:  "用户登录成功",
	LogoutSuccess: "用户注销成功",

	UploadFail:    "上传文件失败",
	UnknownError:  "未知错误",
	CommonSuccess: "成功",
}

// GetMsg get error information based on Code
func GetMsg(flag string) string {
	msg, ok := FlagMsgs[flag]
	if ok {
		return msg
	}

	return FlagMsgs["UnknownError"]
}

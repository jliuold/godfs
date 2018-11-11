package godfs

// 服务响应状态码
const (
	OkStat            = 200
	ErrorParamStat    = 600
	InternalErrorStat = 500
)

// 执行成功状态响应
var OK = &Resp{
	Status: OkStat,
	Msg:    "OK",
}

// 参数错误状态响应
var ErrorParam = &Resp{
	Status: ErrorParamStat,
	Msg:    "param wrong",
}

// 内部错误状态响应
var InternalError = &Resp{
	Status: InternalErrorStat,
	Msg:    "internal error",
}

package constants

import (
	myerrors "logistics/libs/errors"
)

// 0-199	预留
// 200-299  参数错误
// 300-399	类型错误
// 400-499  中间件问题
// 500-599  内部错误
// 600-699  操作系统错误
// 700-699  协议错误
// 10000-19999 业务错误

var (
	// ErrUnsupportedContentType indicates unacceptable or lack of Content-Type
	ErrUnsupportedContentType = myerrors.New("unsupported content type", 300)

	// ErrInvalidQueryParams indicates invalid query parameters
	ErrInvalidQueryParams = myerrors.New("invalid query parameters", 200)

	// ErrNotFoundParam indicates that the parameter was not found in the query
	ErrNotFoundParam = myerrors.New("parameter not found in the query", 201)

	// ErrMalformedEntity indicates a malformed entity specification
	ErrMalformedEntity = myerrors.New("malformed entity specification", 500)
	ErrCompressData    = myerrors.New("数据压缩错误", 501)

	ErrRedisConnection    = myerrors.New("Redis连接错误", 400)
	ErrRedisConfig        = myerrors.New("Redis配置错误", 401)
	ErrDatabaseConnect    = myerrors.New("数据库连接错误", 403)
	ErrDatabaseSessionNil = myerrors.New("数据库会话为空", 405)
	ErrDatabaseCommit     = myerrors.New("数据库提交错误", 406)
	ErrDataMarshal        = myerrors.New("数据序列化错误", 408)
	ErrDataUnmarshal      = myerrors.New("数据反序列化错误", 409)
	ErrAwsInit            = myerrors.New("上传文件初始化错误", 410)
	ErrDatabaseQuery      = myerrors.New("数据库查询错误", 411)
	ErrDatabaseInsert     = myerrors.New("数据库插入错误", 412)
	ErrDatabaseUpdate     = myerrors.New("数据库更新错误", 413)
	ErrorDataConvert      = myerrors.New("数据转换错误", 414)
)

var (
	ErrAuthConnection = myerrors.New("Auth Grpc 连接错误", 404)
	ErrCloseTracer    = myerrors.New("关闭Tracer错误", 407)
	ErrTimeParse      = myerrors.New("时间解析错误", 408)
)

var (
	ErrorInternalServerError = myerrors.New("内部服务器错误", 500)
)
var (
	ErrorSystemBusy               = myerrors.New("系统繁忙,请稍后再试!", 10000)
	ErrorParam                    = myerrors.New("参数错误", 10001)
	ErrorCode                     = myerrors.New("验证码错误", 10002)
	ErrorUserCreate               = myerrors.New("用户创建失败", 10003)
	ErrorUserNotFound             = myerrors.New("用户不存在", 10004)
	ErrorPassword                 = myerrors.New("密码错误", 10005)
	ErrorDataBindError            = myerrors.New("数据绑定错误", 10006)
	ErrorWxCode2Session           = myerrors.New("微信Code转换Session错误", 10007)
	ErrorWxAuthFailed             = myerrors.New("微信授权失败", 10008)
	ErrorInvalidRequest           = myerrors.New("参数验证错误", 10009)
	ErrorPasswordSalt             = myerrors.New("密码加密错误", 10010)
	ErrorAccountIdMissing         = myerrors.New("账号ID错误", 10011)
	ErrorUserHasCorporation       = myerrors.New("用户已存在企业", 10012)
	ErrorUserHasNoCorporation     = myerrors.New("用户没有企业", 10013)
	ErrorContactAlreadyExists     = myerrors.New("联系人已存在", 10014)
	ErrorUserAlreadyExists        = myerrors.New("用户已存在", 10015)
	ErrorUserNotExistsCorporation = myerrors.New("用户不属于该企业", 10016)
	ErrorWxDecryptFailed          = myerrors.New("微信解密失败", 10017)
	ErrorRoleMissing              = myerrors.New("角色信息错误", 10018)
	ErrorPermissionDenied         = myerrors.New("权限不足", 10019)
	ErrorUserNotAudited           = myerrors.New("用户未审核", 10020)
	ErrorOrderIsRelease           = myerrors.New("订单已发布", 10021)
	ErrorOrderNotExist            = myerrors.New("订单不存在", 10022)
	ErrorOrderUnOpen              = myerrors.New("订单未开标", 10023)
	ErrorOrderHasQuote            = myerrors.New("订单已标价", 10024)
	ErrorOrderStateErr            = myerrors.New("订单状态错误", 10025)
	ErrorTokenLoginErr            = myerrors.New("", 10026)
	ErrorAddressErr               = myerrors.New("地址错误", 10027)
)

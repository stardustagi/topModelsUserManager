package constants

var (
	WechatOpenIDKey      string = "wechat:openid"
	UserDefaultPassword  string = "FsNyBfcpdtq1p0063MBu"
	ObtenationIterations int    = 3
	WeChatKeyExpire      int    = 7200 // 微信key过期时间，单位秒
	WebHost              string = "127.0.0.1"
	WebPort              int    = 10001
)

type RoleType int

const (
	RoleTypeService RoleType = iota
	RoleTypeSeller
)

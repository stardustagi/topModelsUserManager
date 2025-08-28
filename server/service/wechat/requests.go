package wechat

type AuthRequest struct {
	JsCode string `json:"js_code" validate:"required"`
}

type JwtTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

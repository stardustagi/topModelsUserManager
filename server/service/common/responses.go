package common

type BaseResponse struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type BasePageResponse struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Page    PageResp    `json:"page"`
}

type PageResp struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Sort     string `json:"sort"`
	Total    int64  `json:"total"`
}

type GeocoderAddress struct {
	Province string  `json:"province"`
	City     string  `json:"city"`
	District string  `json:"district"`
	Address  string  `json:"address"`
	Lng      float64 `json:"lng"`
	Lat      float64 `json:"lat"`
}

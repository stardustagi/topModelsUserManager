package users

type AddAddressBookRequest struct {
	Id          int64  `json:"id"`                               // 编辑的地址，带上ID
	Name        string `json:"name" validate:"required"`         // 姓名
	Phone       string `json:"phone" validate:"required"`        // 手机号
	Province    string `json:"province" validate:"required"`     //省
	City        string `json:"city" validate:"required"`         //市
	District    string `json:"district" validate:"required"`     //区
	Stall       string `json:"stall"`                            //商贸城摊位号
	Address     string `json:"address" validate:"required"`      // 详细地址
	AddressType int    `json:"address_type" validate:"required"` //地址类型
}

type GetAddressBookRequest struct {
	AddressType int `json:"address_type" validate:"required"` //地址类型
}

type DeleteAddressBookRequest struct {
	Id int64 `json:"id"` // 地址ID
}

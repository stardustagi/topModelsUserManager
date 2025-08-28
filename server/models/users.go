package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Users struct {
	Id              int64  `json:"id" xorm:"'id' pk autoincr BIGINT"`
	Name            string `json:"name" xorm:"'name' comment('账号名') VARCHAR(64)"`
	UnionId         string `json:"union_id" xorm:"'union_id' NULL comment('微信union_id') VARCHAR(255) DEFAULT NULL"`
	OpenId          string `json:"open_id" xorm:"'open_id' comment('微信open_id') unique VARCHAR(255) DEFAULT NULL"`
	PhoneNumber     string `json:"phone_number" xorm:"'phone_number' not null comment('移动电话') VARCHAR(24)"`
	PurePhoneNumber string `json:"pure_phone_number" xorm:"'pure_phone_number' VARCHAR(64)"`
	CreatedAt       int64  `json:"created_at" xorm:"'created_at' BIGINT"`
	UpdateAt        int64  `json:"update_at" xorm:"'update_at' BIGINT"`
	AppId           string `json:"app_id" xorm:"'app_id' comment('app_id') VARCHAR(64)"`
	Password        string `json:"password" xorm:"'password' comment('用户密码') VARCHAR(128)"`
	Salt            string `json:"salt" xorm:"'salt' VARCHAR(16)"`
	LastLoginIp     string `json:"last_login_ip" xorm:"'last_login_ip' VARCHAR(24)"`
	LastLoginTime   int64  `json:"last_login_time" xorm:"'last_login_time' BIGINT"`
}

func (o *Users) TableName() string {
	return "users"
}

func (o *Users) GetSliceName(slice string) string {
	return fmt.Sprintf("users_%s", slice)
}

func (o *Users) GetSliceDateMonthTable() string {
	t := time.Now()
	return fmt.Sprintf("users_%d%02d", t.Year(), t.Month())
}

func (o *Users) GetSliceDateDayTable() string {
	t := time.Now()
	return fmt.Sprintf("users_%d%02d%02d", t.Year(), t.Month(), t.Day())
}

func (o *Users) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}

func (o *Users) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}

func (o *Users) PrimaryKey() interface{} {
	return o.Id
}

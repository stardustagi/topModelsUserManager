package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Migrations struct {
	Id string `json:"id" xorm:"'id' not null pk VARCHAR(255)"`
}

func (o *Migrations) TableName() string {
	return "migrations"
}

func (o *Migrations) GetSliceName(slice string) string {
	return fmt.Sprintf("migrations_%s", slice)
}

func (o *Migrations) GetSliceDateMonthTable() string {
	t := time.Now()
	return fmt.Sprintf("migrations_%d%02d", t.Year(), t.Month())
}

func (o *Migrations) GetSliceDateDayTable() string {
	t := time.Now()
	return fmt.Sprintf("migrations_%d%02d%02d", t.Year(), t.Month(), t.Day())
}

func (o *Migrations) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}

func (o *Migrations) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}

func (o *Migrations) PrimaryKey() interface{} {
	return o.Id
}

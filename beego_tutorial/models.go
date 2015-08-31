package main

import (
	"github.com/astaxie/beego/orm"
)

// Rate represent rate in database
type Rate struct {
	Id       int
	Currency string
	Rate     float64
}

func (r *Rate) TableName() string {
	return "rates"
}

// RateSource represent Rate source in database
type RateSource struct {
	Id       int
	Name     string
	Currency string
	Parent   *RateSource `orm:"rel(fk)"`
}

func (rs *RateSource) TableName() string {
	return "rate_sources"
}

func init() {
	orm.RegisterModel(new(Rate), new(RateSource))
}

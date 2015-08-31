package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DR_Postgres)

	orm.RegisterDataBase("default", "postgres", "postgres://ces:ces@localhost:5432/ces?sslmode=disable")
}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	rateSource := RateSource{Id: 217}
	err := o.Read(&rateSource)

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rateSource.Id, rateSource.Currency, rateSource.Name)

	}
}

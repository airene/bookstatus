package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表模型
type User struct {
	Id       int64
	Username string    `orm:"size(20)"`
	Password string    `orm:"unique;size(100)"`
	Userrole string    `orm:"size(40)"`
	Regtime  time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *User) TableName() string {
	return TableName("user")
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {

		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

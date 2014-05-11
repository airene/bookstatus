package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表模型
type Book struct {
	Id          int64
	Bookcode    string    `orm:"size(20)"`
	Bookname    string    `orm:"unique;size(100)"`
	Borrowby    string    `orm:"size(12)"`
	Borrowtime  time.Time `orm:"auto_now_add;type(datetime)"`
	Status      int8
	Borrowcount int64
}

func (m *Book) TableName() string {
	return TableName("book")
}

func (m *Book) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Book) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Book) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Book) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Book) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

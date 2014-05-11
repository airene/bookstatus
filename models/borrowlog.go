package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表模型
type BookBorrowLog struct {
	Id          int64
	Bookid		int64
	Borrowby	string      `orm:"size(20)"`
	Borrowtime  time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *BookBorrowLog) TableName() string {
	return TableName("borrowlog")
}

func (m *BookBorrowLog) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *BookBorrowLog) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *BookBorrowLog) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *BookBorrowLog) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *BookBorrowLog) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

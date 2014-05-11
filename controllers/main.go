package controllers

import (
	"github.com/airene/bookstatus/models"
	"github.com/astaxie/beego"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	// 可借图书
	var bookList []*models.Book
	query := new(models.Book).Query().Filter("status", 0)
	count, _ := query.Count()
	if count > 0 {
		query.All(&bookList)
	}

	// 已借图书
	var borrowedList []*models.Book
	query2 := new(models.Book).Query().Filter("status", 1)
	count2, _ := query2.Count()
	if count2 > 0 {
		query2.All(&borrowedList)
	}
	this.Data["bookList"] = bookList
	this.Data["borrowedList"] = borrowedList
	this.TplNames = "index.html"
}

func (this *MainController) Add() {
	this.TplNames = "add.html"
}

func (this *MainController) Save() {
	if this.Ctx.Request.Method == "POST" {
		bookcode := strings.TrimSpace(this.GetString("bookcode"))
		bookname := strings.TrimSpace(this.GetString("bookname"))
		booktype := strings.TrimSpace(this.GetString("booktype"))

		var book models.Book
		book.Bookcode = bookcode
		book.Bookname = bookname
		book.Booktype = booktype
		if err := book.Insert(); err != nil {
			//
		}
		this.Redirect("/add", 302)
	}
}

func (this *MainController) Borrow() {
	id, _ := this.GetInt("id")
	book := models.Book{Id: id}
	if err := book.Read(); err != nil {
		//
	}
	this.Data["book"] = book
	this.TplNames = "borrow.html"
}

func (this *MainController) BorrowSave() {
	bookid, _ := this.GetInt("id")
	username := strings.TrimSpace(this.GetString("username"))
	book := models.Book{Id: bookid}
	if err := book.Read(); err != nil {
		//
	}
	book.Borrowby = username
	book.Status = 1
	book.Borrowcount += 1
	book.Update()

	var borrowlog models.BookBorrowLog
	borrowlog.Bookid = bookid
	borrowlog.Borrowby = username
	if err := borrowlog.Insert(); err != nil {
		//
	}
	this.Redirect("/", 302)
}

func (this *MainController) Back() {
	bookid, _ := this.GetInt("id")
	book := models.Book{Id: bookid}
	if err := book.Read(); err != nil {
		//
	}
	book.Borrowby = ""
	book.Status = 0
	book.Update()
	this.Redirect("/", 302)
}

package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//跟数据库相关的操作

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return
}

//查数据库数据
func queryAllBook() (bookList []*Book, err error) {
	sqlStr := "select id,title,price from book;"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("查询所有信息失败")
		return
	}
	return
}

func queryBookByID(id int64) (book Book, err error) {
	//跨表查询数据
	//1.直接使用SQL语句组合查询
	//2.一张张表查，查完使用代码把他们组合起来
	sqlStr := `select book.id as book_id,book.title,book.price,publisher.id as publisher_id,publisher.province,publisher.city 
				from book 
				join publisher 
				on book.publisher_id=publisher.id 
				where book.id=?`
	err = db.Get(&book, sqlStr, id)
	if err != nil {
		fmt.Println("查询书籍信息失败")
		return
	}
	return
}

//插入数据
func insertBook(title string, price float64) (err error) {
	sqlStr := "insert into book(title,price) values(?,?);"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("插入书籍失败！")
		return
	}
	return
}

//删除数据
func deleteBook(id int64) (err error) {
	sqlStr := "delete from book where id=?;"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除书籍失败！")
		return
	}
	return
}

//编辑数据
func editBook(title string, price float64, id int64) (err error) {
	sqlStr := "update book set title=?,price=? where id=?"
	_, err = db.Exec(sqlStr, title, price, id)
	if err != nil {
		fmt.Println("编辑书籍信息失败")
		return
	}
	return
}

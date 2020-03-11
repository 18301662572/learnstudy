package main

//专门用来定义与数据对应的结构体

//Book 数据库对应的结构体
type Book struct {
	ID    int64   `db:"book_id"`
	Title string  `db:"title"`
	Price float64 `db:"price"`
	Publisher
}

//Publisher 数据库对应的结构体
type Publisher struct {
	ID       int64  `db:"publisher_id"`
	Province string `db:"province"`
	City     string `db:"city"`
}

package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//数据库指针
var DB *sqlx.DB
var DBTest *sqlx.DB

func InitDB() *sqlx.DB {
	if DB != nil {
		return DB
	}
	// 数据源语法："用户名:密码@[连接方式](主机名:端口号)/数据库名"
	database, err := sqlx.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/welfare?parseTime=true")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	DB = database
	return DB
}

func InitDBTest() *sqlx.DB {
	if DBTest != nil {
		return DBTest
	}
	// 数据源语法："用户名:密码@[连接方式](主机名:端口号)/数据库名"
	database, err := sqlx.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/welfare_test?parseTime=true")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	DBTest = database
	return DBTest
}

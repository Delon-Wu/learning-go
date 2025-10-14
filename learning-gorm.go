package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// go get -u gorm.io/gorm
// 数据库可选：
// 1. go get -u gorm.io/driver/sqlite
// 2. go get -u gorm.io/driver/mysql

// User 带 * 的指针，当为空是默认值是nil
type User struct {
	ID           uint    `gorm:"primarykey"`
	Name         string  `gorm:"size:255"`
	Email        *string `gorm:"size:255"`
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ignored      string
}

type Member struct {
	gorm.Model
	Name string `gorm:"size:255"`
	Age  uint8
}

type Author struct {
	Name  string `gorm:"size:255"`
	Email string `gorm:"size:255"`
}

type Blog struct {
	Author  `gorm:"foreignkey:AuthorID"`
	ID      uint64 `gorm:"primarykey"`
	Upvotes int32  `gorm:"default:0"`
}

type Blog2 struct {
	ID      uint64 `gorm:"primarykey"`
	Upvotes int32  `gorm:"default:0"`
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func main() {
	//db, err := gorm.Open(mysql.Open("root:st123456@tcp（这是用户名，密码）(127.0.0.1:3306)/gorm(这是数据库名)?charset=utf8（这是编码格式）&parseTime=True（将golang的time转成数据库支持的）&loc=Local"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open("root:st123456@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Run(db)
}

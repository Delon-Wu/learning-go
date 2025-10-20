package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// go get -u gorm.io/gorm
// 数据库可选：
// 1. go get -u gorm.io/driver/sqlite
// 2. go get -u gorm.io/driver/mysql

// windows 系统打开mysql服务cl(以管理员身份运行cmd)：net start mysql80

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

	CompanyID   uint    // 要和Code的数据类型保持一致
	Company     Company `gorm:"foreignKey:CompanyID;references: Code"`
	CreditCard1 CreditCard1
	Language    []Language `gorm:"many2many:user_languages;"`
}

type Company struct {
	gorm.Model
	Code uint   `gorm:"size:255;unique"`
	Name string `gorm:"size:255"`
}

type CreditCard1 struct {
	gorm.Model
	Number string `gorm:"size:255"`
	UserID uint
}

type Language struct {
	gorm.Model
	Name string `gorm:"size:255"`
}

type Member struct {
	gorm.Model
	Name string `gorm:"size:255"`
	Age  uint8
}

type Author struct {
	Name  string `gorm:"size:255;"`
	Email string `gorm:"size:255"`
}

type Blog struct {
	Author  `gorm:"foreignkey:AuthorID"`
	ID      uint64 `gorm:"primarykey"`
	Upvotes int32  `gorm:"default:0"`
}

func (b *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	b.Upvotes = 0
	return nil
}

type Blog2 struct {
	ID      uint64 `gorm:"primarykey"`
	Upvotes int32  `gorm:"default:0"`
}

func Run1(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.Create(&User{
		Name: "John Doe",
		Age:  18,
	})
}

func Run2(db *gorm.DB) {
	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&Company{})
	//db.AutoMigrate(&CreditCard1{})
	//db.AutoMigrate(&Language{})
	//
	//now := time.Now()
	//
	//company := Company{Name: "Johnson", Code: 123888}
	//result := db.Create(&company)
	//if result.Error != nil {
	//	fmt.Println(result.Error)
	//}
	//fmt.Println("Company ID:", company.ID)
	//user := User{Name: "Demon Doe", Age: 90, Birthday: &now, CompanyID: 123888}
	//db.Create(&user)

	//user := User{}
	//db.Preload("Company").First(&user)
	//fmt.Printf("user: %+v\n", user)

	//db.First(&user)
	//fmt.Printf("user: %+v\n", user)
	//cards := []*CreditCard1{{Number: "800012", UserID: user.ID}, {Number: "800013", UserID: user.ID}}
	//db.Create(cards)

	//user := User{}
	//db.Preload("CreditCard1").First(&user)
	//fmt.Println(user)

	//languages := []*Language{{Name: "English"}, {Name: "Chinese"}, {Name: "French"}}
	//db.Create(languages)
	var languages []Language
	db.Find(&languages)
	fmt.Println("All languages: ", languages)
	user := User{Name: "Old John", Age: 199, Language: languages, CompanyID: 123888}
	db.Create(&user)
}

func main() {
	//db, err := gorm.Open(mysql.Open("root:st123456@tcp（这是用户名，密码）(127.0.0.1:3306)/gorm(这是数据库名)?charset=utf8（这是编码格式）&parseTime=True（将golang的time转成数据库支持的）&loc=Local"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	//Run1(db)
	Run2(db)
}

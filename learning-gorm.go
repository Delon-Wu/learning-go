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

// User 带 * 的指针，当为空是默认值是nil
type User struct {
	ID                uint   `gorm:"primarykey"`
	Name              string `gorm:"size:255"`
	Age               uint8
	Birthday          *time.Time
	MemberNumber      sql.NullString
	ActivatedAt       sql.NullTime
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ignored           string
	BillingAddressID  uint
	BillingAddress    Address `gorm:"foreignkey:BillingAddressID"`
	ShippingAddressID uint
	ShippingAddress   Address    `gorm:"foreignkey:ShippingAddressID"`
	Languages         []Language `gorm:"many2many:user_languages;"`
	Emails            []Email
}

type Address struct {
	gorm.Model
	Address1 string
	UserID   uint
}

type Language struct {
	gorm.Model
	Name   string
	UserID uint
}

type Email struct {
	gorm.Model
	Email  string
	UserID uint
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

func Run1(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func Run2(db *gorm.DB) {
	// CURD
	user := User{
		Name: "Delon",
		Age:  18,
	}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println("Error creating user", result.Error)
	}
	fmt.Println("Single user created", user)

	users := []*User{
		{Name: "Tom", Age: 9},
		{Name: "Jack", Age: 17},
		{Name: "Janice", Age: 59},
	}
	result = db.Create(&users)
	if result.Error != nil {
		fmt.Println("Error creating users", result.Error)
	}
	fmt.Println("Users created", result)
}

func Run3(db *gorm.DB) {
	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&Address{})
	//db.AutoMigrate(&Language{})
	//db.AutoMigrate(&Email{})
	//user := User{
	//	Name:            "Chandler",
	//	Age:             25,
	//	BillingAddress:  Address{Address1: "Billing Address1 - Address 1"},
	//	ShippingAddress: Address{Address1: "Shipping Address1 - Address 1"},
	//	Emails: []Email{
	//		{Email: "666666@qq.com"},
	//		{Email: "888@qq.com"},
	//	},
	//	Languages: []Language{
	//		{Name: "English"},
	//		{Name: "Chinese"},
	//		{Name: "French"},
	//	},
	//}
	////创建时自动保存关联
	//db.Create(&user)
	//db.Save(&user)

	//user := User{}
	//db.Find(&user, 1)
	//user.BillingAddress = Address{Address1: "new billing address 2"} // 新增记录
	//user.BillingAddress.Address1 = "new billing address 2" // 更新记录
	//db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

	//db.Preload("BillingAddress").Find(&user, 1)
	//fmt.Println("User: ", user)

	// 完善User的定义（已包含所有关联字段）

	// 删除id为2和3的用户及其所有关联信息
	//var users []User
	//db.Find(&users, []int{2, 3})
	//for _, u := range users {
	//	db.Select(clause.Associations).Delete(&u)
	//}

	var user User
	////error := db.Model(&user).Association("Languages").Error
	//var languages []Language
	//err := db.Model(&user).Association("Languages").Find(&languages)
	//if err != nil {
	//	return
	//}
	//fmt.Println("Languages:\n", languages)

	ids := []uint{0, 1, 2}
	db.Model(&user).Where("ID IN ?", ids).Association("Languages").Clear()
}

func main() {
	//db, err := gorm.Open(mysql.Open("root:st123456@tcp（这是用户名，密码）(127.0.0.1:3306)/gorm(这是数据库名)?charset=utf8（这是编码格式）&parseTime=True（将golang的time转成数据库支持的）&loc=Local"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	//Run1(db)
	//Run2(db)
	Run3(db)
}

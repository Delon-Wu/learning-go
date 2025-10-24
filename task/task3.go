package task

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//题目1：基本CRUD操作
//假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
//要求 ：
//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

type Student struct {
	gorm.Model
	Name  string
	Age   uint8
	Grade string
}

func task3_1_1(db *gorm.DB) {
	if err := db.AutoMigrate(&Student{}); err != nil {
		panic("迁移失败")
	}
	db.Create(&Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	})

	var students []Student
	db.Where("Age > ?", 18).Find(&students)
	fmt.Println("All students older than 18:", students)

	zhangsan := Student{}
	db.Where("name = ?", "张三").First(&zhangsan)
	zhangsan.Grade = "四年级"
	db.Save(&zhangsan)

	db.Create(&Student{
		Name:  "李四2",
		Age:   12,
		Grade: "三年级",
	})
	db.Debug().Where("Age < ?", 15).Delete(&Student{})
	db.Unscoped().Where("age < ?", 15).Find(&students)
	db.Debug().Unscoped().Delete(&students)
}

//题目2：事务语句
//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
//要求 ：
//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

type Account struct {
	gorm.Model
	Balance float64
}
type Transaction struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
}

func TransferMoney(tx *gorm.DB, fromAccountID, toAccountID uint, amount float64) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		fromAccount := Account{}
		if err := tx.First(&fromAccount, fromAccountID).Error; err != nil {
			return err
		}
		if fromAccount.Balance < amount {
			return errors.New("余额不足")
		}
		fromAccount.Balance -= amount
		toAccount := Account{}
		if err := tx.First(&toAccount, toAccountID).Error; err != nil {
			return err
		} else {
			toAccount.Balance += amount
			tx.Save(&fromAccount)
			tx.Save(&toAccount)
			tx.Create(&Transaction{ToAccountID: toAccountID, FromAccountID: fromAccountID, Amount: amount})
		}
		return nil
	})
}

func task3_1_2(db *gorm.DB) {
	db.AutoMigrate(&Transaction{}, &Account{})
	//accountA, accountB := Account{Balance: 1000}, Account{Balance: 0}
	//db.Create([]*Account{&accountA, &accountB})

	if err := TransferMoney(db, 1, 2, 0.99); err != nil {
		fmt.Println("转账失败:", err)
	} else {
		fmt.Println("转账成功")
	}
}

//Sqlx入门
//题目1：使用SQL扩展库进行查询
//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
//要求 ：
//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Employee1 struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Salary     uint      `db:"salary"`
	Department string    `db:"department"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func task3_2_1(db *sqlx.DB) {
	employeeTable := `
    CREATE TABLE IF NOT EXISTS employees (
        id BIGINT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        salary INT,
        department VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`
	if _, err := db.Exec(employeeTable); err != nil {
		fmt.Println("建表失败:", err)
	}

	//db.Exec(
	//	"INSERT INTO employees (name, salary, department) VALUES (?, ?, ?)",
	//	"Wang", 9000, "技术部",
	//)
	//db.Exec(
	//	"INSERT INTO employees (name, salary, department) VALUES (?, ?, ?)",
	//	"Wu", 30000, "技术部",
	//)

	//var employees []Employee1
	//if err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部"); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("技术部人员: ", employees)
	//}

	var employee Employee1
	if err := db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("工资最高的员工:", employee)
	}
}

//题目2：实现类型安全映射
//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
//要求 ：
//定义一个 Book 结构体，包含与 books 表对应的字段。
//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type Book struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	Price     float64   `db:"price"`
	Author    string    `db:"author"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func task3_2_2(db *sqlx.DB) {
	employeeTable := `
	CREATE TABLE IF NOT EXISTS books (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	price FLOAT NOT NULL,
	author VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci
`
	if _, err := db.Exec(employeeTable); err != nil {
		fmt.Println("建表失败:", err)
	}

	//books := []Book{
	//	{Title: "三国演义", Price: 50.1, Author: "罗贯中"},
	//	{Title: "水浒传", Price: 49, Author: "施耐庵"},
	//	{Title: "西游记", Price: 99, Author: "吴承恩"},
	//}
	//query := "INSERT INTO books (title, price, author) VALUES (:title, :price, :author)"
	//_, err := db.NamedExec(query, books)
	//if err != nil {
	//	fmt.Println("批量插入失败: ", err)
	//}
	//fmt.Println("批量插入成功")

	var books1 []Book
	if err := db.Select(&books1, "SELECT * FROM books WHERE price > 50"); err != nil {
		fmt.Println("查询失败", err)
	} else {
		fmt.Println("查询成功,所有价格大于50的书有:", books1)
	}
}

//进阶gorm
//题目1：模型定义
//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//要求 ：
//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//编写Go代码，使用Gorm创建这些模型对应的数据库表。
//题目2：关联查询
//基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息。

type User struct {
	gorm.Model
	Name       string
	Posts      []Post    // 一对多：一个用户有多个文章
	Comments   []Comment // 一对多：一个用户有多个评论
	PostsCount int64     `gorm:"default:0"` // 文章数量统计
}

type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint      // 外键字段
	User          User      // 关联用户
	Comments      []Comment // 一对多：一篇文章有多个评论
	CommentsCount int64     `gorm:"default:0"` // 评论数量统计
	CommentStatus string    `gorm:"default:'无评论'"`
}

type Comment struct {
	gorm.Model
	Content string
	UserID  uint // 外键字段
	User    User // 关联用户
	PostID  uint // 外键字段
	Post    Post // 关联文章
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 如果没有关联用户 ID，则不处理
	if p.UserID == 0 {
		return nil
	}
	// 原子地递增 users.posts_count，使用 UpdateColumn + gorm.Expr 防止触发其他钩子或回调
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		UpdateColumn("posts_count", gorm.Expr("COALESCE(posts_count,0) + ?", 1)).Error
}

func (c *Comment) AfterCreate(tx *gorm.DB) error {
	if c.PostID == 0 {
		return nil
	}
	// 原子地 +1 并更新状态为 "N 条评论"
	return tx.Model(&Post{}).Where("id = ?", c.PostID).
		UpdateColumns(map[string]interface{}{
			"comments_count": gorm.Expr("COALESCE(comments_count,0) + ?", 1),
			"comment_status": gorm.Expr("CONCAT(COALESCE(comments_count,0) + ?, ' 条评论')", 1),
		}).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	if c.PostID == 0 {
		return nil
	}
	// 原子地 -1（不低于0）并更新状态为 "无评论" 或 "N 条评论"
	return tx.Model(&Post{}).Where("id = ?", c.PostID).
		UpdateColumns(map[string]interface{}{
			"comments_count": gorm.Expr("GREATEST(COALESCE(comments_count,0) - ?, 0)", 1),
			"comment_status": gorm.Expr("CASE WHEN GREATEST(COALESCE(comments_count,0) - ?, 0) = 0 THEN '无评论' ELSE CONCAT(GREATEST(COALESCE(comments_count,0) - ?, 0), ' 条评论') END", 1, 1),
		}).Error
}

func task3_3_2(db *gorm.DB) {
	db.AutoMigrate(&Post{}, &Comment{}, &User{})
	//var users = []User{
	//	{Name: "Tom"},
	//	{Name: "Jerry"},
	//	{Name: "Ham"},
	//}
	//db.Create(&users)
	//fmt.Println(users)
	var tom User
	var jerry User
	//var posts []Post
	//var comments []Comment
	db.Where("name = ?", "Tom").First(&tom)
	db.Where("name = ?", "Jerry").First(&jerry)
	//
	//posts = []Post{
	//	{Title: "怎么吃饭", Content: "饭是这样吃的", User: tom},
	//	{Title: "钢铁是怎么养成的", Content: "饭是这样吃的", User: jerry},
	//}
	//db.Create(&posts)
	//db.Save(&posts)

	//var tomPosts []Post
	//db.Model(&Post{}).Where("user_id = ?", tom.ID).Find(&tomPosts)
	//if len(tomPosts) > 0 {
	//	comment := Comment{
	//		Content: "点我头像",
	//		UserID:  tom.ID,
	//		PostID:  tomPosts[0].ID,
	//	}
	//	db.Create(&comment)
	//}

	var hottestPost Post
	var hottestCount int64
	// 查询所有Post及其评论数
	var posts []Post
	db.Preload("Comments").Find(&posts)
	for _, post := range posts {
		if int64(len(post.Comments)) > hottestCount {
			hottestCount = int64(len(post.Comments))
			hottestPost = post
		}
	}
	fmt.Printf("评论最多的Post: %+v, 评论数: %d\n", hottestPost, hottestCount)

}

func task3_3_3(db *gorm.DB) {
	db.AutoMigrate(&Post{}, &Comment{}, &User{})
	var tom User
	var posts []Post
	db.Where("name = ?", "Tom").First(&tom)
	//posts = []Post{
	//	{Title: "怎么吃饭2", Content: "饭是这样吃的2", UserID: tom.ID},
	//	//{Title: "钢铁是怎么养成的", Content: "饭是这样吃的", UserID: jerry.ID},
	//}
	//if err := db.Create(&posts).Error; err != nil {
	//	fmt.Println("创建 posts 失败:", err)
	//}
	//
	//// 重新查询用户以显示 PostsCount（钩子会在创建时递增）
	//db.First(&tom, tom.ID)
	//fmt.Printf("Tom PostsCount=%d,", tom.PostsCount)

	//var comment Comment
	//comment.UserID = tom.ID
	//comment.PostID = 2
	//db.Create(&comment
	var comments []Comment
	db.Model(&Comment{}).Where("post_id = ? AND user_id = ?", 2, tom.ID).Find(&comments)
	db.Preload("Comments").Delete(&comments)

	db.Find(&posts, 2)
	fmt.Println("Post: ", posts)
}

func RunTask3() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	//task3_1_1(db)
	//task3_1_2(db)

	//db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	fmt.Println(err)
	//	panic("failed to connect database")
	//}
	//task3_2_1(db)
	//task3_2_2(db)

	//task3_3_2(db)
	task3_3_3(db)
}

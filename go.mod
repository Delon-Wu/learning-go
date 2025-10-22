module github.com/Delon-Wu/learning-go.git

go 1.25.1

require (
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.31.0
)

//# 下载所有依赖
//go mod download
//
//# 下载依赖并验证校验和
//go mod verify
//
//# 下载依赖并整理 go.mod 文件
//go mod tidy
require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.30.0 // indirect
)

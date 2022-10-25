package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func test1() {
	// session级别的禁用事务
	db.Session(&gorm.Session{SkipDefaultTransaction: true})
}

type User struct {
	gorm.Model
	Name string
}

func test2() {
	// 测试事务操作 没有使用事务控制
	user := User{
		Name: "Tom",
	}
	db.Create(&user)
	db.Create(nil)
}

func test3()  {
	user := User{
		Name: "Tom",
	}
	// 手动事务控制
	begin := db.Begin()
	db.Create(&user)
	err := db.Create(nil).Error
	if err != nil {
		begin.Rollback()
	}else{
		begin.Commit()
	}
}

func test4()  {
	user := User{
		Name: "Tom",
	}
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil{
			return err
		}
		if err := tx.Create(nil).Error; err != nil{
			return err
		}
		return nil
	})
}

func main() {
	//test2()
	//test3()
}

func init() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ // 全局配置
		Logger:                 logger.Default.LogMode(logger.Info),
		DryRun:                 true,
		SkipDefaultTransaction: true, // 禁用全局事务
	})
	if err != nil {
		log.Fatalln(err)
	}
	db = DB
}

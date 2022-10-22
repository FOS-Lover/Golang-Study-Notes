package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
	Active   bool
}

var DB *gorm.DB

func createTable() {
	DB.AutoMigrate(&User{})
}

func insert() {
	user := User{
		Name:     "test",
		Age:      20,
		Birthday: time.Now(),
	}
	DB.Create(&user)
}

func delete1() {
	var user User
	DB.First(&user)
	DB.Delete(&user) // 软删除
}

func delete2() {
	// 根据主键删除
	var user User
	DB.Delete(&user, 4)
}

func delete3() {
	// 批量删除
	var user User
	DB.Where("1 = 1").Delete(&user)
}

// 删除之前
func (receiver *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("before delete...")
	return nil
}
func main() {
	insert()
	//delete1()
	//delete2()
	delete3()
}

func init() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
}

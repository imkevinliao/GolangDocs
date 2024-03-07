package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"myProject/code"
	"path/filepath"
	"strings"
)

type User struct {
	gorm.Model
	ID      uint `gorm:"primary_key;auto_increment"`
	Name    string
	Email   string `gorm:"type:varchar(100);unique_index"`
	Age     uint
	Default string `gorm:"column:DEFAULT_COLUMN"` //通过这种方式控制列名
}

// TableName 使用TableName方法来自定义表名
func (User) TableName() string {
	return "custom_table"
}

func GetDb() *gorm.DB {
	dir, _ := filepath.Abs(`.`)
	dbName := "gorm.db"
	dbPath := filepath.Join(dir, dbName)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!")
	}
	return db
}

func DeleteTable(db *gorm.DB) {
	migrator := db.Migrator()
	if migrator.HasTable(&User{}) {
		err := migrator.DropTable(&User{})
		if err != nil {
			panic("failed to drop table!")
		}
	}
}

func GenerateInfo() (u *User) {
	emailPrefix := code.RandomString(10)
	emailSuffix := "@gmail.com"
	email := emailPrefix + emailSuffix
	age := uint(code.RandomInt(15, 30))
	name := code.RandomString(5)
	userA := &User{
		Name:  name,
		Age:   age,
		Email: email,
	}
	userB := new(User)
	userB.Name = name
	userB.Age = age
	userB.Email = email
	return userA
}

func DbOperation(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic("Database Migrate Wrong!")
	}
	var length = 3
	var UserSlice []*User
	for i := 0; i < length; i++ {
		user := GenerateInfo()
		UserSlice = append(UserSlice, user)
	}
	for _, value := range UserSlice {
		db.Create(value)
	}
	db.Create(&User{Name: "kevin", Email: "kevin@example.com", Age: 25})
	// 找不到满足条件的数据的时候才会创建
	db.FirstOrCreate(&User{}, map[string]interface{}{"name": "Natalia", "email": "Natalia@example.com", "age": 18})

	var user User
	db.First(&user, "name = ?", "kevin")
	db.Model(&user).Update("Email", "kevin_only_change_email@example.com")
	db.Model(&user).Updates(User{Name: "kevin_new", Email: "kevin@example.com"})              // 多个字段更新
	db.Model(&user).Updates(map[string]interface{}{"Name": "kevin update by map", "Age": 30}) // 使用 map 更新多个字段
	//无论软删除还是硬删除，常规的查找都无法找到。软删除可以通过特殊方法找到，硬删除则无解
	db.Delete(&user) //软删除
	var users []User
	db.Find(&users)
	for key, value := range users {
		fmt.Println("index:", key, "id:", value.ID, "name:", value.Name, "email:", value.Email)
	}
	db.Find(&users, "name = ?", "Natalia")
	//硬删除
	db.Unscoped().Delete(&users)
	fmt.Println(strings.Repeat("*", 100))
	db.Find(&users)
	for key, value := range users {
		fmt.Println("index:", key, "id:", value.ID, "name:", value.Name, "email:", value.Email)
	}
	fmt.Println(strings.Repeat("*", 100))
	var alreadyDeleted []User
	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&alreadyDeleted)
	// 打印所有软删除产品信息
	for key, value := range alreadyDeleted {
		fmt.Println("index:", key, "id:", value.ID, "name:", value.Name, "email:", value.Email)
		// 恢复指定的软删除记录
		err := db.Model(&User{}).Unscoped().Where("id = ?", value.ID).Update("deleted_at", nil).Error
		if err != nil {
			fmt.Println("Error while restoring record:", err)
		} else {
			fmt.Println("Record restored successfully")
		}
	}

}

func main() {
	db := GetDb()
	DeleteTable(db)
	DbOperation(db)
}

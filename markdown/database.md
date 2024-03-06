# 数据库
GORM SQLite 驱动使用CGO实现

"github.com/glebarez/sqlite" 使用纯go实现，不依赖CGO，牺牲了性能换取方便性

```golang
package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("error here.")
	}
	type User struct {
		ID   uint   `gorm:"primary_key;auto_increment"`
		Name string `gorm:"not null"`
	}
	db.AutoMigrate(&User{})
	user := User{Name: "John Doe"}
	db.Create(&user)
}
```
# ORM
https://gorm.io/zh_CN/docs/index.html

```golang
package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint `gorm:"primary_key;auto_increment"`
	Name  string
	Email string `gorm:"type:varchar(100);unique_index"`
	Age   uint
}

// TableName 使用TableName方法来自定义表名
func (User) TableName() string {
	return "custom_table"
}

func GetDb() *gorm.DB {
	dbName := "gorm.db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
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

func DbOperation(db *gorm.DB) {
	// GORM 为这些操作内部管理了事务，确保每个操作都是原子性执行的。每次调用如 Create、Save、Delete 等方法时，GORM 都会自动开始一个新的事务，执行操作，并在操作成功完成后提交事务。
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "John", Email: "john@example.com", Age: 25})
	db.Create(&User{Name: "kevin", Email: "kevin@example.com", Age: 25})
	// 找不到满足条件的数据的时候才会创建
	db.FirstOrCreate(&User{}, map[string]interface{}{"name": "Natalia", "email": "Natalia@example.com", "age": 18})
	var user User
	db.First(&user, "name = ?", "kevin") // 查询名字为John的用户
	fmt.Println(user.Name, user.Email, user.Age)

	db.Model(&user).Update("Email", "john_new@example.com")
	db.Model(&user).Updates(User{Name: "John Updated", Email: "john_updated@example.com"})   // 多个字段更新
	db.Model(&user).Updates(map[string]interface{}{"Name": "John Updated Again", "Age": 30}) // 使用 map 更新多个字段
	//软删除
	db.Delete(&user)
	//彻底删除
	//db.Unscoped().Delete(&user)

}

func main() {
	db := GetDb()
	DbOperation(db)
}
```

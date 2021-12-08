package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id       int
	Username string
	Password string
}

func main() {
	dsn := "root:Xb005869@@tcp(127.0.0.1:3306)/My-Golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 自动迁移 - 自动创建表
	//db.AutoMigrate(&user{})

	/*
		插入
	*/
	//user1 := User{1,"123","12"}
	//db.Create(&user1)

	/*
		用指定的字段创建记录保存在指定的表
	*/
	//db.Table("user").Select("username","password").Create(&user1)

	/*
		批量插入
	*/

	//var users = []User{
	//	{8,"xb","444"},
	//	{9,"xxxx","111"},
	//}
	//db.Table("user").Omit("id").Create(&users)

	var user User
	// 获取第一条记录（主键升序）SELECT * FROM users ORDER BY id LIMIT 1;
	db.Table("user").First(&user)
	fmt.Println(user)

	var user2 User
	// 获取一条记录，没有指定排序字段
	take := db.Table("user").Take(&user2)
	fmt.Println(user2, "记录数为:", take.RowsAffected)

	//根据主键id进行查询 SELECT * FROM user WHERE id = 4;
	var user3 User
	db.Table("user").First(&user3, 4)
	fmt.Println(user3)

	var user5 []User
	// SELECT * FROM user WHERE id IN (1,2,3);
	db.Table("user").Find(&user5, []int{1, 2, 3})
	fmt.Println("user5:", user5)

	//查询所有 SELECT * FROM user;
	var user4 []User
	db.Table("user").Find(&user4)
	fmt.Println("user4", user4)

	//如果id是字符串 SELECT * FROM user WHERE id = "2"
	//var user6 User
	//db.Table("user").First(&user6,"id = ?",2)

	var users User
	// 获取第一条匹配的记录 SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	db.Where("name = ?", "jinzhu").First(&user)

	// 获取全部匹配的记录 SELECT * FROM users WHERE name <> 'jinzhu';
	db.Where("name <> ?", "jinzhu").Find(&users)

	// IN SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');
	db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)

	// LIKE SELECT * FROM users WHERE name LIKE '%jin%';
	db.Where("name LIKE ?", "%jin%").Find(&users)

	// AND SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

	// Time SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
	//db.Where("updated_at > ?", lastWeek).Find(&users)

	// BETWEEN SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
	//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)

	/*
		注意：
			当使用结构作为条件查询时，GORM 只会查询非零值字段，如果需要查询零值需要使用map进行查询
	*/
	// Struct	SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
	db.Where(&User{Username: "jinzhu", Password: "20"}).First(&user)

	// Map		SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)

	// 主键切片条件	SELECT * FROM users WHERE id IN (20, 21, 22);
	db.Where([]int64{20, 21, 22}).Find(&users)

	/*
		选择特定字段查询
	*/
	// SELECT name, age FROM users;
	db.Select("name", "age").Find(&users)

	// SELECT name, age FROM users;
	db.Select([]string{"name", "age"}).Find(&users)

	// SELECT COALESCE(age,'42') FROM users;
	db.Table("users").Select("COALESCE(age,?)", 42).Rows()

	/*
		按字段排序
	*/

	// SELECT * FROM users ORDER BY age desc, name;
	db.Order("age desc, name").Find(&users)

	/*
		Limit
	*/

	// SELECT * FROM users LIMIT 3;
	db.Limit(3).Find(&users)

	/*
		去重
	*/
	db.Distinct("name", "age").Order("name, age desc").Find(&users)

	/*
		group by
	*/
	type result struct {
		Date  time.Time
		Total int
	}
	var results result
	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&results)
	// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"
}

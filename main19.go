package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID   uint
	Name string
	//Age        sql.NullInt64 `gorm:"default:1"` // 零值类型
	Age        *int8 `gorm:"default:1"`
	Mobile     string
	Gender     int8
	CreateTime int64 `gorm:"AutoCreateTime"` // 使用时间戳秒数填充创建时间
	UpdateTime int64 `gorm:"AutoUpdateTime"` // 使用时间戳秒数填充更新时间
	//UpdateTime int64 `gorm:"autoUpdateTime:nano"` // 使用时间戳纳秒数填充更新时间
}

//func (u User) TableName() string {
//	return "user"
//}

// CURD查询
func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:123456@(192.168.10.33:3306)/test?charset=utf8mb4"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 创建表
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println("failed to create table")
	}

	// ----------------------------- 创建 -----------------------------
	// 创建记录
	age := int8(18)
	u1 := User{Name: "Shanla", Age: &age, Mobile: "1809990009", Gender: 1}
	u1Created := db.Create(&u1) // 通过数据指针创建记录
	if u1Created.Error != nil && u1Created.RowsAffected <= 0 {
		log.Println("failed to insert record")
		return
	}

	fmt.Printf("u1.id: %d, u1.name : %s, u1.Age : %d, u1.mobile : %s\n", u1.ID, u1.Name, u1.Age, u1.Mobile)

	// 用指定的字段创建记录
	u2 := User{Name: "xiaokang", Age: &age, Mobile: "19889892298"}
	db.Select("Name", "Age").Create(&u2) // 只保存Name和Age
	db.Omit("Name").Create(&u2)          // 忽略Name字段

	// 批量插入记录
	// 一批次批量插入
	var users24 = []User{
		{Name: "张三"},
		{Name: "李四"},
		{Name: "王麻子"},
	}
	db.Create(&users24)
	for _, u := range users24 {
		fmt.Println(u.ID)
	}

	// 多批次批量插入
	var users25 = []User{
		{Name: "Name1"},
		{Name: "Name2"},
		{Name: "Name3"},
		{Name: "Name4"},
		{Name: "Name5"},
		{Name: "Name6"},
		{Name: "Name7"},
		{Name: "Name8"},
		{Name: "Name9"},
		{Name: "Name10"},
	}
	db.CreateInBatches(users25, 5)

	// 插入{Age:0}
	user14 := User{
		//Age:  0, // 如果字段不是指针类型或Scanner/Valuer，通过0，''，false等值会被忽略
		//Age:  sql.NullInt64{0, true},
		Age:  new(int8),
		Name: "测试0值",
	}
	db.Create(&user14)
	return

	// ----------------------------- 查询 -----------------------------
	//检索单个对象
	var user, user1, user2 User
	// SELECT * FROM users ORDER BY id LIMIT 1;
	db.First(&user)
	fmt.Println(user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	db.Last(&user1)
	fmt.Println(user1)
	// SELECT * FROM users LIMIT 1;
	user2Result := db.Take(&user2)
	// 判断数据是否为空
	if errors.Is(user2Result.Error, gorm.ErrRecordNotFound) {
		log.Println("Take record is empty!")
	}
	fmt.Println(user2)

	// 根据主键检索
	var user3 User
	//db.First(&user3, 11)
	db.First(&user3, "11")
	fmt.Println(user3)
	// 主键in
	var users4 []User
	db.Find(&users4, []int{12, 13, 14})
	fmt.Println(users4)

	// 条件
	// 1. string 条件
	var user5 User
	db.Where("name = ?", "Name9").First(&user5) // ==
	fmt.Println(user5)
	var users6, users7, users8, users9 []User
	db.Where("age <> ?", 0).Find(&users6) // !=
	fmt.Println(users6)
	db.Where("gender in ?", []int8{1}).Find(&users7) // IN
	fmt.Println(users7)
	search := "Name"
	db.Where("name like ?", "%"+search+"%").Find(&users8) // like
	fmt.Println(users8)
	db.Where("gender = ? AND age > ?", 1, 18).Find(&users9) // AND
	fmt.Println(users9)
	//db.Where("") // Between

	// SELECT * FROM users WHERE id = 10 and id = 20 ORDER BY id ASC LIMIT 1
	var user10 = User{ID: 10}
	db.Where("id = ?", 20).First(&user10) // record not found

	// 2. Struct & Map 条件
	var user11, user12 User
	var users11, users12 []User
	// Struct 条件
	db.Where(&User{Name: "Name7"}).Find(&user11)
	db.Where(&User{Name: "Shanla"}).Find(&users11)
	fmt.Println(user11)
	fmt.Println(users11)
	// Map 条件
	db.Where(map[string]interface{}{"name": "Shanla", "age": 18}).Find(&user12)
	fmt.Println(user12)
	// Map 主键
	db.Where([]int{1, 2, 3}).Find(&users12)
	fmt.Println(users12)

	// 3. 指定结构体查询字段
	var users13, users14 []User
	var age1 = int8(1)
	db.Debug().Where(&User{Name: "1212", Age: &age1}, "name", "age").Find(&users13)
	fmt.Println(users13)
	db.Debug().Where(&User{Name: "1212"}, "Age").Find(&users14)
	fmt.Println(users14)

	// 4. 内联条件
	var users15 []User
	db.Find(&users15, "name = ?", "1212")
	fmt.Println(users15)

	// 5. Not 条件
	var users16 []User
	db.Debug().Not("name = ?", "1212").Find(&users16)
	fmt.Println(users16)

	// 6. Or 条件
	var users17 []User
	db.Debug().Where("name= ?", "1212").Or("name = ?", "2323").Find(&users17)
	fmt.Println(users17)

	// 选择查询字段
	var users18 []User
	db.Debug().Select("name", "age").Find(&users18)
	fmt.Printf("%#v\n", users18)

	// 排序
	var users19 []User
	db.Debug().Order("id desc").Find(&users19)
	fmt.Println(users19)

	// Limit && Offset
	var users20, users21 []User
	db.Debug().Limit(2).Offset(1).Find(&users20).Limit(1).Find(&users21)
	fmt.Println(users20, users21)

	// Distinct
	var users22 []User
	db.Debug().Distinct("name", "age").Order("name, age desc").Find(&users22)
	fmt.Println(users22)

	// Joins
	var users23 []User
	db.Debug().Model(&User{}).Joins("left join address on address.user_id = users.id").Select("users.name, address.address").Scan(&users23)
	fmt.Printf("result:%#v\n", users23)

	// 智能选择字段
	type ApiUser struct {
		Name string
		Age  int8
	}
	var apiUsers []ApiUser
	// 查询时会自动选择 `name`, `age` 字段
	db.Model(&User{}).Limit(10).Find(&apiUsers)
	for k, apiuser := range apiUsers {
		fmt.Printf("第%d条数据：name:%s, age:%d\n", k+1, apiuser.Name, apiuser.Age)
	}
}

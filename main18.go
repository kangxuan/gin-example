package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Animal struct {
	AnimalId      uint          `gorm:"primaryKey"`                                                       // 定义为主键
	AnimalName    string        `gorm:"type:varchar(100);not null;default:EMPTY STRING;index;comment:名称"` // 定义类型为varchar类型
	AnimalAge     sql.NullInt16 `gorm:"not null;default:0;comment:年龄"`
	AnimalContent string        `gorm:"type:text;not null;comment:说明"`
	CreateTime    uint          `gorm:"AutoCreateTime;not null;default:0;comment:创建时间"` // 自动默认创建时间戳
	UpdateTime    uint          `gorm:"AutoUpdateTime;not null;default:0;comment:更新时间"` // 自动默认更新时间戳
	UpdateAt      time.Time     `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:变更时间"`
}

// TableName 自定义表名
func (animal Animal) TableName() string {
	return "animal2"
}

// 演示模型定义
func main() {
	db, err := gorm.Open(mysql.Open("root:123456@(192.168.10.33:3306)/test?charset=utf8mb4"), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql database")
	}
	_ = db.AutoMigrate(&Animal{})
}

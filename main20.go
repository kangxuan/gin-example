package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Lesson struct {
	LessonId       uint32    `gorm:"primaryKey"`
	LessonName     string    `gorm:"type:varchar(100);not null;default:'';comment:课程名称"`
	LessonLongTime uint8     `gorm:"not null;default:0;comment:课程时长"`
	CreateTime     uint32    `gorm:"AutoCreateTime;not null;comment:创建时间"`
	UpdateTime     uint32    `gorm:"AutoUpdateTime;not null;comment:更新时间"`
	UpdateAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;ON UPDATE CURRENT_TIMESTAMP;comment:变更时间"`
}

func main() {
	dsn := "root:123456@(192.168.10.33:3306)/test?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql")
	}

	_ = db.AutoMigrate(&Lesson{})

	db.Create(&Lesson{
		LessonName:     "英语课",
		LessonLongTime: 40,
	})
	insertLesson := []Lesson{
		{
			LessonName:     "化学课",
			LessonLongTime: 40,
		},
		{
			LessonName:     "物理课",
			LessonLongTime: 40,
		},
		{
			LessonName:     "生物课",
			LessonLongTime: 40,
		},
		{
			LessonName:     "体育课",
			LessonLongTime: 50,
		},
	}
	db.Create(&insertLesson)

	// Save 默认修改所有字段
	var huaXueLesson Lesson
	db.Where("lesson_name = ?", "化学课").First(&huaXueLesson)
	huaXueLesson.LessonLongTime = 30
	db.Save(huaXueLesson)

	// 更新单列
	db.Model(&Lesson{}).Where("lesson_name = ?", "化学课").Update("lesson_name", "雨课堂")

	// 更新多列，当使用 struct 进行更新时，GORM 只会更新非零值的字段。 你可以使用 map 更新字段，或者使用 Select 指定要更新的字段
	db.Debug().Model(&Lesson{}).Where("lesson_long_time = ?", 40).Updates(map[string]interface{}{
		"lesson_long_time": 45,
		"create_time":      0,
	})

	// 更新选定字段 UPDATE `lessons` SET `create_time`=0,`update_time`=1677218430 WHERE create_time = 0
	// 有点坑：实际上没有更新，RowsAffected返回的是查询到的数量
	updated := db.Debug().Model(&Lesson{}).Where("create_time = ?", 0).Select("create_time").Updates(Lesson{LessonName: "课程名称", CreateTime: 0})
	fmt.Println("effected rows:", updated.RowsAffected, ",err:", updated.Error)

	// SQL表达式更新
	db.Model(&Lesson{}).Debug().Where("lesson_long_time = ?", 45).Updates(map[string]interface{}{
		"lesson_long_time": gorm.Expr("lesson_long_time + ?", 5),
	})

	// 删除一条记录
	lesson1 := Lesson{LessonId: 1}
	db.Delete(&lesson1)

	// 根据主键删除记录
	db.Delete(&Lesson{}, 2)
	db.Delete(&Lesson{}, []int{2, 3, 4})

	// 批量删除
	db.Where("create_time = ?", 0).Delete(&Lesson{})
	db.Delete(&Lesson{}, "create_time = ?", 0)
}

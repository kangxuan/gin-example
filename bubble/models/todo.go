package models

import (
	dao2 "gin-example/bubble/dao"
)

type Todo struct {
	ID     int    `gorm:"primaryKey;not null" json:"id"`
	Title  string `gorm:"type:varchar(100);not null;default:'';comment:代办标题" json:"title"`
	Status bool   `json:"status"`
	IsDel  uint8  `gorm:"not null;default:0" json:"is_del"`
}

// GetAllGoto 获取代办列表
func GetAllGoto() (todoList []*Todo, err error) {
	if err := dao2.DB.Where("is_del = ?", 0).Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// AddAGoto 添加代办
func AddAGoto(addTodo *Todo) (err error) {
	if err = dao2.DB.Create(&addTodo).Error; err != nil {
		return
	}
	return nil
}

// GetAGotoById 通过ID获取代办
func GetAGotoById(id string) (upTodo *Todo, err error) {
	upTodo = new(Todo)
	// 先判断todo是否存在
	if err := dao2.DB.Where("id = ? AND is_del = ?", id, 0).First(&upTodo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateAGoto 更新代办
func UpdateAGoto(upTodo *Todo) (err error) {
	if err := dao2.DB.Save(&upTodo).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAGoto 删除代办
func DeleteAGoto(id string) (err error) {
	if err = dao2.DB.Model(&Todo{}).Where("id = ?", id).Update("is_del", 1).Error; err != nil {
		return err
	}
	return nil
}

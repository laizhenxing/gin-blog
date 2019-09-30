package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	// 给予附属属性 json， 在 c.JSON 输出的时候就会自动转换格式
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 获取标签
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// 查询标签的数量
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// 检查标签是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// 新增标签
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

// 新增时回调
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}
// 更新时回调
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

// 更新标签
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}

// 删除标签
func DeleteTag(id int) bool  {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

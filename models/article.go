package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	// index 用户声明这个字段是索引，使用自动迁移会有影响
	// TagID 是一个外键
	TagID int `json:"tag_id" gorm:"index"`
	// Tag 结构体，可以通过 Related 进行关联查询
	Tag Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

// 通过ID查询模型是否存在
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}

	return false
}

// 获取文章的总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

// 获取一篇文章
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	// 查找文章关联的标签
	db.Model(&article).Related(&article.Tag)

	return
}

//新增文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

// 获取所有文章
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	// Preload是预加载器，执行两条sql
	// select * from blog_articles
	// select * from blog_tag whre id in (1,2,3,4)
	// 查询出结构后，将其填充到 Article 的 Tag 中
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

// 修改文章
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

// 删除文章
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ? ",0).Delete(&Article{})

	return true
}

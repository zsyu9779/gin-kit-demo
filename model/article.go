package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	Id            int64  `json:"id"`
	TagID         int    `json:"tag_id" gorm:"index"`
	Tag           Tag    `json:"tag"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

func GetBlogsByTagId(tagId int64) []Article {
	var result []Article
	Db.Raw("select * from blog where id in(select blog_id from where tag_id = ?)",tagId).Scan(&result)
	return result
}

// ExistArticleByID checks if an article exists based on ID
func (article *Article) ExistArticleByID(id int) (bool, error) {
	err := Db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err == nil {
		return true, nil
	} else if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, err
}

// GetArticleTotal gets the total number of articles based on the constraints
func (article *Article) GetArticleTotal() (int, error) {
	var count int
	if err := Db.Model(&article).Where(&article).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticles gets a list of articles based on paging constraints
func (article *Article) GetArticles(pageNum int, pageSize int) ([]*Article, error) {
	var articles []*Article
	err := Db.Preload("Tag").Where(&article).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

// GetArticle Get a single article based on ID
func (article *Article) GetArticle() (*Article, error) {
	err := Db.Where("id = ? AND deleted_on = ? ", article.Id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = Db.Model(article).Related(article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return article, nil
}

// EditArticle modify a single article
func EditArticle(id int, data interface{}) error {
	if err := Db.Model(&Article{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddArticle add a single article
func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}
	if err := Db.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

// DeleteArticle delete a single article
func DeleteArticle(id int) error {
	if err := Db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllArticle clear all article
func CleanAllArticle() error {
	if err := Db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error; err != nil {
		return err
	}

	return nil
}

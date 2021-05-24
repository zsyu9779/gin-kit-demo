package model

import "fmt"

type Article struct {
	Id             int64  `json:"id"`
	Appreciation   bool   `json:"appreciation"`
	Commentable    bool   `json:"commentable"`
	CreateTime     string `json:"create_time"`
	Description    string `json:"description"`
	Title          string `json:"title"`
	FirstPicture   string `json:"first_picture"`
	Flag           string `json:"flag"`
	Published      bool   `json:"published"`
	Recommend      bool   `json:"recommend"`
	ShareStatement bool   `json:"share_statement"`
	UpdateTime     string `json:"update_time"`
	Views          int    `json:"views"`
	UserId         int64  `json:"user_id"`
	TypeId         int64  `json:"type_id"`
	Content        string `json:"content"`
}
type ArticleListItem struct {
	Id             int64  `json:"id"`
	Appreciation   bool   `json:"appreciation"`
	Commentable    bool   `json:"commentable"`
	CreateTime     string `json:"create_time"`
	Description    string `json:"description"`
	Title          string `json:"title"`
	FirstPicture   string `json:"first_picture"`
	Flag           string `json:"flag"`
	Published      bool   `json:"published"`
	Recommend      bool   `json:"recommend"`
	ShareStatement bool   `json:"share_statement"`
	UpdateTime     string `json:"update_time"`
	Views          int    `json:"views"`
	UserId         int64  `json:"user_id"`
	TypeId         int64  `json:"type_id"`
}
type ArticleResp struct {
	Article
	Tags []Tag
	Type Type
	User User
}
type ArticleListResp struct {
	ArticleListItem
	Tags []Tag
	Type Type
	User User
}

func GetBlogsByTagId(tagId int64) []ArticleListResp {
	var result []ArticleListResp
	Db.Raw("select * from blog where id in(select blog_id from where tag_id = ?)", tagId).Find(&result)
	return result
}
func GetDistinctYearFromBlog() []string{
	var result []string
	Db.Table("article").Raw("SELECT DISTINCT DATE_FORMAT(create_time,\"%Y\") as time from t_blog").Find(&result)
	return result
}
func GetBlogsByYear(year int) []ArticleListResp {
	var result []ArticleListResp
	Db.Raw("SELECT * from t_blog where create_time like ?", fmt.Sprintf("%d%",year)).Find(&result)
	return result
}
func (article *Article) AddArticle() error {
	if err := Db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}
func (article *Article) FindArticleById() error {
	if err := Db.Find(&article, article.Id).Error; err != nil {
		return err
	}
	return nil
}
func (article *Article) FindArticleByTypeId() []ArticleListResp {
	var result []ArticleListResp
	if err := Db.Where("type_id <> ?",article.TypeId).Find(&article).Error; err != nil {
		return nil
	}
	return result
}
func (article *Article) FindAllArticles() []ArticleListResp {
	var result []ArticleListResp
	Db.Find(&result)
	return result
}
func (article *Article) FindArticleList(page, pageSize int) []ArticleListItem {
	var result []ArticleListItem
	Db.Scopes(Paginate(page, pageSize)).Find(&result)
	return result
}
func (article *Article) DeleteArticle() bool {
	num := Db.Delete(&article, article.Id).RowsAffected
	if num > 0 {
		return true
	} else {
		return false
	}
}
func (article *Article) UpdateArticle() bool {
	num := Db.Save(&article).RowsAffected
	if num > 0 {
		return true
	} else {
		return false
	}
}

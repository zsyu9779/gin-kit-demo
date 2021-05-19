package service

import (
	"fmt"
	"gin-kit-demo/constant"
	"gin-kit-demo/g_rediscache"
	"gin-kit-demo/model"
	"github.com/sirupsen/logrus"
	"reflect"
	"time"
)

var typeService TypeService
func GetArticleById(articleId int64) model.Article {
	result, _, err := g_rediscache.UseSimpleAop(fmt.Sprintf(constant.BLOG_ARTICLE_BY_ID, articleId), reflect.TypeOf(model.Article{})).
		WithExpires(time.Hour).
		WithEmptyExpires(5 * time.Minute).
		Then(func() (interface{}, error) {
			article := model.Article{}
			article.Id = articleId
			article.FindArticleById()
			return article,nil
		})
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	return result.(model.Article)
}
func GetArticleRespById(articleId int64) model.ArticleResp {
	article := GetArticleById(articleId)
	var result  = model.ArticleResp{
		Article: article,
		Tags:    nil,
		Type:    *typeService.GetType(article.TypeId),
		User:    model.User{},
	}
	return result
}

type ArticleListItem struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	FirstPicture string `json:"firstPicture"`
	Flag         string `json:"flag"`
	TypeId       int64  `json:"type_id"`

}


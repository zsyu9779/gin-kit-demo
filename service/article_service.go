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

func GetArticleById(articleId int64) (model.Article, error) {
	result, _, err := g_rediscache.UseSimpleAop(fmt.Sprintf(constant.BLOG_ARTICLE_BY_ID, articleId), reflect.TypeOf(model.Article{})).
		WithExpires(time.Hour).
		WithEmptyExpires(5 * time.Minute).
		Then(func() (interface{}, error) {
			article := model.Article{}
			article.Id = articleId
			return article.GetArticle()
		})
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	return result.(model.Article), nil
}

type ArticleListItem struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	FirstPicture string `json:"firstPicture"`
	Flag         string `json:"flag"`
	TypeId       int64  `json:"type_id"`

}


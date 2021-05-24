package service

import (
	"fmt"
	"gin-kit-demo/constant"
	"gin-kit-demo/g_rediscache"
	"gin-kit-demo/model"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"time"
)

type ArticleService struct {

}

var typeService TypeService
var userService UserService

func (a *ArticleService)GetArticleById(articleId int64) model.Article {
	result, _, err := g_rediscache.UseSimpleAop(fmt.Sprintf(constant.BLOG_ARTICLE_BY_ID, articleId), reflect.TypeOf(model.Article{})).
		WithExpires(time.Hour).
		WithEmptyExpires(5 * time.Minute).
		Then(func() (interface{}, error) {
			article := model.Article{}
			article.Id = articleId
			article.FindArticleById()
			return article, nil
		})
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	return result.(model.Article)
}
func (a *ArticleService)GetArticleRespById(articleId int64) model.ArticleResp {
	article := a.GetArticleById(articleId)
	var result = model.ArticleResp{
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

func (a *ArticleService)GetArticleListByType(typeId,start,stop int64) []model.ArticleListResp {
	result, _, err := g_rediscache.UseSimpleAop(fmt.Sprintf(constant.BLOG_ARTICLE_BY_TYPE, typeId), reflect.TypeOf([]model.ArticleListResp{})).
		WithExpires(time.Hour).
		WithEmptyExpires(5 * time.Minute).
		Then(func() (interface{}, error) {
			article := model.Article{}
			article.TypeId = typeId
			articleList := article.FindArticleByTypeId()
			for i, item := range articleList {
				articleList[i].Tags = model.GetTagsByArticleId(item.Id)
				articleList[i].User = userService.FindUserById(item.UserId)
				articleList[i].Type = *typeService.GetType(item.TypeId)
			}
			return articleList, nil
		})
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	var articleList = result.([]model.ArticleListResp)
	return articleList[start:stop]
}

func (a *ArticleService)GetArticleListByTag(tagId,start,stop int64) []model.ArticleListResp {
	result, _, err := g_rediscache.UseSimpleAop(fmt.Sprintf(constant.BLOG_ARTICLE_BY_TAG, tagId), reflect.TypeOf([]model.ArticleListResp{})).
		WithExpires(time.Hour).
		WithEmptyExpires(5 * time.Minute).
		Then(func() (interface{}, error) {
			articleList := model.GetBlogsByTagId(tagId)
			for i, item := range articleList {
				articleList[i].Tags = model.GetTagsByArticleId(item.Id)
				articleList[i].User = userService.FindUserById(item.UserId)
				articleList[i].Type = *typeService.GetType(item.TypeId)
			}
			return articleList, nil
		})
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	var articleList = result.([]model.ArticleListResp)
	return articleList[start:stop]
}

func (a *ArticleService)GetArticleArchive() map[string][]model.ArticleListResp {
	years := model.GetDistinctYearFromBlog()
	result := make(map[string][]model.ArticleListResp)
	for _, year := range years {
		yearInt, _ := strconv.Atoi(year)
		result[year] = a.GetArticleByYear(yearInt)
	}
	return result
}
func (a *ArticleService)GetArticleByYear(year int) []model.ArticleListResp {
	result, _, err := g_rediscache.UseSimpleAop(fmt.Sprintf(constant.BLOG_ARTICLE_BY_TAG, year), reflect.TypeOf([]model.ArticleListResp{})).
		WithExpires(time.Hour).
		WithEmptyExpires(5 * time.Minute).
		Then(func() (interface{}, error) {
			return model.GetBlogsByYear(year),nil
		})
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	return result.([]model.ArticleListResp)
}


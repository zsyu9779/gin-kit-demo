package service

import "gin-kit-demo/model"

type TagService struct {
}

func (t *TagService) GetTags(pageSize, page int) []model.Tag {
	return model.GetTagsPage(page, pageSize)
}

func (t *TagService) GetTag(tagId int64) *model.Tag {
	return model.GetTagById(tagId)
}

func (t *TagService) DeleteTag(tagId int64) bool {
	tag := model.Tag{Id: tagId}
	return tag.DeleteTag()
}

func (t *TagService) UpdateTags(tag *model.Tag) bool {
	return tag.UpdateTag()
}

func (t *TagService) AddTag(tag *model.Tag) bool {
	return tag.AddTag()
}

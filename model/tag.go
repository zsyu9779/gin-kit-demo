package model

type Tag struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
}

type BlogTag struct {
	Id int64 `json:"id"`
	BlogId int64 `json:"blog_id"`
	TagId int64 `json:"tag_id"`
}

func (b *BlogTag) AddBlogTag() {
	Db.Create(&b)
}

func (b *BlogTag)Delete()  {
	Db.Where("blog_id",b.BlogId).Delete(&b)
}

func GetAllTags(page,pageSize int) []Tag {
	var result []Tag
	Db.Scopes(Paginate(page,pageSize)).Find(&result)
	return result
}

func GetTagById(id int64) *Tag {
	var result Tag
	Db.Find(&result,id)
	return &result
}
func (t *Tag)DeleteTag() bool {
	num := Db.Delete(&t).RowsAffected
	if num>0 {
		return true
	}else {
		return false
	}
}

func (t *Tag)UpdateTag() bool {
	num := Db.Save(&t).RowsAffected
	if num>0 {
		return true
	}else {
		return false
	}
}
func (t *Tag)AddTag() bool {
	num := Db.Create(&t).RowsAffected
	if num>0 {
		return true
	}else {
		return false
	}
}
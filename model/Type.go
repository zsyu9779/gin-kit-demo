package model

type Type struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

func GetAllTypes(page,pageSize int) []Type {
	var result []Type
	Db.Scopes(Paginate(page,pageSize)).Find(&result)
	return result
}

func GetTypeById(id int64) *Type {
	var result Type
	Db.Find(&result,id)
	return &result
}
func (t *Type)DeleteType() bool {
	num := Db.Delete(&t,t.Id).RowsAffected
	if num>0 {
		return true
	}else {
		return false
	}
}

func (t *Type)UpdateType() bool {
	num := Db.Save(&t).RowsAffected
	if num>0 {
		return true
	}else {
		return false
	}
}
func (t *Type)AddType() bool {
	num := Db.Create(&t).RowsAffected
	if num>0 {
		return true
	}else {
		return false
	}
}
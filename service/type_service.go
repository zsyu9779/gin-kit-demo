package service

import "gin-kit-demo/model"

type TypeService struct {

}

func (t *TypeService) GetTypes(pageSize, page int) []model.Type {
	return model.GetAllTypes(page, pageSize)
}

func (t *TypeService) GetType(typeId int64) *model.Type {
	return model.GetTypeById(typeId)
}

func (t *TypeService) DeleteType(typeId int64) bool {
	type1 := model.Type{Id: typeId}
	return type1.DeleteType()
}

func (t *TypeService) UpdateType(type1 *model.Type) bool {
	return type1.UpdateType()
}

func (t *TypeService) AddType(type1 *model.Type) bool {
	return type1.UpdateType()
}
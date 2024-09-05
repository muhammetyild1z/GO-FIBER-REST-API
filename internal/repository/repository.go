package repository

import (
	"go-auth/database"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{DB: database.DB}
}

func (r *Repository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}
func (r *Repository[T]) Dpdate(entity *T) error {
	return r.DB.Save(entity).Error
}

func (r *Repository[T]) Delete(entity *T) error {
	return r.DB.Delete(entity).Error
}
func (r *Repository[T]) GetByID(id uint, entity *T) error {
	return r.DB.First(entity, id).Error
}

func (r *Repository[T]) GetAll(entities *[]T) error {
	return r.DB.Find(entities).Error
}
func (r *Repository[T]) FindByField(field string, value interface{}, entity *T) error {
	return r.DB.Where(field+" = ?", value).First(entity).Error
}

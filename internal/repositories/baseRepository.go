package repositories

import (

	"gorm.io/gorm"
)


type BaseRepository[T any] interface {
	Create(entity T) error
	Update(entity T) error
	FindById(id int) (*T, error)
	Delete(id int) error
	FindAll(condition T) ([]*T, error)
	FindOne(condition T) (*T, error)
}

type baseRepository[T any] struct {
	db *gorm.DB
}


func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}


func (r *baseRepository[T]) Create(entity T) error {
	return r.db.Create(entity).Error
}

func (r *baseRepository[T]) Update(entity T) error {
	return r.db.Save(entity).Error
}

func (r *baseRepository[T]) FindById(id int) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	return &entity, err
}

func (r *baseRepository[T]) Delete(id int) error {
	return r.db.Delete(new(T), id).Error
}

func (r *baseRepository[T]) FindAll(condition T) ([]*T, error) {
	var entities []*T
	err := r.db.Where(condition).Find(&entities).Error
	return entities, err
}

func (r *baseRepository[T]) FindOne(condition T) (*T, error) {
	var entity T
	err := r.db.Where(condition).Find(&entity).Error
	return &entity, err
}
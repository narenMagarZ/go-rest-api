package repositories

import (

	"gorm.io/gorm"
)


type BaseRepository[T any] interface {
	Create(entity T) error
	UpdateOne(id int, entity T) error
	FindById(id int) (*T, error)
	Delete(id int) error
	FindAll(condition T) ([]*T, error)
	FindOne(condition T) (*T, error)
	Count(condition T) *int64
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}

func (r *baseRepository[T]) Create(entity T) error {
	return r.db.Create(&entity).Error
}

func (r *baseRepository[T]) UpdateOne(id int, entity T) error {
	return r.db.Where("id = ?", id).Updates(&entity).Error
}

func (r *baseRepository[T]) FindById(id int) (*T, error) {
	var entity T
	err := r.db.Select("id", "username", "email", "created_at", "updated_at").First(&entity, id).Error
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
	err := r.db.Where(condition).First(&entity).Error
	return &entity, err
}

func (r *baseRepository[T]) Count(condition T) *int64 {
	var count *int64
	r.db.Where(condition).Count(count);
	return count
}
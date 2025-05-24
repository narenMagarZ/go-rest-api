package repositories

import (
	"rest-api/internal/types"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type BaseRepository[T any] interface {
	Create(entity T) error
	UpdateOne(id int, entity T) error
	FindById(id int) (*T, error)
	Delete(id int) error
	FindAll(args types.CursorPaginationArgs) ([]*T, error)
	FindOne(condition T) (*T, error)
	Count(where any) *int64
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

func (r *baseRepository[T]) FindAll(args types.CursorPaginationArgs) ([]*T, error) {
	var entities []*T
	err := r.db.Where(args.Where).Find(&entities).
	Limit(args.Limit).
	Order(clause.OrderByColumn{
		Column: clause.Column{
			Name: args.Order}, 
			Desc: args.Sort,
	}).
	Error
	return entities, err
}

func (r *baseRepository[T]) FindOne(condition T) (*T, error) {
	var entity T
	err := r.db.Where(condition).First(&entity).Error
	return &entity, err
}

func (r *baseRepository[T]) Count(where any) *int64 {
	var count *int64
	r.db.Where(where).Count(count);
	return count
}
package repositories

import (

	"gorm.io/gorm"
)


type BaseRepository[T any] interface {
	Create(entity T)
	Update(entity T)
	FindById(id int)
	Delete(id int)
	FindAll()
}

type baseRepository[T any] struct {
	db *gorm.DB
}


func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}


func (r *baseRepository[T]) Create(entity T) {
	
}

func (r *baseRepository[T]) Update(entity T) {

}

func (r *baseRepository[T]) FindById(id int) {
}

func (r *baseRepository[T]) Delete(id int) {
	
}

func (r *baseRepository[T]) FindAll() {

}
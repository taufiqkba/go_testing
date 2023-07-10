package repository

import "github.com/taufiqkba/go_testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

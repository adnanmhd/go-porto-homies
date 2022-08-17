package repo

import (
	"context"
	"github.com/adnanmhd/go-porto-homies/app/entity"
)

type IPropertyRepo interface {
	SaveProperty(context.Context, entity.Property) error
	GetProperties(context.Context) ([]entity.Property, error)
}

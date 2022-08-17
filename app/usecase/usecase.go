package usecase

import (
	"context"
	"github.com/adnanmhd/go-porto-homies/app/entity"
)

type Property interface {
	List(context.Context) ([]entity.Property, error)
	Add(context.Context, entity.Property) error
}

package usecase

import (
	"context"

	"github.com/adnanmhd/go-porto-homies/internal/entity"
)

type (
	Property interface {
		List(context.Context) ([]entity.Property, error)
		Add(context.Context, entity.Property) error
	}
	PropertyRepo interface {
		SaveProperty(context.Context, entity.Property) error
		GetProperties(context.Context) ([]entity.Property, error)
	}
)

package usecase

import (
	"context"
	"fmt"
	"github.com/adnanmhd/go-porto-homies/app/entity"
	"github.com/adnanmhd/go-porto-homies/app/repo"
)

type PropertyUseCase struct {
	repo repo.IPropertyRepo
}

func New(r repo.IPropertyRepo) *PropertyUseCase {
	return &PropertyUseCase{
		repo: r,
	}
}

func (uc *PropertyUseCase) List(ctx context.Context) ([]entity.Property, error) {
	list, err := uc.repo.GetProperties(context.Background())
	if err != nil {
		return nil, fmt.Errorf("PropertyUseCase - List - repo.GetProperties: %w", err)
	}
	return list, nil
}

func (uc *PropertyUseCase) Add(ctx context.Context, property entity.Property) error {
	err := uc.repo.SaveProperty(context.Background(), property)
	return err
}

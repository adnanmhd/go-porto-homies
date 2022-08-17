package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/adnanmhd/go-porto-homies/app/entity"
)

const _defaultEntityCap = 64

type PropertyRepo struct {
	*sql.DB
}

func New(db *sql.DB) *PropertyRepo {
	return &PropertyRepo{db}
}

func (c *PropertyRepo) GetProperties(ctx context.Context) ([]entity.Property, error) {
	list := make([]entity.Property, 0, _defaultEntityCap)
	rec, err := c.Query("select p.id, p.name, p.sub_name, p.address, pc.category_name, p.selling_price, p.renting_price, p.building_area, p.building_floor, p.rooms, IFNULL(p.description, ''), p.is_sold from property p, property_category pc where pc.id = p.property_category and p.is_deleted = 0")
	if err != nil {
		return nil, fmt.Errorf("PropertyRepo - GetProperties: %w", err)

	}
	defer rec.Close()

	for rec.Next() {
		e := entity.Property{}
		e.PropertyDetail = entity.PropertyDetail{}
		e.PropertyType = entity.PropertyType{}

		err = rec.Scan(&e.Id, &e.Name, &e.SubName, &e.Address, &e.PropertyType.Type, &e.Price, &e.Price, &e.PropertyDetail.BuildingArea, &e.PropertyDetail.SurfaceArea, &e.PropertyDetail.Rooms, &e.Description, &e.Status)
		if err != nil {
			return nil, fmt.Errorf("PropertyRepo - Error Scan: %w", err)
		}
		list = append(list, e)
	}
	return list, nil
}

func (c *PropertyRepo) SaveProperty(ctx context.Context, property entity.Property) error {
	return nil
}

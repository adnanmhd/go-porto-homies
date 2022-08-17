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
	rec, err := c.Query("select p.id, p.name, p.sub_name, p.address, pc.category_name, p.selling_price, p.renting_price, IFNULL(p.renting_type, ''), IFNULL(p.img, ''), p.building_area, p.building_floor, p.rooms, IFNULL(p.description, ''), p.is_sold from property p, property_category pc where pc.id = p.property_category and p.is_deleted = 0")
	if err != nil {
		return nil, fmt.Errorf("PropertyRepo - GetProperties: %w", err)

	}
	defer rec.Close()

	for rec.Next() {
		e := entity.Property{}

		err = rec.Scan(&e.Id, &e.Name, &e.SubName, &e.Address, &e.CategoryName, &e.SellingPrice, &e.RentingPrice, &e.RentingType, &e.ImgPath, &e.BuildingArea, &e.BuildingFloor, &e.Rooms, &e.Description, &e.IsSold)
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

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
	rec, err := c.Query("select p.id, p.name, IFNULL(p.description, ''), pt.`type`, p.amenities, p.address, p.city, " +
		"p.price, IFNULL(p.renting_type, ''), p.status, p.images, p.created_by, p.created_at, " +
		"IFNULL(pd.rooms, 0), IFNULL(pd.bathroom, 0), IFNULL(pd.surface_area, 0), IFNULL(pd.building_area, 0), IFNULL(pd.year_built, 0), IFNULL(pd.legal_paper, '') " +
		"from property_type pt, property p " +
		"left join property_detail pd " +
		"on pd.property_id = p.id " +
		"where pt.id = p.property_type_id")

	if err != nil {
		return nil, fmt.Errorf("PropertyRepo - GetProperties: %w", err)

	}
	defer rec.Close()

	for rec.Next() {
		e := entity.Property{}
		e.PropertyDetail = entity.PropertyDetail{}
		e.PropertyType = entity.PropertyType{}

		err = rec.Scan(&e.Id, &e.Name, &e.Description, &e.PropertyType.Type, &e.Amenities, &e.Address, &e.City,
			&e.Price, &e.RentingType, &e.Status, &e.Images, &e.CreatedBy, &e.CreatedAt,
			&e.PropertyDetail.Rooms, &e.PropertyDetail.Bathroom, &e.PropertyDetail.SurfaceArea, &e.PropertyDetail.BuildingArea,
			&e.PropertyDetail.YearBuilt, &e.PropertyDetail.LegalPaper)

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

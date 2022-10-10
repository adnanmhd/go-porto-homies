package entity

type Response struct {
	Status     string     `json:"status,omitempty"`
	Message    string     `json:"message,omitempty"`
	Properties []Property `json:"properties,omitempty"`
}
type Request struct {
	Property Property
}

type Property struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	SubName        string         `json:"sub_name,omitempty"`
	Description    string         `json:"description,omitempty"`
	Address        string         `json:"address,omitempty"`
	City           string         `json:"city,omitempty"`
	Price          int            `json:"price,omitempty"`
	Status         string         `json:"status,omitempty"`
	RentingType    string         `json:"renting_type,omitempty"`
	Images         string         `json:"images,omitempty"`
	Amenities      string         `json:"amenities,omitempty"`
	PropertyType   PropertyType   `json:"property_type,omitempty"`
	PropertyDetail PropertyDetail `json:"property_detail,omitempty"`
	CreatedBy      string         `json:"created_by,omitempty"`
	CreatedAt      string         `json:"created_at,omitempty"`
	UpdatedBy      string         `json:"updated_by,omitempty"`
	UpdatedAt      string         `json:"updated_at,omitempty"`
}

type PropertyType struct {
	Id   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type PropertyDetail struct {
	Id           string `json:"id,omitempty"`
	Rooms        int    `json:"rooms,omitempty"`
	Bathroom     int    `json:"bath_room,omitempty"`
	SurfaceArea  int    `json:"surface_area,omitempty"`
	BuildingArea int    `json:"building_area,omitempty"`
	YearBuilt    int    `json:"year_built,omitempty"`
	LegalPaper   string `json:"legal_paper,omitempty"`
}

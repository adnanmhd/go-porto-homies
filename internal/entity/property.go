package entity

type Property struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	SubName       string `json:"sub_name,omitempty"`
	Address       string `json:"address,omitempty"`
	CategoryName  string `json:"category_name,omitempty"`
	SellingPrice  int64  `json:"selling_price,omitempty"`
	RentingPrice  int64  `json:"renting_price,omitempty"`
	RentingType   string `json:"renting_type,omitempty"`
	ImgPath       string `json:"img_path,omitempty"`
	BuildingArea  string `json:"building_area,omitempty"`
	BuildingFloor string `json:"building_floor,omitempty"`
	Rooms         int8   `json:"rooms,omitempty"`
	Description   string `json:"description,omitempty"`
	IsSold        string `json:"is_sold,omitempty"`
	CreatedBy     string `json:"created_by,omitempty"`
	CreatedDate   string `json:"created_date,omitempty"`
	UpdatedBy     string `json:"updated_by,omitempty"`
	UpdatedDate   string `json:"updated_date,omitempty"`
}

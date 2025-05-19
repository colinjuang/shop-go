package response

type AddressResponse struct {
	ID           uint64 `json:"id"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	City         string `json:"city"`
	CityCode     string `json:"cityCode"`
	Province     string `json:"province"`
	ProvinceCode string `json:"provinceCode"`
	District     string `json:"district"`
	DistrictCode string `json:"districtCode"`
	DetailAddr   string `json:"detailAddr"`
	FullAddr     string `json:"fullAddr"`
	IsDefault    int8   `json:"isDefault"`
}

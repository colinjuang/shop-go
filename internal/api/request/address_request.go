package request

type AddressRequest struct {
	ID           uint64 `json:"id"`
	Phone        string `json:"phone" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Province     string `json:"province" binding:"required"`
	ProvinceCode string `json:"provinceCode" binding:"required"`
	City         string `json:"city" binding:"required"`
	CityCode     string `json:"cityCode" binding:"required"`
	District     string `json:"district" binding:"required"`
	DistrictCode string `json:"districtCode" binding:"required"`
	DetailAddr   string `json:"detailAddr" binding:"required"`
	IsDefault    int8   `json:"isDefault"`
}

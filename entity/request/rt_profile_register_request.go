package request

type RTProfileRegisterRequest struct {
	RTNumber    string  `json:"rt_number" validate:"required"`
	VillageID   int     `json:"village_id" validate:"required"`
	RTEmail     string  `json:"rt_email" validate:"required" gorm:"uniqueIndex"`
	MobilePhone string  `json:"mobile_phone" validate:"required"`
	FullAddress string  `json:"full_address" validate:"required"`
	Description string  `json:"description"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	ReferalCode string  `json:"referal_code"`
}

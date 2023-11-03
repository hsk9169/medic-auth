package http

type GetAccountInfoReqToBasic struct {
	UserID string
	Role   string
}

type MedicDataFromBasicApi struct {
	Phone        string `json:"phone"`
	Username     string `json:"username"`
	HospitalCode string `json:"hospitalCode"`
	WardCode     string `json:"wardCode"`
	Position     string `json:"position"`
	AuthFlag     bool   `json:"authFlag"`
}

type UserDataFromBasicApi struct {
}

package types

type GetUserDataRes struct {
	Id       uint64              `json:"id"`
	Username string              `json:"username"`
	Password string              `json:"password"`
	Role     string              `json:"role"`
	Allowed  []UserAllowedAccess `json:"allowed"`
}

type UserAllowedAccess struct {
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
}

type GetUserDataReq struct {
	Id uint64 `json:"id"`
}

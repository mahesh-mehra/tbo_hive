package objects

type AuthReq struct {
	Mobile string `json:"mobile"`
}

type ValidateOtpReq struct {
	Mobile string `json:"mobile"`
	Otp    string `json:"otp"`
}

type ValidateProfileReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

type FollowReq struct {
	UserId string `json:"user_id"`
}

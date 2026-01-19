package objects

type Response struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

type ResponseWithData struct {
	Status bool            `json:"status"`
	Data   UserProfileData `json:"data"`
}
